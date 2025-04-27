package mg

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type DecodedNumber struct {
	Sign            int
	IsIntegral      bool
	IsNonFinite     bool
	IsBalanced      bool
	HasExplicitBase bool
	Base            int
	Exponent        int
	Mantissa        string
}

func (d *DecodedNumber) IntValue() (int, error) {
	if !d.IsIntegral {
		f, e := d.FloatValue()
		if e != nil {
			return 0, e
		}
		return int(f), nil
	}
	sofar := 0
	base := d.Base
	for _, r := range d.Mantissa {
		increment := 0
		d0 := int(r - '0')
		if 0 <= d0 && d0 <= 9 {
			increment = d0
		} else if d.IsBalanced {
			if r == 'T' {
				increment = -1
			} else {
				return 0, strconv.ErrSyntax
			}
		} else {
			d1 := int(r - 'A')
			if 0 <= d1 && d1 <= 35 {
				increment = d1 + 10
			} else {
				return 0, strconv.ErrSyntax
			}
		}
		if increment < 0 && !d.IsBalanced || increment >= base {
			return 0, strconv.ErrSyntax
		}
		sofar = sofar*base + increment
	}
	return d.Sign * sofar, nil
}

func (d *DecodedNumber) FloatValue() (float64, error) {
	if d.IsIntegral {
		i, e := d.IntValue()
		if e != nil {
			return 0, e
		}
		return float64(i), nil
	}
	seenDot := false
	sofar := 0.0
	base := d.Base
	fbase := float64(base)
	decimalPlaces := 0
	for _, r := range d.Mantissa {
		if seenDot {
			decimalPlaces++
		}
		if r == '.' {
			seenDot = true
		} else {
			increment := 0
			d0 := int(r - '0')
			if 0 <= d0 && d0 <= 9 {
				increment = d0
			} else if d.IsBalanced {
				if r == 'T' {
					increment = -1
				} else {
					return 0, strconv.ErrSyntax
				}
			} else {
				d1 := int(r - 'A')
				if 0 <= d1 && d1 <= 35 {
					increment = d1 + 10
				} else {
					return 0, strconv.ErrSyntax
				}
			}
			if increment < 0 && !d.IsBalanced || increment >= base {
				return 0, strconv.ErrSyntax
			}
			sofar = sofar*fbase + float64(increment)
		}
	}
	return float64(d.Sign) * sofar * math.Pow(fbase, float64(d.Exponent-decimalPlaces)), nil
}

func decodeNumber(text string) (DecodedNumber, error) {
	decoded := DecodedNumber{}

	decoded.Sign = 1
	if text[0] == '-' {
		decoded.Sign = -1
		text = text[1:]
	} else if text[0] == '+' {
		text = text[1:]
	}

	decoded.Base = 10
	if strings.HasPrefix(text, hexPrefix) {
		decoded.Base = 16
		text = text[len(hexPrefix):]
	} else if strings.HasPrefix(text, binPrefix) {
		decoded.Base = 2
		text = text[len(binPrefix):]
	} else if strings.HasPrefix(text, octPrefix) {
		decoded.Base = 8
		text = text[len(octPrefix):]
	} else if strings.HasPrefix(text, nonFinitePrefix) {
		decoded.IsNonFinite = true
		text = text[len(nonFinitePrefix):]
	} else if strings.HasPrefix(text, balancedTernaryPrefix) {
		decoded.IsBalanced = true
		decoded.Base = 3
		text = text[len(balancedTernaryPrefix):]
	} else if strings.ContainsRune(text, radixRune) {
		decoded.HasExplicitBase = true
		d0 := int(text[0] - '0')
		if d0 < 0 || d0 > 35 {
			return decoded, strconv.ErrSyntax
		}
		if text[1] == radixRune {
			decoded.Base = d0
			text = text[2:]
		} else if text[2] == radixRune {
			d1 := int(text[1] - '0')
			if d1 < 0 || d1 > 35 {
				return decoded, strconv.ErrSyntax
			}
			decoded.Base = d0*10 + d1
			text = text[3:]
		} else {
			return decoded, strconv.ErrSyntax
		}
		if decoded.Base < 2 || decoded.Base > 36 {
			return decoded, strconv.ErrSyntax
		}
	} else if strings.HasPrefix(text, "∞") {
		var rewritten strings.Builder
		if decoded.Sign < 0 {
			rewritten.WriteRune('-')
		}
		rewritten.WriteString(nonFinitePrefix)
		rewritten.WriteRune('1')
		rewritten.WriteString(text[1:])

		d, e := decodeNumber(rewritten.String())
		return d, e
	} else if strings.HasPrefix(text, "⦰") {
		var rewritten strings.Builder
		rewritten.WriteString(nonFinitePrefix)
		rewritten.WriteRune('0')
		rewritten.WriteString(text[1:])

		d, e := decodeNumber(rewritten.String())
		return d, e
	}

	text = strings.ReplaceAll(text, "_", "")

	decoded.Exponent = 0
	if strings.ContainsRune(text, 'e') {
		parts := strings.Split(text, "e")
		text = parts[0]
		exponent, err := strconv.Atoi(parts[1])
		if err == nil {
			// convert exp to an integer
			decoded.Exponent = exponent
		}
	}

	decoded.Mantissa = text
	decoded.IsIntegral = !strings.ContainsRune(text, '.') && decoded.Exponent == 0

	if isOnlyZeros(text) {
		decoded.Sign = 0
	}

	return decoded, nil
}

func isOnlyZeros(text string) bool {
	for _, r := range text {
		if r != '0' && r != '.' {
			// We allow 0.0, 0.00, 0.000, etc.
			return false
		}
	}
	return true
}

// ConvertToDenary would be a better name in many ways but "denary" is much less
// common in English than decimal. However decimal can imply a decimal point,
// which is not necessarily the case here.
func ConvertToDecimal(text string) (string, error) {
	decoded, err := decodeNumber(text)
	if err != nil {
		return "", err
	}
	if decoded.Base == 10 && !decoded.HasExplicitBase {
		// Guarantee stability for base 10 - as long as `r` notation is not used.
		return strings.ReplaceAll(text, "_", ""), nil
	}
	if decoded.IsIntegral {
		n, e := decoded.IntValue()
		if e != nil {
			return "", e
		}
		return fmt.Sprintf("%d", n), nil
	}
	f, e := decoded.FloatValue()
	if e != nil {
		return "", e
	}
	return strconv.FormatFloat(f, 'g', -1, 64), nil
}

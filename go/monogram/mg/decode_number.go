package mg

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type DecodedNumber struct {
	Sign         int
	IsIntegral   bool
	Base         int
	ExplicitBase bool
	Exponent     int
	Mantissa     string
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
		d0 := int(r - '0')
		if 0 <= d0 && d0 <= 9 {
			sofar = sofar*base + (d0)
		} else {
			d1 := int(r - 'A')
			if 0 <= d1 && d1 <= 35 {
				sofar = sofar*base + (d1 + 10)
			} else {
				return 0, strconv.ErrSyntax
			}
		}
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
	base := float64(d.Base)
	decimalPlaces := 0
	for _, r := range d.Mantissa {
		if seenDot {
			decimalPlaces++
		}
		if r == '.' {
			seenDot = true
		} else {
			d0 := int(r - '0')
			if 0 <= d0 && d0 <= 9 {
				sofar = sofar*base + float64(d0)
			} else {
				d1 := int(r - 'A')
				if 0 <= d1 && d1 <= 35 {
					sofar = sofar*base + float64(d1+10)
				} else {
					fmt.Println("Got wrong here", r)
					return 0, strconv.ErrSyntax
				}
			}
		}
	}
	return float64(d.Sign) * sofar * math.Pow(base, float64(d.Exponent-decimalPlaces)), nil
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
	if strings.HasPrefix(text, "0x") {
		decoded.Base = 16
		text = text[2:]
	} else if strings.HasPrefix(text, "0b") {
		decoded.Base = 2
		text = text[2:]
	} else if strings.HasPrefix(text, "0o") {
		decoded.Base = 8
		text = text[2:]
	} else if strings.ContainsRune(text, radixRune) {
		decoded.ExplicitBase = true
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
	}

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
	if decoded.Base == 10 && !decoded.ExplicitBase {
		// Guarantee stability for base 10 - as long as `r` notation is not used.
		return text, nil
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

package mg

import (
	"testing"
)

func TestAsInt(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"0x1", 1},
		{"0x10", 16},
		{"0x100", 256},
		{"0x1000", 4096},
		{"123", 123},
		{"36rZZ", 1295},
		{"-36rZZ", -1295},
	}

	for _, c := range cases {
		d, err := decodeNumber(c.input)
		if err != nil {
			t.Errorf("Error decoding '%s': %v", c.input, err)
			continue
		}
		i, err := d.IntValue()
		if err != nil {
			t.Errorf("Error converting '%s' to int: %v", c.input, err)
			continue
		}
		if i != c.expected {
			t.Errorf("Expected %d, got %d for input '%s'", c.expected, i, c.input)
		}
	}
}

func TestAsFloat(t *testing.T) {
	cases := []struct {
		input    string
		expected float64
	}{
		{"99.123", 99.123},
		{"0x1.8", 1.5},
		{"-0x1.8e+1", -24.0},
		{"-0x18e-2", -0.09375},
		{"-123.567e+1", -1235.67},
		{"0b1101.101", 13.625},
	}

	for _, c := range cases {
		d, err := decodeNumber(c.input)
		if err != nil {
			t.Errorf("Error decoding '%s': %v", c.input, err)
			continue
		}
		f, err := d.FloatValue()
		if err != nil {
			t.Errorf("Error converting '%s' to float: %v", c.input, err)
			continue
		}
		if f != c.expected {
			t.Errorf("Expected %v, got %v for input '%s'", c.expected, f, c.input)
		}
	}
}

func TestDecode(t *testing.T) {
	cases := []struct {
		input       string
		sign        int
		is_integral bool
		base        int
		mantissa    string
		exponent    int
	}{
		{"0x1", 1, true, 16, "1", 0},
		{"0x1.8", 1, false, 16, "1.8", 0},
		{"0x1.8e+1", 1, false, 16, "1.8", 1},
		{"0x1.8e-1", 1, false, 16, "1.8", -1},
		{"16r1.8e-1", 1, false, 16, "1.8", -1},
		{"123", 1, true, 10, "123", 0},
		{"123.456", 1, false, 10, "123.456", 0},
		{"123.456e+1", 1, false, 10, "123.456", 1},
		{"123.456e-1", 1, false, 10, "123.456", -1},
		{"-123", -1, true, 10, "123", 0},
		{"-123.456", -1, false, 10, "123.456", 0},
		{"-123.456e+1", -1, false, 10, "123.456", 1},
		{"-123.456e-1", -1, false, 10, "123.456", -1},
		{"0", 0, true, 10, "0", 0},
		{"000", 0, true, 10, "000", 0},
		{"0.0", 0, false, 10, "0.0", 0},
		{"000.00", 0, false, 10, "000.00", 0},
		{"0.0e+1", 0, false, 10, "0.0", 1},
		{"0.0e-1", 0, false, 10, "0.0", -1},
	}

	for _, c := range cases {
		d, err := decodeNumber(c.input)
		if err != nil {
			t.Errorf("Error decoding '%s': %v", c.input, err)
		}
		if d.IsIntegral != c.is_integral {
			t.Errorf("Expected IsIntegral to be %v, got %v", c.is_integral, d.IsIntegral)
		}
		if d.Base != c.base {
			t.Errorf("Expected Base to be %d, got %d", c.base, d.Base)
		}
		if d.Mantissa != c.mantissa {
			t.Errorf("Expected Mantissa to be '%s', got '%s'", c.mantissa, d.Mantissa)
		}
		if d.Exponent != c.exponent {
			t.Errorf("Expected Exponent to be %d, got %d", c.exponent, d.Exponent)
		}
	}
}

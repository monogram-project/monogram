package lib

import (
	"testing"
)

func TestReadStringInterpolation(t *testing.T) {
	cases := []struct {
		input        string
		expectedText string
		shouldError  bool
	}{
		{"\\(x)", "(x)", false},
		{"\\[x]", "[x]", false},
		{"\\(x + y)", "(x + y)", false},
		{"\\(x.f(y, z))", "(x.f(y, z))", false},
		{"\\(mismatched", "", true},
		{"\\{'string!'}", "{'string!'}", false},
		{"\\{'string!\\[99]'}", "{'string!\\[99]'}", false},
		{"\\{'string!\\[99'}", "", true},
	}

	for _, c := range cases {
		tokenizer := NewTokenizer(c.input)
		tokenizer.consume() // Consume the backslash
		token, err := tokenizer.readStringInterpolation()

		if c.shouldError {
			if err == nil {
				t.Errorf("Expected error for input %q, but got none", c.input)
			}
			continue
		}

		if err != nil {
			t.Errorf("Unexpected error for input %q: %v", c.input, err)
			continue
		}

		if token.Type != Literal || token.SubType != LiteralExpressionString {
			t.Errorf("Unexpected token type/subtype for input %q: got (%v, %v), want (%v, %v)",
				c.input, token.Type, token.SubType, Literal, LiteralExpressionString)
		}

		if token.Text != c.expectedText {
			t.Errorf("Unexpected token text for input %q: got %v, want %v", c.input, token.Text, c.expectedText)
		}
	}
}

func TestReadStringWithoutInterpolation(t *testing.T) {
	cases := []struct {
		input        string
		expectedText string
		shouldError  bool
	}{
		// Simple string test
		{"\"simple string\"", "simple string", false},

		// String with escape sequences
		{"\"escaped \\\"quote\\\" and newline\\n\"", "escaped \"quote\" and newline\n", false},

		// Unterminated string
		{"\"unterminated", "", true},
	}

	for _, c := range cases {
		tokenizer := NewTokenizer(c.input)
		token, err := tokenizer.readString()

		if c.shouldError {
			if err == nil {
				t.Errorf("Expected error for input %q, but got none", c.input)
			}
			continue
		}

		if err != nil {
			t.Errorf("Unexpected error for input %q: %v", c.input, err)
			continue
		}

		if token.Type != Literal || token.SubType != LiteralString {
			t.Errorf("Unexpected token type for input %q: got (%v, %v), want (%v, %v)",
				c.input, token.Type, token.SubType, Literal, LiteralString)
		}

		if token.Text != c.expectedText {
			t.Errorf("Unexpected token text for input %q: got %v, want %v", c.input, token.Text, c.expectedText)
		}
	}
}

func TestReadStringWithInterpolation(t *testing.T) {
	cases := []struct {
		input             string
		expectedSubTokens []string // Expected sub-token texts
		shouldError       bool
	}{
		// String with interpolation
		{"\"Hello \\(name)!\"", []string{"Hello ", "(name)", "!"}, false},
		{"\"Hello \\(name)\\(`!`)\"", []string{"Hello ", "(name)", "(`!`)"}, false},

		// Interpolation with mismatched brackets
		{"\"Mismatched \\(x + y\"", nil, true},
	}

	for _, c := range cases {
		tokenizer := NewTokenizer(c.input)
		token, err := tokenizer.readString()

		if c.shouldError {
			if err == nil {
				t.Errorf("Expected error for input %q, but got none", c.input)
			}
			continue
		}

		if err != nil {
			t.Errorf("Unexpected error for input %q: %v", c.input, err)
			continue
		}

		if token.Type != Literal || token.SubType != LiteralInterpolatedString {
			t.Errorf("Unexpected token type for input %q: got (%v, %v), want (%v, %v)",
				c.input, token.Type, token.SubType, Literal, LiteralInterpolatedString)
		}

		// Check sub-tokens for interpolated strings
		if len(token.SubTokens) != len(c.expectedSubTokens) {
			t.Errorf("Unexpected number of sub-tokens for input %q: got %d, want %d",
				c.input, len(token.SubTokens), len(c.expectedSubTokens))
		}

		for i, sub := range token.SubTokens {
			if sub.Text != c.expectedSubTokens[i] {
				t.Errorf("Unexpected sub-token text for input %q: got %v, want %v",
					c.input, sub.Text, c.expectedSubTokens[i])
			}
		}
	}
}

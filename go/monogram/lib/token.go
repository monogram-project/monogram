package lib

import (
	"strings"
)

type TokenType int

const (
	// Major Types
	Literal TokenType = iota
	Identifier
	Punctuation
	OpenBracket
	CloseBracket
	Sign
)

// Subtypes for Literal
const (
	LiteralString uint8 = iota
	LiteralNumber
	LiteralInterpolatedString
	LiteralExpressionString
	LiteralMultilineString
)

// Subtypes for Identifier
const (
	IdentifierVariable uint8 = iota
	IdentifierFormStart
	IdentifierFormEnd
	IdentifierBreaker
	IdentifierCompoundBreaker
)

// Subtypes for Punctuation
const (
	PunctuationComma uint8 = iota
	PunctuationSemicolon
)

// Subtypes for Bracket
const (
	BracketParenthesis uint8 = iota
	BracketBrace
	BracketBracket
)

// Subtypes for Sign
const (
	SignLabel uint8 = iota
	SignForce
	SignDot
	SignMinus
	SignOperator
)

type Token struct {
	Type                 TokenType // The type of token (Sign, Bracket, etc.)
	SubType              uint8     // The specific subtype of the token (if any)
	Text                 string    // The raw text of the token
	Span                 Span      // The span of the token in the source code
	PrecededByNewline    bool      // New field to indicate if the token is preceded by a newline
	FollowedByWhitespace bool      // New field to indicate if the token is followed by whitespace
	EscapeSeen           bool      // New field to indicate if an escape sequence was seen
	IsMultiLine          bool      // New field to indicate if the token is a multi-line string
	QuoteRune            rune      // New field to indicate the quote rune for strings
	NextToken            *Token    // The next token in the chain

	SubTokens []*Token // Subtokens for interpolated string tokens

	// Cache for precedence
	precValue int  // Cached precedence value
	precValid bool // Indicates if the precedence has been computed
	errFlag   bool // Cached error flag for precedence validity
}

func (t *Token) SpanString() string {
	return t.Span.SpanString()
}

func (t *Token) QuoteWord() string {
	if t.QuoteRune == '"' {
		return "double"
	} else if t.QuoteRune == '\'' {
		return "single"
	} else if t.QuoteRune == '`' {
		return "backtick"
	} else {
		return "undefined"
	}
}

func (t *Token) IsBreaker(formStart *Token) bool {
	return t.IsSimpleBreaker() || t.IsCompoundBreaker(formStart)
}

func (t *Token) IsSimpleBreaker() bool {
	if t.Type != Identifier || t.SubType != IdentifierVariable {
		return false
	}
	if t.FollowedByWhitespace {
		return false
	}
	if t.NextToken == nil {
		return false
	}
	if t.NextToken.Type != Sign || t.NextToken.SubType != SignLabel {
		return false
	}
	return true
}

func (t *Token) IsCompoundBreaker(formStart *Token) bool {
	if t.Type != Identifier || t.SubType != IdentifierVariable {
		return false
	}
	if t.FollowedByWhitespace {
		return false
	}
	t1 := t.NextToken
	if t1 == nil {
		return false
	}
	if t1.Type != Sign || t1.SubType != SignMinus {
		return false
	}
	t2 := t1.NextToken
	if t2 == nil {
		return false
	}
	if t2.Type != Identifier || t2.SubType != IdentifierFormStart {
		return false
	}
	if t2.Text != formStart.Text {
		return false
	}
	return true
}

func (t *Token) IsLabel() bool {
	return t.Type == Sign && t.SubType == SignLabel
}

func (t *Token) IsMacro() bool {
	if t.Type != Identifier || t.SubType != IdentifierFormStart || t.FollowedByWhitespace {
		return false
	}
	t1 := t.NextToken
	if t1 == nil {
		return false
	}
	return t1.Type == Sign && t1.SubType == SignForce
}

func (t *Token) SetSeen(tokenizer *Tokenizer, seen bool) {
	t.PrecededByNewline = seen
	t.Span.EndLine = tokenizer.lineNo
	t.Span.EndColumn = tokenizer.colNo
}

const signChars = ".*/%+-<>~!&|?:="
const precCharacters = ".([*/%+-<>~!&|?:="

func (t *Token) DelimiterName() string {
	switch t.Type {
	case OpenBracket, CloseBracket:
		switch t.SubType {
		case BracketParenthesis:
			return "parentheses"
		case BracketBrace:
			return "braces"
		case BracketBracket:
			return "brackets"
		}
	}
	return ""
}

const (
	maxPrecedence    int = 999
	prefixPrecedence int = 1
)

func (t *Token) Precedence() (int, bool) {
	// Check if precedence is already cached
	if t.precValid {
		return t.precValue, !t.errFlag // Return cached result
	}

	// Precedence is only meaningful for Signs and Brackets
	if t.Type != Sign && t.Type != OpenBracket {
		t.precValue = 0
		t.precValid = true
		t.errFlag = true // Cache that this token has no valid precedence
		return 0, false
	}

	// Get the first rune of the token's text
	runes := []rune(t.Text)
	if len(runes) == 0 {
		// Invalid token with empty text
		t.precValue = 0
		t.precValid = true
		t.errFlag = true // Cache the error
		return 0, false
	}
	firstRune := runes[0]

	// Find the position of the first rune in the signs string
	pos := strings.IndexRune(precCharacters, firstRune)
	if pos == -1 {
		// If the rune is not in the signs string
		t.precValue = 0
		t.precValid = true
		t.errFlag = true // Cache the error
		return 0, false
	}

	// Calculate precedence
	P := (pos + 1) * 10
	if len(runes) > 1 && runes[0] == runes[1] {
		// If the first rune occurs twice in the token, subtract 1
		P--
	}

	// Cache the precedence result and success
	t.precValue = P
	t.precValid = true
	t.errFlag = false // Cache success (no error)

	return P, true
}

// VSCodeTokenType maps the token's type and subtype to a VSCode semantic token type.
func (t *Token) VSCodeTokenType() string {
	switch t.Type {
	case Literal:
		switch t.SubType {
		case LiteralString:
			return "string"
		case LiteralNumber:
			return "number"
		default:
			return "literal"
		}
	case Identifier:
		switch t.SubType {
		case IdentifierVariable:
			return "variable"
		case IdentifierFormStart:
			// Assuming a callable-like entity
			return "function"
		case IdentifierFormEnd:
			// End markers can be styled as keywords
			return "keyword"
		case IdentifierBreaker, IdentifierCompoundBreaker:
			return "operator"
		default:
			return "identifier"
		}
	case Punctuation:
		switch t.SubType {
		case PunctuationComma, PunctuationSemicolon:
			return "delimiter"
		default:
			return "punctuation"
		}
	case OpenBracket, CloseBracket:
		switch t.SubType {
		case BracketParenthesis:
			return "delimiter.parenthesis"
		case BracketBrace:
			return "delimiter.brace"
		case BracketBracket:
			return "delimiter.bracket"
		default:
			return "delimiter"
		}
	case Sign:
		switch t.SubType {
		case SignLabel:
			return "macro"
		case SignForce, SignDot, SignMinus, SignOperator:
			return "operator"
		default:
			return "operator"
		}
	default:
		return "unknown"
	}
}

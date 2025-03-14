package main

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type TokenType int

const (
	// Major Types
	Literal TokenType = iota
	Identifier
	Punctuation
	Bracket
	Sign
)

// Subtypes for Literal
const (
	LiteralString = iota
	LiteralNumber
)

// Subtypes for Identifier
const (
	IdentifierVariable = iota
	IdentifierFormStart
	IdentifierFormEnd
	IdentifierBreaker
	IdentifierCompoundBreaker
)

// Subtypes for Punctuation
const (
	PunctuationComma = iota
	PunctuationSemicolon
)

// Subtypes for Bracket
const (
	BracketOpenParenthesis = iota
	BracketOpenBrace
	BracketOpenBracket
	BracketCloseParenthesis
	BracketCloseBrace
	BracketCloseBracket
)

// Subtypes for Sign
const (
	SignLabel = iota
	SignForce
	SignDot
	SignOperator
)

type Token struct {
	Type        TokenType // The type of token (Sign, Bracket, etc.)
	SubType     int       // The specific subtype of the token (if any)
	Text        string    // The raw text of the token
	StartLine   int       // The starting line number of the token
	StartColumn int       // The starting column number of the token

	// Cache for precedence
	precValue int  // Cached precedence value
	precValid bool // Indicates if the precedence has been computed
	errFlag   bool // Cached error flag for precedence validity
}

type Tokenizer struct {
	input  string  // The input string to tokenize
	tokens []Token // The array of tokens generated
	lineNo int     // Current line number
	colNo  int     // Current column number
	pos    int     // Current byte position in the input
}

// Create a new Tokenizer
func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{
		input:  input,
		tokens: []Token{},
		lineNo: 1,
		colNo:  1,
		pos:    0,
	}
}

const signCharacters = ".({[*/%+-<~!&|?:="

func (t *Token) Precedence() (int, bool) {
	// Check if precedence is already cached
	if t.precValid {
		return t.precValue, !t.errFlag // Return cached result
	}

	// Precedence is only meaningful for Signs and Brackets
	if t.Type != Sign && t.Type != Bracket {
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
	pos := strings.IndexRune(signCharacters, firstRune)
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

// Advance the position within the input, updating line and column numbers
func (t *Tokenizer) advancePosition(r rune) {
	if r == '\n' {
		t.lineNo++
		t.colNo = 1
	} else {
		t.colNo++
	}
	t.pos += utf8.RuneLen(r) // Move the byte position forward
}

// Peek at the current rune in the input without advancing
func (t *Tokenizer) peek() (rune, int) {
	if t.pos >= len(t.input) {
		return rune(0), 0 // End of input
	}
	return utf8.DecodeRuneInString(t.input[t.pos:])
}

func (t *Tokenizer) peekN(n int) (rune, bool) {
	currentPos := t.pos // Start at the current position
	var r rune
	var size int

	// Iterate through the input to locate the nth rune
	for i := 0; i < n; i++ {
		if currentPos >= len(t.input) {
			// If we reach the end of input before finding the nth rune, return false
			return 0, false
		}

		r, size = utf8.DecodeRuneInString(t.input[currentPos:])
		if r == utf8.RuneError {
			// Handle invalid UTF-8 character by returning false
			return 0, false
		}

		// Advance to the next rune
		currentPos += size
	}

	// Return the nth rune and true (indicating success)
	return r, true
}

// Consume the current rune and advance the position
func (t *Tokenizer) consume() rune {
	r, _ := t.peek()
	t.advancePosition(r)
	return r
}

// Add a token to the token list
func (t *Tokenizer) addToken(tokenType TokenType, subType int, text string, startLine int, startCol int) {
	t.tokens = append(t.tokens, Token{
		Type:        tokenType,
		SubType:     subType,
		Text:        text,
		StartLine:   startLine,
		StartColumn: startCol,
	})
}

func (t *Tokenizer) tokenize() {
	for t.pos < len(t.input) {
		r, _ := t.peek()

		// Skip whitespace
		if unicode.IsSpace(r) {
			t.consume()
			continue
		}

		// Match strings
		if r == '"' || r == '\'' {
			t.readString()
			continue
		}

		// Match numbers
		if unicode.IsDigit(r) {
			t.readNumber()
			continue
		}

		// Match identifiers
		if unicode.IsLetter(r) || r == '_' {
			t.readIdentifier()
			continue
		}

		// Match punctuation
		if r == ',' || r == ';' {
			t.readPunctuation()
			continue
		}

		// Match brackets
		if r == '(' || r == ')' || r == '[' || r == ']' || r == '{' || r == '}' {
			t.readBracket()
			continue
		}

		// Match signs
		if t.isSign(r) {
			t.readSign()
			continue
		}

		// Discard unexpected characters
		t.consume()
	}
}

func (t *Tokenizer) isSign(r rune) bool {
	signChars := ".*/%+-<~!&|?:="
	return strings.ContainsRune(signChars, r)
}

func (t *Tokenizer) readSign() {
	startLine, startCol := t.lineNo, t.colNo
	start := t.pos

	for t.pos < len(t.input) {
		r, _ := t.peek()
		if !t.isSign(r) {
			break
		}
		t.consume()
	}

	// Add the sign token
	text := t.input[start:t.pos]
	t.addToken(Sign, -1, text, startLine, startCol) // -1 for now as signs may not have subtypes yet
}

func (t *Tokenizer) readBracket() {
	startLine, startCol := t.lineNo, t.colNo
	r := t.consume() // Consume the bracket character

	// Determine the subtype
	var subType int
	switch r {
	case '(':
		subType = BracketOpenParenthesis
	case ')':
		subType = BracketCloseParenthesis
	case '[':
		subType = BracketOpenBracket
	case ']':
		subType = BracketCloseBracket
	case '{':
		subType = BracketOpenBrace
	case '}':
		subType = BracketCloseBrace
	}

	// Add the bracket token
	t.addToken(Bracket, subType, string(r), startLine, startCol)
}

func (t *Tokenizer) readPunctuation() {
	startLine, startCol := t.lineNo, t.colNo
	r := t.consume() // Consume the punctuation character

	// Determine the subtype
	var subType int
	if r == ',' {
		subType = PunctuationComma
	} else if r == ';' {
		subType = PunctuationSemicolon
	}

	// Add the punctuation token
	t.addToken(Punctuation, subType, string(r), startLine, startCol)
}

func (t *Tokenizer) readString() {
	startLine, startCol := t.lineNo, t.colNo
	quote := t.consume() // Consume the opening quote
	start := t.pos

	for t.pos < len(t.input) {
		r := t.consume()
		if r == quote { // Closing quote found
			break
		}
		if r == '\\' && t.pos < len(t.input) { // Handle escape sequences
			t.consume()
		}
	}

	// Add the string token
	t.addToken(Literal, LiteralString, t.input[start-1:t.pos], startLine, startCol)
}

func (t *Tokenizer) readNumber() {
	startLine, startCol := t.lineNo, t.colNo
	start := t.pos
	hasDot := false

	for t.pos < len(t.input) {
		r, _ := t.peek()
		if r == '.' {
			if hasDot { // Invalid: multiple dots
				break
			}
			hasDot = true
		} else if !unicode.IsDigit(r) {
			break
		}
		t.consume()
	}

	// Add the number token
	text := t.input[start:t.pos]
	t.addToken(Literal, LiteralNumber, text, startLine, startCol)
}

func (t *Tokenizer) readIdentifier() {
	startLine, startCol := t.lineNo, t.colNo
	start := t.pos

	for t.pos < len(t.input) {
		r, _ := t.peek()
		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_') {
			break
		}
		t.consume()
	}

	// Add the identifier token
	t.addToken(Identifier, IdentifierVariable, t.input[start:t.pos], startLine, startCol)
}

func tokenizeInput(input string) []Token {
	// Create a new Tokenizer instance
	tokenizer := NewTokenizer(input)

	fmt.Println("Discovering Tokens ...")

	// Perform tokenization
	tokenizer.tokenize()

	// Print the tokens for scaffolding purposes
	fmt.Println("Discovered Tokens:")
	for _, token := range tokenizer.tokens {
		p, ok := token.Precedence()
		if !ok {
			p = -1
		}
		fmt.Printf("Token: {Type: %d, SubType: %d, Text: %q, StartLine: %d, StartColumn: %d, Precedence: %d}\n",
			token.Type, token.SubType, token.Text, token.StartLine, token.StartColumn, p)
	}
	fmt.Println()

	// Return the list of tokens
	return tokenizer.tokens
}

package main

import (
	"fmt"
	"strconv"
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

// Peek at the current rune in the input without advancing. If it is the end of
// input, return 0 for the rune and 0 for the size. Otherwise, return the rune.
func (t *Tokenizer) peek() (rune, bool) {
	if t.pos >= len(t.input) {
		return rune(0), false // End of input
	}
	r, b := utf8.DecodeRuneInString(t.input[t.pos:])
	return r, b > 0
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
		if r == '"' || r == '\'' || r == '`' {
			_, ok := t.tryPeekTripleQuotes()
			if ok {
				t.readMultilineString()
				continue
			}
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

		// Match tokens starting with backslash (`\`)
		if r == '\\' {
			// Look ahead to check for a quote (`, ', or `)
			secondRune, ok := t.peekN(2)
			if ok && (secondRune == '"' || secondRune == '\'' || secondRune == '`') {
				t.consume()       // Consume the backslash
				t.readRawString() // Process as a raw string
			} else {
				// Consume and discard unexpected backslashes or handle other cases here
				t.consume()
			}
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

func (t *Tokenizer) tryPeekTripleQuotes() (rune, bool) {
	// Peek the first character to check if it’s a valid quote character
	r1, ok1 := t.peek()
	if !ok1 || (r1 != '\'' && r1 != '"' && r1 != '`') {
		return 0, false // Invalid or non-quote character
	}

	// Check if the next two characters match the first one
	r2, ok2 := t.peekN(2)
	r3, ok3 := t.peekN(3)
	if !(ok2 && ok3 && r2 == r1 && r3 == r1) {
		return 0, false // Not a triple quote
	}

	return r1, true // Successfully read triple quotes
}

func (t *Tokenizer) tryReadTripleQuotes() (rune, bool) {
	r, b := t.tryPeekTripleQuotes()

	if b {
		// Consume all three quotes
		t.consume() // Consume the first quote
		t.consume() // Consume the second quote
		t.consume() // Consume the third quote
	}

	return r, b
}

func (t *Tokenizer) tryReadMatchingTripleQuotes(q rune) bool {
	if t.tryPeekMatchingTripleQuotes(q) {
		// Consume all three quotes
		t.consume() // Consume the first quote
		t.consume() // Consume the second quote
		t.consume() // Consume the third quote
		return true
	}
	return false
}

func (t *Tokenizer) tryPeekMatchingTripleQuotes(q rune) bool {
	r, b := t.tryPeekTripleQuotes()
	return b && r == q
}

// Method to ensure there are no non-whitespace characters on the same line
func (t *Tokenizer) ensureOnlyTripleQuotesOnLine() {
	// Check for non-whitespace characters on the same line
	for t.pos < len(t.input) {
		r, _ := t.peek()
		if r == '\n' { // End of line
			t.consume()
			break
		}
		if r == '\r' { // Handle \r\n line endings
			t.consume() // Consume \r
			if t.pos < len(t.input) && t.input[t.pos] == '\n' {
				t.consume() // Consume \n
			}
			break
		}
		if !unicode.IsSpace(r) {
			// Panic if any non-space character is found
			panic(fmt.Sprintf("Opening triple quote must be on its own line (line %d, column %d)", t.lineNo, t.colNo))
		}
		t.consume() // Consume the current character
	}
}

func (t *Tokenizer) readMultilineString(rawFlag bool) {
	startLine, startCol := t.lineNo, t.colNo

	// Validate and consume the opening triple quotes
	openingQuote, ok := t.tryReadTripleQuotes()
	if !ok {
		panic(fmt.Sprintf("Malformed opening triple quotes at line %d, column %d", startLine, startCol))
	}

	// Ensure no other non-space characters appear on the opening line
	t.ensureOnlyTripleQuotesOnLine()

	// Buffer to temporarily hold each line
	var lines []string
	done := false

	for t.pos < len(t.input) && !done {
		// Read the current line
		start := t.pos

		for t.pos < len(t.input) && t.input[t.pos] != '\n' && t.input[t.pos] != '\r' {
			if t.tryPeekMatchingTripleQuotes(openingQuote) {
				done = true
				break
			}
			t.consume()
		}
		line := t.input[start:t.pos]

		// Consume the newline using the helper (if it exists)
		t.consumeNewline()

		// Process the line based on the `rawFlag`
		if !rawFlag {
			var processedLine strings.Builder
			for _, r := range line {
				if r == '\\' {
					processedLine.WriteString(handleEscapeSequence(t, openingQuote))
				} else {
					processedLine.WriteRune(r)
				}
			}
			lines = append(lines, processedLine.String())
		} else {
			lines = append(lines, line)
		}
	}

	// Consume the closing triple quotes
	if !t.tryReadMatchingTripleQuotes(openingQuote) {
		panic(fmt.Sprintf("Closing triple quote not found (line %d, column %d)", t.lineNo, t.colNo))
	}

	// Verify that the last line consists only of whitespace
	if len(lines) > 0 {
		lastLine := lines[len(lines)-1]
		if strings.TrimSpace(lastLine) != "" {
			panic(fmt.Sprintf("Closing triple quote must be on its own line (line %d, column %d)", t.lineNo, t.colNo))
		}
	}

	// Validate and process each line based on closing indent
	closingIndent := lines[len(lines)-1]
	var text strings.Builder
	for i, line := range lines[:len(lines)-1] {
		processedLine := processLineWithIndent(line, closingIndent, startLine+i, t.lineNo, t.colNo)
		text.WriteString(processedLine)
	}

	// Add the multiline string token
	t.addToken(Literal, LiteralString, text.String(), startLine, startCol)
}

func processLineWithIndent(line string, closingIndent string, lineNumber int, closingLine int, closingCol int) string {
	// Allow empty lines (return as-is)
	if strings.TrimSpace(line) == "" {
		return "\n"
	}

	// Check if the line starts with the closing indent
	if !strings.HasPrefix(line, closingIndent) {
		panic(fmt.Sprintf(
			"Line %d does not start with the required closing indent at line %d, column %d",
			lineNumber, closingLine, closingCol,
		))
	}

	// Remove the closing indent from the line and return the processed result
	return line[len(closingIndent):] + "\n"
}

func (t *Tokenizer) consumeNewline() {
	// Consume '\r' and optionally '\n' to handle both '\n' and '\r\n' line endings
	if t.pos < len(t.input) && t.input[t.pos] == '\r' {
		t.consume() // Consume '\r'
		if t.pos < len(t.input) && t.input[t.pos] == '\n' {
			t.consume() // Consume '\n' if it follows
		}
	} else if t.pos < len(t.input) && t.input[t.pos] == '\n' {
		t.consume() // Consume '\n'
	}
}

func (t *Tokenizer) readRawString() {
	startLine, startCol := t.lineNo, t.colNo
	quote := t.consume() // Consume the opening quote
	var text strings.Builder

	for t.pos < len(t.input) {
		r := t.consume()
		if r == quote { // Closing quote found
			break
		}
		// Backslashes are treated as normal characters in raw strings
		text.WriteRune(r)
	}

	// Add the raw string token
	t.addToken(Literal, LiteralString, text.String(), startLine, startCol)
}

func (t *Tokenizer) readString() {
	startLine, startCol := t.lineNo, t.colNo
	quote := t.consume() // Consume the opening quote
	var text strings.Builder

	for t.pos < len(t.input) {
		r := t.consume()
		if r == quote { // Closing quote found
			break
		}
		if r == '\\' && t.pos < len(t.input) { // Handle escape sequences
			text.WriteString(handleEscapeSequence(t, quote))
		} else {
			text.WriteRune(r)
		}
	}

	// Add the string token
	t.addToken(Literal, LiteralString, text.String(), startLine, startCol)
}

// Helper method to process escape sequences
func handleEscapeSequence(t *Tokenizer, quote rune) string {
	var text strings.Builder
	r := t.consume() // Consume the escape character

	switch r {
	case 'b':
		text.WriteRune('\b')
	case 'f':
		text.WriteRune('\f')
	case 'n':
		text.WriteRune('\n')
	case 'r':
		text.WriteRune('\r')
	case 't':
		text.WriteRune('\t')
	case '\\', '/', quote: // Escaped backslash, slash, or matching quote
		text.WriteRune(r)
	case 'u': // Unicode escape sequence
		if t.pos+4 <= len(t.input) { // Ensure there are enough characters
			code := t.input[t.pos : t.pos+4]
			t.pos += 4 // Consume 4 characters
			if decoded, err := decodeUnicodeEscape(code); err == nil {
				text.WriteRune(decoded)
			} else {
				text.WriteString(`\u` + code) // Keep invalid escape sequences intact
			}
		} else {
			text.WriteString(`\u`) // Handle incomplete Unicode sequence
		}
	case '_': // Non-standard escape sequence: \_
		// Expand into no characters (do nothing)
		// This has a couple of use-cases. 1. It helps break up a dense sequence
		// of characters, making it easier to read. 2. It can be used to introduce
		// a non-standard identifier.
	default:
		text.WriteRune('\\') // Keep invalid escape sequences as-is
		text.WriteRune(r)
	}

	return text.String()
}

// Decode a Unicode escape sequence (\uXXXX) into a rune
func decodeUnicodeEscape(code string) (rune, error) {
	if r, err := strconv.ParseInt(code, 16, 32); err == nil {
		return rune(r), nil
	} else {
		return 0, err
	}
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

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
	OpenBracket
	CloseBracket
	Sign
)

// Subtypes for Literal
const (
	LiteralString uint8 = iota
	LiteralNumber
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
	StartLine            int       // The starting line number of the token
	StartColumn          int       // The starting column number of the token
	PrecededByNewline    bool      // New field to indicate if the token is preceded by a newline
	FollowedByWhitespace bool      // New field to indicate if the token is followed by whitespace
	EscapeSeen           bool      // New field to indicate if an escape sequence was seen
	IsMultiLine          bool      // New field to indicate if the token is a multi-line string
	QuoteRune            rune      // New field to indicate the quote rune for strings
	NextToken            *Token    // The next token in the chain

	// Cache for precedence
	precValue int  // Cached precedence value
	precValid bool // Indicates if the precedence has been computed
	errFlag   bool // Cached error flag for precedence validity
}

type Tokenizer struct {
	input       string   // The input string to tokenize
	tokens      []*Token // The array of tokens generated
	lineNo      int      // Current line number
	colNo       int      // Current column number
	pos         int      // Current byte position in the input
	NewlineSeen bool     // New field to indicate if a newline has been seen
}

// Create a new Tokenizer
func NewTokenizer(input string) *Tokenizer {
	return &Tokenizer{
		input:  input,
		tokens: []*Token{},
		lineNo: 1,
		colNo:  1,
		pos:    0,
	}
}

func (t *Token) IsBreaker() bool {
	return t.IsSimpleBreaker() || t.IsCompoundBreaker()
}

func (t *Token) IsSimpleBreaker() bool {
	if t.Type != Identifier || t.SubType != IdentifierVariable {
		return false
	}
	if t.FollowedByWhitespace {
		return false
	}
	if t.NextToken.Type != Sign || t.NextToken.SubType != SignLabel {
		return false
	}
	return true
}

func (t *Token) IsCompoundBreaker() bool {
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
	return true
}

func (t *Token) IsMacro() bool {
	if t.Type != Identifier || t.SubType == IdentifierFormEnd {
		return false
	}
	t1 := t.NextToken
	if t1 == nil {
		return false
	}
	return t1.Type == Sign || t1.SubType == SignForce
}

func (t *Token) SetSeen(seen bool) {
	t.PrecededByNewline = seen
}

const signChars = ".*/%+-<~!&|?:="
const precCharacters = ".({[*/%+-<~!&|?:="

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
	prefixPrecedence     = 1
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

// hasMoreInput checks whether there is any remaining input to be processed.
// It returns true if the current position has not reached the end of the input
// string, indicating that there is more content to tokenize.
func (t *Tokenizer) hasMoreInput() bool {
	return t.pos < len(t.input)
}

func (t *Tokenizer) peekN(n int) (rune, bool) {
	currentPos := t.pos // Start at the current position
	var r rune
	var size int

	// Iterate through the input to locate the nth rune
	for range n {
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
	if r == '\n' {
		t.NewlineSeen = true
	} else if !unicode.IsSpace(r) {
		t.NewlineSeen = false
	}
	return r
}

// Add a token to the token list
func (t *Tokenizer) addToken(tokenType TokenType, subType uint8, text string, startLine int, startCol int) *Token {
	token := Token{
		Type:        tokenType,
		SubType:     subType,
		Text:        text,
		StartLine:   startLine,
		StartColumn: startCol,
	}
	t.tokens = append(t.tokens, &token)
	return &token
}

func (t *Tokenizer) tokenize() {
	for t.hasMoreInput() {
		r, _ := t.peek()
		seen := t.NewlineSeen

		// Skip whitespace
		if unicode.IsSpace(r) {
			t.consume()
			continue
		}

		// Skip comments
		if r == '#' {
			for t.hasMoreInput() {
				r, _ := t.peek()
				if r == '\n' || r == '\r' {
					t.consumeNewline()
					break
				}
				t.consume()
			}
			continue
		}

		// Match strings
		if r == '"' || r == '\'' || r == '`' {
			_, ok := t.tryPeekTripleQuotes()
			if ok {
				t.readMultilineString(false)
				continue
			}
			t.readString().SetSeen(seen)
			continue
		}

		// Match numbers
		if unicode.IsDigit(r) {
			t.readNumber().SetSeen(seen)
			continue
		}

		// Match identifiers
		if unicode.IsLetter(r) || r == '_' {
			t.readIdentifier().SetSeen(seen)
			continue
		}

		// Match punctuation
		if r == ',' || r == ';' {
			t.readPunctuation().SetSeen(seen)
			continue
		}

		// Match brackets
		if r == '(' || r == ')' || r == '[' || r == ']' || r == '{' || r == '}' {
			t.readBracket().SetSeen(seen)
			continue
		}

		// Match signs
		if t.isSign(r) {
			t.readSign().SetSeen(seen)
			continue
		}

		// Match tokens starting with backslash (`\`)
		if r == '\\' {
			// Look ahead to check for a quote
			secondRune, ok := t.peekN(2)
			if ok && (secondRune == '"' || secondRune == '\'' || secondRune == '`') {
				if _, ok := t.tryPeekTripleQuotes(); ok {
					t.readMultilineString(true).SetSeen(seen)
				} else {
					t.consume()                     // Consume the backslash
					t.readRawString().SetSeen(seen) // Process as a raw string
				}
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
	return strings.ContainsRune(signChars, r)
}

func (t *Tokenizer) readSign() *Token {
	startLine, startCol := t.lineNo, t.colNo
	start := t.pos

	for t.hasMoreInput() {
		r, _ := t.peek()
		if !t.isSign(r) {
			break
		}
		t.consume()
	}

	// Add the sign token
	text := t.input[start:t.pos]
	subType := SignOperator
	if text == "." {
		subType = SignDot
	} else if text == "!" {
		subType = SignForce
	} else if text == ":" {
		subType = SignLabel
	} else if text == "-" {
		subType = SignMinus
	}
	return t.addToken(Sign, subType, text, startLine, startCol) // 0 for now as signs may not have subtypes yet
}

func (t *Tokenizer) readBracket() *Token {
	startLine, startCol := t.lineNo, t.colNo
	r := t.consume() // Consume the bracket character

	// Determine the subtype
	var subType uint8
	var ttype TokenType
	switch r {
	case '(':
		ttype = OpenBracket
		subType = BracketParenthesis
	case ')':
		ttype = CloseBracket
		subType = BracketParenthesis
	case '[':
		ttype = OpenBracket
		subType = BracketBracket
	case ']':
		ttype = CloseBracket
		subType = BracketBracket
	case '{':
		ttype = OpenBracket
		subType = BracketBrace
	case '}':
		ttype = CloseBracket
		subType = BracketBrace
	}

	// Add the bracket token
	return t.addToken(ttype, subType, string(r), startLine, startCol)
}

func (t *Tokenizer) readPunctuation() *Token {
	startLine, startCol := t.lineNo, t.colNo
	r := t.consume() // Consume the punctuation character

	// Determine the subtype
	var subType uint8
	if r == ',' {
		subType = PunctuationComma
	} else if r == ';' {
		subType = PunctuationSemicolon
	}

	// Add the punctuation token
	return t.addToken(Punctuation, subType, string(r), startLine, startCol)
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
	for t.hasMoreInput() {
		r, _ := t.peek()
		if r == '\n' { // End of line
			t.consume()
			break
		}
		if r == '\r' { // Handle \r\n line endings
			t.consume() // Consume \r
			if t.hasMoreInput() && t.input[t.pos] == '\n' {
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

func (t *Tokenizer) readMultilineString(rawFlag bool) *Token {
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

	for t.hasMoreInput() && !done {
		// Read the current line
		var text strings.Builder

		for t.hasMoreInput() && t.input[t.pos] != '\n' && t.input[t.pos] != '\r' {
			if t.tryPeekMatchingTripleQuotes(openingQuote) {
				done = true
				break
			}
			if !rawFlag && t.input[t.pos] == '\\' {
				t.consume()
				text.WriteString(handleEscapeSequence(t))
				continue
			}
			text.WriteRune(t.consume())
		}

		// Consume the newline using the helper (if it exists)
		t.consumeNewline()

		lines = append(lines, text.String())
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
	token := t.addToken(Literal, LiteralString, text.String(), startLine, startCol)
	token.IsMultiLine = true
	token.QuoteRune = openingQuote
	return token
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
	if t.hasMoreInput() && t.input[t.pos] == '\r' {
		t.consume() // Consume '\r'
		if t.hasMoreInput() && t.input[t.pos] == '\n' {
			t.consume() // Consume '\n' if it follows
		}
	} else if t.hasMoreInput() && t.input[t.pos] == '\n' {
		t.consume() // Consume '\n'
	}
}

func (t *Tokenizer) readRawString() *Token {
	startLine, startCol := t.lineNo, t.colNo
	quote := t.consume() // Consume the opening quote
	var text strings.Builder

	for t.hasMoreInput() {
		r := t.consume()
		if r == quote { // Closing quote found
			break
		}
		// Backslashes are treated as normal characters in raw strings
		text.WriteRune(r)
	}

	// Add the raw string token
	token := t.addToken(Literal, LiteralString, text.String(), startLine, startCol)
	token.QuoteRune = quote
	return token
}

func (t *Tokenizer) readString() *Token {
	startLine, startCol := t.lineNo, t.colNo
	quote := t.consume() // Consume the opening quote
	var text strings.Builder

	for t.hasMoreInput() {
		r := t.consume()
		if r == quote { // Closing quote found
			break
		}
		if r == '\\' && t.hasMoreInput() { // Handle escape sequences
			text.WriteString(handleEscapeSequence(t))
		} else {
			text.WriteRune(r)
		}
	}

	// Add the string token
	token := t.addToken(Literal, LiteralString, text.String(), startLine, startCol)
	token.QuoteRune = quote
	return token
}

// Helper method to process escape sequences
func handleEscapeSequence(t *Tokenizer) string {
	var text strings.Builder
	r := t.consume() // Consume the escape character

	fmt.Println("esc", r)

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
	case '\\', '/', '"', '\'', '`': // Escaped backslash, slash, or matching quote
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
		fmt.Println("Underscore")
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

func (t *Tokenizer) readNumber() *Token {
	startLine, startCol := t.lineNo, t.colNo
	start := t.pos
	hasDot := false

	for t.hasMoreInput() {
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
	token := t.addToken(Literal, LiteralNumber, text, startLine, startCol)
	return token
}

func (t *Tokenizer) readIdentifier() *Token {
	startLine, startCol := t.lineNo, t.colNo
	var text strings.Builder
	var escSeen bool = false

	for t.hasMoreInput() {
		r, _ := t.peek()

		if !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_') {
			// Handle escape sequences when a backslash is encountered
			if r == '\\' {
				t.consume() // Consume the backslash
				escapeText := handleEscapeSequence(t)
				text.WriteString(escapeText)
				escSeen = true
				continue
			}
			break
		}

		t.consume() // Consume the current character
		text.WriteRune(r)
	}

	// Peek at the next character after the identifier to check for whitespace
	r, ok := t.peek()
	followedByWhitespace := ok && unicode.IsSpace(r)

	// Add the identifier token with the new field
	token := t.addToken(Identifier, IdentifierVariable, text.String(), startLine, startCol)
	token.FollowedByWhitespace = followedByWhitespace
	token.EscapeSeen = escSeen
	return token
}

func (t *Tokenizer) markReservedTokens() {
	// This function marks tokens that act as reserved words. To do this we
	// collect all identifiers in a set `idents``. Then we iterate through the
	// tokens again. For each identifier token T we check if it is one of a
	// pair "XXX" and "endXXX", where XXX is a common string.
	idents := make(map[string]*Token)
	for _, token := range t.tokens {
		if token.Type == Identifier {
			idents[token.Text] = token
		}
	}
	for _, token := range t.tokens {
		fmt.Println("Token", token.Text, "HasPrefix", strings.HasPrefix(token.Text, "end"))
		if token.Type == Identifier && strings.HasPrefix(token.Text, "end") && idents[token.Text] != nil {
			start := idents[token.Text[3:]]
			if start != nil {
				start.SubType = IdentifierFormStart
				token.SubType = IdentifierFormEnd
			}
		}
	}
}

func (t *Tokenizer) chainTokens() {
	for n, token := range t.tokens {
		if n == 0 {
			continue
		}
		prev := t.tokens[n-1]
		prev.NextToken = token
	}
}

func tokenizeInput(input string) []*Token {
	// Create a new Tokenizer instance
	tokenizer := NewTokenizer(input)

	fmt.Println("Discovering Tokens ...")

	// Perform tokenization
	tokenizer.tokenize()
	tokenizer.markReservedTokens()
	tokenizer.chainTokens()

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

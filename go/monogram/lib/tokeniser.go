package lib

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type TokenizerError struct {
	Message string `json:"message"`
	Line    int    `json:"line"`
	Column  int    `json:"column"`
}

type Tokenizer struct {
	input        string   // The input string to tokenize
	tokens       []*Token // The array of tokens generated
	lineNo       int      // Current line number
	colNo        int      // Current column number
	pos          int      // Current byte position in the input
	NewlineSeen  bool     // New field to indicate if a newline has been seen
	markStack    []int    // Flag to indicate if the position should be marked
	lineNoStack  []int    // Array to store line numbers for each token
	lineColStack []int    // Array to store column numbers for each token
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

func (t *Tokenizer) markPosition() {
	// Mark the current position in the input
	t.markStack = append(t.markStack, t.pos)
	t.lineNoStack = append(t.lineNoStack, t.lineNo)
	t.lineColStack = append(t.lineColStack, t.colNo)
}

// Reset the position to the last marked position
func (t *Tokenizer) resetPosition() {
	// Reset the position to the last marked position
	if len(t.markStack) == 0 {
		return
	}
	t.pos = t.markStack[len(t.markStack)-1]
	t.lineNo = t.lineNoStack[len(t.lineNoStack)-1]
	t.colNo = t.lineColStack[len(t.lineColStack)-1]
	t.markStack = t.markStack[:len(t.markStack)-1]
	t.lineNoStack = t.lineNoStack[:len(t.lineNoStack)-1]
	t.lineColStack = t.lineColStack[:len(t.lineColStack)-1]
}

func (t *Tokenizer) popMark() string {
	// Pop the last marked position and return the corresponding substring
	if len(t.markStack) == 0 {
		return ""
	}
	start := t.markStack[len(t.markStack)-1]
	t.markStack = t.markStack[:len(t.markStack)-1]
	t.lineNoStack = t.lineNoStack[:len(t.lineNoStack)-1]
	t.lineColStack = t.lineColStack[:len(t.lineColStack)-1]
	return t.input[start:t.pos]
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

func (t *Tokenizer) consumeN(n int) {
	// Consume n runes from the input
	for range n {
		if t.hasMoreInput() {
			t.consume()
		} else {
			break // Stop if we reach the end of input
		}
	}
}

// Add a token to the token list
func (t *Tokenizer) addToken(tokenType TokenType, subType uint8, text string, startLine int, startCol int) *Token {
	span := Span{startLine, startCol, -1, -1}
	token := t.newToken(tokenType, subType, text, span)
	t.appendToken(token)
	return token
}

// Pop the last token from the token list
func (t *Tokenizer) popToken() *Token {
	// Pop the last token from the token list
	if len(t.tokens) == 0 {
		return nil // No tokens to pop
	}
	token := t.tokens[len(t.tokens)-1]
	t.tokens = t.tokens[:len(t.tokens)-1] // Remove the last token
	return token
}

// Append a token to the token list
func (t *Tokenizer) appendToken(token *Token) {
	// Append the token to the token list
	t.tokens = append(t.tokens, token)
}

func (t *Tokenizer) newToken(tokenType TokenType, subType uint8, text string, span Span) *Token {
	// Create a new token with the current line and column numbers
	return &Token{
		Type:    tokenType,
		SubType: subType,
		Text:    text,
		Span:    span,
	}
}

func (t *Tokenizer) tokenize() *TokenizerError {
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
					t.tryConsumeNewline()
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
				token, terr := t.readMultilineString(false)
				if terr != nil {
					return terr
				}
				token.SetSeen(t, seen)
				continue
			}
			token, terr := t.readString(false, r)
			if terr != nil {
				return terr
			}
			token.SetSeen(t, seen)
			continue
		}

		// Match numbers
		if unicode.IsDigit(r) || (r == '-' && t.IsNegativeNumber()) {
			token, terr := t.readNumber()
			if terr != nil {
				return terr
			}
			token.SetSeen(t, seen)
			continue
		}

		// Match identifiers
		if unicode.IsLetter(r) || r == '_' {
			t.readIdentifier().SetSeen(t, seen)
			continue
		}

		// Match punctuation
		if r == ',' || r == ';' {
			t.readPunctuation().SetSeen(t, seen)
			continue
		}

		// Match brackets
		if r == '(' || r == ')' || r == '[' || r == ']' || r == '{' || r == '}' {
			t.readBracket().SetSeen(t, seen)
			continue
		}

		// Match signs
		if t.isSign(r) {
			t.readSign().SetSeen(t, seen)
			continue
		}

		// Match tokens starting with backslash (`\`)
		if r == '\\' {
			// Look ahead to check for a quote
			secondRune, ok := t.peekN(2)
			if ok && (secondRune == '"' || secondRune == '\'' || secondRune == '`') {
				t.consume() // Consume the backslash
				_, is_triple := t.tryPeekTripleQuotes()
				if is_triple {
					token, terr := t.readMultilineString(true)
					if terr != nil {
						return terr
					}
					token.SetSeen(t, seen)
				} else {
					token, terr := t.readRawString(false, secondRune)
					if terr != nil {
						return terr
					}
					token.SetSeen(t, seen) // Process as a raw string
				}
			} else {
				t.readIdentifier().SetSeen(t, seen)
			}
			continue
		}

		// Discard unexpected characters
		return &TokenizerError{Message: fmt.Sprintf("Unexpected character: %c", r), Line: t.lineNo, Column: t.colNo}
	}
	return nil
}

func (t *Tokenizer) IsNegativeNumber() bool {
	r, b := t.peekN(2)
	return b && unicode.IsDigit(r)
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

func (t *Tokenizer) consumeTripleQuotes(quote rune) *TokenizerError {
	r, b := t.tryReadTripleQuotes()
	if !b {
		return &TokenizerError{Message: "Missing triple quotes", Line: t.lineNo, Column: t.colNo}
	}
	if r != quote {
		return &TokenizerError{Message: fmt.Sprintf("Expected %c, but found %c", quote, r), Line: t.lineNo, Column: t.colNo}
	}
	return nil
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

func (t *Tokenizer) skipSpacesUpToNewline() {
	// Skip whitespace characters
	for t.hasMoreInput() {
		r, ok := t.peek()
		if !ok || r == '\n' || r == '\r' {
			break
		}
		if !unicode.IsSpace(r) {
			break // Stop if a non-whitespace character is found
		}
		t.consume() // Consume the whitespace character
	}
}

// Method to ensure there are no non-whitespace characters on the same line
func (t *Tokenizer) ensureRestOfLineIsWhitespace() *TokenizerError {
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
			return &TokenizerError{Message: "Opening triple quote must be on its own line", Line: t.lineNo, Column: t.colNo}
		}
		t.consume() // Consume the current character
	}
	return nil
}

// Calculates the closing indent if we are on the last line of a multiline string.
func textIsWhitespaceFollowedBy3Quotes(text string, quote rune) (bool, string) {
	// Check if the text is whitespace followed by three quotes
	if len(text) < 3 {
		return false, ""
	}
	var indent strings.Builder
	whitespace := true
	count := 0
	for _, r := range text {
		if whitespace {
			if unicode.IsSpace(r) {
				indent.WriteRune(r)
				continue
			} else if r == quote {
				whitespace = false
				count++
			} else {
				return false, ""
			}
		} else {
			if r == quote {
				count++
				if count >= 3 {
					return true, indent.String()
				}
			} else {
				return false, ""
			}
		}
	}
	return false, ""
}

func (t *Tokenizer) findClosingIndent() (rune, string, int, *TokenizerError) {
	t.markPosition()

	// Validate and consume the opening triple quotes
	quote, ok := t.tryReadTripleQuotes()
	if !ok {
		return 0, "", 0, &TokenizerError{Message: "Malformed opening triple quotes", Line: t.lineNo, Column: t.colNo}
	}

	// Ensure no other non-space characters appear on the opening line
	terr := t.ensureRestOfLineIsWhitespace()
	if terr != nil {
		return 0, "", 0, terr
	}

	// Now read each line in order until we find the closing line.
	startLine, startCol := t.lineNo, t.colNo
	lines := []string{}
	var match bool
	var closingIndent string
	for t.hasMoreInput() {
		line := t.readRestOfLine()
		match, closingIndent = textIsWhitespaceFollowedBy3Quotes(line, quote)
		if match {
			break
		}
		lines = append(lines, line)
	}

	if !match {
		return 0, "", 0, &TokenizerError{Message: "Closing triple quote not found", Line: t.lineNo, Column: t.colNo}
	}

	for i, line := range lines {
		// Allow empty lines
		if line == "" {
			continue
		}
		// Check if the line starts with the closing indent
		if !strings.HasPrefix(line, closingIndent) {
			return 0, "", 0, &TokenizerError{
				Message: "not indented consistently with the closing triple quote",
				Line:    startLine + i,
				Column:  startCol,
			}
		}
	}

	t.resetPosition()
	return quote, closingIndent, len(lines), nil
}

func (t *Tokenizer) readRestOfLine() string {
	// Read the rest of the line until a newline or end of input
	var text strings.Builder
	for t.hasMoreInput() {
		r, _ := t.peek()
		if r == '\n' || r == '\r' {
			break // End of line
		}
		text.WriteRune(t.consume())
	}
	t.tryConsumeNewline()
	return text.String()
}

func (t *Tokenizer) readMultilineString(rawFlag bool) (*Token, *TokenizerError) {

	startLine, startCol := t.lineNo, t.colNo
	var subTokens []*Token

	openingQuote, closingIndent, nlines, terr := t.findClosingIndent()
	if terr != nil {
		return nil, terr
	}

	// Discard the rest of this line, which are the opening quotes.
	t.readRestOfLine()

	// The next N lines should be either all whitespace or start with the
	// closing indent.
	for range nlines {
		if t.tryConsumeText(closingIndent) {
			if rawFlag {
				t.readRawString(true, openingQuote)
			} else {
				t.readString(true, openingQuote)
			}
		} else {
			token := t.addToken(Literal, LiteralString, "", t.lineNo, t.colNo)
			token.QuoteRune = openingQuote
		}
		subTokens = append(subTokens, t.popToken())
	}

	// Discard the rest of the next line, which will be the closing quotes.
	t.skipSpacesUpToNewline()
	t.consumeTripleQuotes(openingQuote)

	// Add the multiline string token
	token := t.addToken(Literal, LiteralMultilineString, "", startLine, startCol)
	token.IsMultiLine = true
	token.QuoteRune = openingQuote
	token.SubTokens = subTokens

	return token, nil
}

func (t *Tokenizer) tryConsumeRune(char rune) bool {
	// Check if the next rune matches the given character
	r, ok := t.peek()
	if !ok {
		return false // End of input
	}
	if r != char {
		return false // No match
	}
	t.consume() // Consume the character
	return true
}

func (t *Tokenizer) tryConsumeNewline() bool {
	// Consume '\r' and optionally '\n' to handle both '\n' and '\r\n' line endings
	if t.hasMoreInput() && t.input[t.pos] == '\r' {
		t.consume() // Consume '\r'
		if t.hasMoreInput() && t.input[t.pos] == '\n' {
			t.consume() // Consume '\n' if it follows
		}
		return true
	} else if t.hasMoreInput() && t.input[t.pos] == '\n' {
		t.consume() // Consume '\n'
		return true
	}
	return false // No newline consumed
}

func (t *Tokenizer) tryConsumeText(text string) bool {
	// Check if the next characters match the given text
	if strings.HasPrefix(t.input[t.pos:], text) {
		t.consumeN(len(text)) // Consume the matching text
		return true
	}
	return false
}

func (t *Tokenizer) readRawString(unquoted bool, default_quote rune) (*Token, *TokenizerError) {
	startLine, startCol := t.lineNo, t.colNo
	quote := default_quote
	if !unquoted {
		quote = t.consume() // Consume the opening quote
	}
	var text strings.Builder

	for {
		if !t.hasMoreInput() {
			return nil, &TokenizerError{Message: "Unterminated raw string", Line: startLine, Column: startCol}
		}
		r := t.consume()
		if r == quote { // Closing quote found
			break
		} else if r == '\n' || r == '\r' { // Handle newlines
			if unquoted {
				if r == '\r' {
					t.tryConsumeRune('\n') // Consume '\n' if it follows
				}
				break
			}
			return nil, &TokenizerError{Message: "Line break in raw string", Line: startLine, Column: startCol}
		}
		// Backslashes are treated as normal characters in raw strings
		text.WriteRune(r)
	}

	// Add the raw string token
	token := t.addToken(Literal, LiteralString, text.String(), startLine, startCol)
	token.QuoteRune = quote
	return token, nil
}

func (t *Tokenizer) readString(unquoted bool, default_quote rune) (*Token, *TokenizerError) {
	startLine, startCol := t.lineNo, t.colNo
	currSpan := Span{startLine, startCol, -1, -1}
	quote := default_quote
	if !unquoted {
		quote = t.consume() // Consume the opening quote
	}
	var text strings.Builder
	var interpolationTokens []*Token

	for {
		if !t.hasMoreInput() {
			return nil, &TokenizerError{Message: "Unterminated string", Line: startLine, Column: startCol}
		}
		r := t.consume()
		if !unquoted && r == quote { // Closing quote found
			break
		}
		if r == '\\' && t.hasMoreInput() { // Handle escape or interpolation
			next, _ := t.peek()
			if next == '(' || next == '[' || next == '{' {
				// End the current StringToken and handle interpolation
				if text.Len() > 0 {
					currSpan.EndLine, currSpan.EndColumn = t.lineNo, t.colNo-1 // Do not include the backslash
					current := t.newToken(Literal, LiteralString, text.String(), currSpan)
					current.QuoteRune = quote
					interpolationTokens = append(interpolationTokens, current)
					text.Reset()
				}
				interpolatedToken, err := t.readStringInterpolation()
				if err != nil {
					return nil, err
				}
				interpolationTokens = append(interpolationTokens, interpolatedToken)
				currSpan = Span{t.lineNo, t.colNo, -1, -1}
			} else {
				text.WriteString(handleEscapeSequence(t))
			}
		} else if r == '\n' || r == '\r' { // Handle newlines
			if unquoted {
				if r == '\r' {
					t.tryConsumeRune('\n') // Consume '\n' if it follows
				}
				break
			}
			return nil, &TokenizerError{Message: "Line break in string", Line: startLine, Column: startCol}
		} else {
			text.WriteRune(r)
		}
	}

	// Add the final StringToken if there's remaining text
	if text.Len() > 0 {
		currSpan.EndLine, currSpan.EndColumn = t.lineNo, t.colNo
		token := t.newToken(Literal, LiteralString, text.String(), currSpan)
		token.QuoteRune = quote
		interpolationTokens = append(interpolationTokens, token)
	}

	// Is this just a literal string?
	if len(interpolationTokens) == 1 && interpolationTokens[0].SubType == LiteralString {
		t.appendToken(interpolationTokens[0])
		return interpolationTokens[0], nil
	}

	// Combine into a StringInterpolationToken if interpolation occurred
	compoundToken := t.addToken(Literal, LiteralInterpolatedString, "", startLine, startCol)
	compoundToken.QuoteRune = quote
	compoundToken.SubTokens = interpolationTokens
	return compoundToken, nil
}

func (t *Tokenizer) readStringInterpolation() (*Token, *TokenizerError) {
	span := Span{t.lineNo, t.colNo, -1, -1}
	state := 0       // State 0: inside expression, State 1: inside string
	var stack []rune // Pushdown stack

	t.markPosition()                   // Mark the position for the interpolation
	openingRune := t.consume()         // Consume the opening bracket
	stack = append(stack, openingRune) // Push opening bracket onto stack

	for {
		if !t.hasMoreInput() {
			return nil, &TokenizerError{Message: "Unterminated interpolation", Line: span.StartLine, Column: span.StartColumn}
		}
		r := t.consume()
		switch state {
		case 0: // Inside expression
			switch r {
			case '\\': // Escape sequence
				handleEscapeSequence(t)
			case '(', '[', '{': // Opening brackets
				stack = append(stack, r)
			case ')', ']', '}': // Closing brackets
				if len(stack) > 0 && matches(stack[len(stack)-1], r) {
					stack = stack[:len(stack)-1] // Pop stack
					if len(stack) == 0 {         // End of interpolation
						text := t.popMark() // Pop the marked position
						span.EndLine, span.EndColumn = t.lineNo, t.colNo
						token := t.newToken(Literal, LiteralExpressionString, text, span)
						return token, nil
					}
				} else {
					return nil, &TokenizerError{Message: "Mismatched bracket", Line: span.StartLine, Column: span.StartColumn}
				}
			case '"', '\'', '`': // Enter string state
				stack = append(stack, r)
				state = 1
			case 'r', '\n': // Line breaks are not allowed
				return nil, &TokenizerError{Message: "Line break in interpolation", Line: t.lineNo, Column: t.colNo}
			}
		case 1: // Inside string
			switch r {
			case '\\': // Escape sequence
				if t.hasMoreInput() {
					next, _ := t.peek()
					if next == '(' || next == '[' || next == '{' {
						_, err := t.readStringInterpolation()
						if err != nil {
							return nil, err
						}
					} else {
						handleEscapeSequence(t)
					}
				} else {
					return nil, &TokenizerError{Message: "Unterminated escape sequence", Line: span.StartLine, Column: span.StartColumn}
				}
			case stack[len(stack)-1]: // Matching closing quote
				stack = stack[:len(stack)-1] // Pop stack
				state = 0
			}
		}
	}
}

// Helper to check if brackets match
func matches(open, close rune) bool {
	return (open == '(' && close == ')') || (open == '[' && close == ']') || (open == '{' && close == '}')
}

// Helper method to process escape sequences
func handleEscapeSequence(t *Tokenizer) string {
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

func (t *Tokenizer) readNumber() (*Token, *TokenizerError) {
	startLine, startCol := t.lineNo, t.colNo
	start := t.pos

	// Initialize prev to zero.
	var prev rune = 0

	// Handle an optional leading '-' sign.
	if t.hasMoreInput() && t.tryConsumeRune('-') {
		prev = '-' // Assign immediately since this affects later parsing.
	}

	hasDot := false

	for t.hasMoreInput() {
		r, _ := t.peek()

		if r == '.' {
			// Reject if multiple dots or if the previous character was an underscore.
			if hasDot || prev == '_' {
				break
			}
			hasDot = true
			t.consume()
		} else if unicode.IsDigit(r) {
			t.consume()
		} else if r == '_' {
			// Allow underscores only if the previous character was a digit.
			if !unicode.IsDigit(prev) {
				// Invalid underscore placement
				break
			}
			// Use peekIf to verify that the following character is a digit.
			r2, b := t.peekN(2)
			if !b || !unicode.IsDigit(r2) {
				// Invalid underscore placement
				break
			}
			t.consume() // Consume the underscore
		} else {
			break
		}

		// Update prev at the end of each iteration.
		prev = r
	}

	// If no runes were consumed or the only rune consumed was a sign, return an error.
	if start == t.pos || (start == t.pos-1 && prev == '-') {
		return nil, &TokenizerError{Message: "Invalid number format", Line: startLine, Column: startCol}
	}

	text := t.input[start:t.pos]
	token := t.addToken(Literal, LiteralNumber, text, startLine, startCol)
	return token, nil
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

func (t *Tokenizer) markReservedTokens() *TokenizerError {
	ident_exists := make(map[string]bool)
	is_reserved := make(map[string]bool) // A subset of ident_exists
	for _, token := range t.tokens {
		if token.Type == Identifier {
			ident_exists[token.Text] = true
		}
	}
	for n, token := range t.tokens {
		if token.Type != Identifier {
			continue
		}
		var next *Token
		if n < len(t.tokens)-1 {
			next = t.tokens[n+1]
		}
		if next != nil && next.Type == Sign && next.SubType == SignForce {
			if strings.HasPrefix(token.Text, "end") {
				//return fmt.Errorf("cannot use %s as an opening keyword", token.Text)
				return &TokenizerError{
					Message: fmt.Sprintf("cannot use '%s' as an opening keyword", token.Text),
					Line:    token.Span.StartLine,
					Column:  token.Span.StartColumn,
				}
			}
			token.SubType = IdentifierFormStart
			is_reserved[token.Text] = true
		}
	}
	for _, token := range t.tokens {
		if token.Type != Identifier || strings.HasPrefix(token.Text, "endend") {
			continue
		}
		if strings.HasPrefix(token.Text, "end") {
			stem := token.Text[3:]
			if ident_exists[stem] {
				token.SubType = IdentifierFormEnd
			}
		} else if is_reserved[token.Text] {
			token.SubType = IdentifierFormStart
		} else {
			if ident_exists["end"+token.Text] {
				token.SubType = IdentifierFormStart
			}
		}
	}
	return nil
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

func tokenizeInput(input string, colOffset int) ([]*Token, Span, *TokenizerError) {
	// Create a new Tokenizer instance
	tokenizer := NewTokenizer(input)

	// Perform tokenization
	terr := tokenizer.tokenize()
	if terr != nil {
		return nil, Span{}, terr
	}

	terr = tokenizer.markReservedTokens()
	if terr != nil {
		return nil, Span{}, terr
	}
	tokenizer.chainTokens()
	if colOffset > 0 {
		for _, token := range tokenizer.tokens {
			token.Span.StartColumn += colOffset
			token.Span.EndColumn += colOffset
		}
	}

	// Return the list of tokens
	return tokenizer.tokens, Span{1, 1, tokenizer.lineNo, tokenizer.colNo}, nil
}

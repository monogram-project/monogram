We are collaborating on a Go program, codename `monogram`. The `monogram` tool
translates from `monogram` notation into XML, JSON and other formats. The 
notation is designed to represent program-like texts. However it is just a
notation and not a programming language, although it does have an opinionated
grammar. Consequently it has no built-in variables, no built-in operators and
even the reserved words are dynamically discovered during the parse.

The program has several phases, which we have completed:

- an initial ingestion phase in which the input is tokenised.
- a pass to discover and mark the identifiers that are used as keywords.
- a parsing of the tokens to form an internal AST.
- walking the AST to generate output.

We are now working on an extension to tokenisation that will support
string interpolation. Firstly we will create two new token types:

- ExpressionStringToken, which is a string-like token whose text value
  is marked as being a monogram expression.
- StringInterpolationToken, which is a compound token that consists of
  StringTokens and ExpressionStringTokens.

The strategy for tokenising a string becomes more complicated. As soon as we
detect a backslash followed by any one of `(`, `[` or `{` we complete the
in-progress StringToken and enter `readStringInterpolation`, passing that token
as a parameter. This constructs a ExpressionStringToken and pushes the 
token as the first sub-token. Then it enters a loop, checking to see if the
next characters are backslash followed by a bracket, which is guaranteed on
the first time around the loop, or a plain string section. If it is a backslash
followed by a bracket then it is the start of an interpolated section.

Here, the aim is to figure out the end of the interpolated section. To do this
we employ a finite-state-machine with a pushdown stack. The finite stack machine
has two states: 1 means that it is inside a string and 0 means that it is inside
a monogram expression. 

In state 0, we look for the following special characters: `\`, `()`, `[]`, `{}`,
`"`, `'` and backtick itself. Line breaks are not allowed in this section to
protect against runaway consumption. Any other character is simply consumed. A
`\` escapes the next character (or possibly several characters). An opening
bracket will push onto the pushdown stack. A closing bracket pops the pushdown
stack and checks for a match. An opening quote (single, double or backtick) will
push onto the stack and change to state 1. 

In state 1, we look for the following special characters: `\` and the 
matching closing quote (which will be the rune on top of the stack). Line
breaks are allowed in this state. If a matching closing quote is found then
we pop the stack and re-enter state 0. If a backslash is found then the next
rune (or several runes) are consumed. Any other character is simply consumed.

We stop when in state 0 we pop a matching bracket and find the state is empty.
At this point we have reached the end of the interpolated section.

The interpolated section is then pushed as a sub-token of type
StringInterpolationToken and the loop is reentered - unless we have come
to the end of the string.

Here is the current implementation of readString.

```go
func (t *Tokenizer) readString() (*Token, *TokenizerError) {
	startLine, startCol := t.lineNo, t.colNo
	quote := t.consume() // Consume the opening quote
	var text strings.Builder

	for {
		if !t.hasMoreInput() {
			return nil, &TokenizerError{Message: "Unterminated string", Line: startLine, Column: startCol}
		}
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
	return token, nil
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
```

Please rewrite readString as I have outlined.
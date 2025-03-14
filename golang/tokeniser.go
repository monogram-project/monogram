package main

type LexicalClass int

const (
	Identifier  LexicalClass = iota // 0
	Keyword                         // 1
	Literal                         // 2
	Operator                        // 3
	Punctuation                     // 4
	Comment                         // 5
	Whitespace                      // 6
)

// return the string representation of each lexical class for debugging or printing:
func (lc LexicalClass) String() string {
	switch lc {
	case Identifier:
		return "Identifier"
	case Keyword:
		return "Keyword"
	case Literal:
		return "Literal"
	case Operator:
		return "Operator"
	case Punctuation:
		return "Punctuation"
	case Comment:
		return "Comment"
	case Whitespace:
		return "Whitespace"
	default:
		return "Unknown"
	}
}

type Token struct {
	LexicalClass LexicalClass // The lexical class of the token
	Text         string       // The text representation of the token
	StartLine    int          // The starting line number of the token in the source
	StartColumn  int          // The starting column number of the token in the source
}

func tokenizeInput(input string) []Token {
	// Dummy implementation for now: produces a few example tokens
	return []Token{
		{
			LexicalClass: Keyword,
			Text:         "node1",
			StartLine:    1,
			StartColumn:  1,
		},
		{
			LexicalClass: Identifier,
			Text:         "key1",
			StartLine:    1,
			StartColumn:  7,
		},
		{
			LexicalClass: Literal,
			Text:         "value1",
			StartLine:    1,
			StartColumn:  12,
		},
	}
}

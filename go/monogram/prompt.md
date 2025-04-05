We are collaborating on a Go program, codename `monogram`. The `monogram` tool
translates from `monogram` notation into XML, JSON and other formats. The 
notation is designed to represent program-like texts. However it is just a
notation and not a programming language, although it does have an opinionated
grammar. Consequently it has no built-in variables, no built-in operators and
even the reserved words are dynamically discovered during the parse.

We have completed a good first version of that program. The repo that it resides in is a monorepo with a parallel implementation,
written in some other programming language (Pop-11). As a consequence the
go folder is not at the top-level of the repo but in `~/go/monogram/`. The
application itself is in `~/go/monogram/cmd/monogram/main.go`.

Much of the syntax is aligned with Python, since that is currently a very 
popular programming language. Our next task is to allow numbers to have 
underscores appear in them. At thee moment number syntax is more closely
modelled on JSON, which does not allow underscores. Here is the code for
reading numbers:

```go
func (t *Tokenizer) readNumber() *Token {
	startLine, startCol := t.lineNo, t.colNo
	start := t.pos

	if t.hasMoreInput() && t.input[t.pos] == '-' {
		t.consume() // Consume the negative sign
	}

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
```

Summarise the rules on underscores in numbers in Python. Then apply this to the
code.


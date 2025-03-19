package main

import (
	"fmt"
)

type Node struct {
	Name     string            // The name of the node
	Options  map[string]string // Attributes (name-value pairs)
	Children []*Node           // Child nodes
}

// Parser holds the list of tokens and our current reading position.
type Parser struct {
	tokens       []*Token
	pos          int
	UnglueOption *Token
}

type Context struct {
	InsideForm    bool
	AcceptNewline bool
}

// hasNext checks if there are tokens left to consume.
func (p *Parser) hasNext() bool {
	return p.pos < len(p.tokens)
}

// next returns the current token and advances our pointer.
func (p *Parser) next() *Token {
	tok := p.tokens[p.pos]
	p.pos++
	return tok
}

// peek returns the current token without advancing.
func (p *Parser) peek() *Token {
	if p.hasNext() {
		return p.tokens[p.pos]
	}
	return nil
}

func (p *Parser) readExpr(context Context) (*Node, error) {
	// fmt.Println(">>> READ EXPR")
	n, e := p.readExprPrec(maxPrecedence, context)
	// fmt.Println("<<< READ EXPR")
	return n, e
}

func (p *Parser) readArguments(subType uint8, context Context) (string, *Node, error) {
	// fmt.Println(">>> READ ARGUMENTS")
	c := context
	c.AcceptNewline = false
	sep, seq, err := p.readExprSeqTo(subType, true, c)
	if err != nil {
		return "", nil, err
	}
	node := &Node{
		Name:     "arguments",
		Children: seq,
	}
	// fmt.Println("<<< READ ARGUMENTS")
	return sep, node, nil
}

func (p *Parser) readOptExprPrec(formStart *Token, outer_prec int, context Context) (*Node, error) {
	if !p.hasNext() {
		return nil, nil
	}
	token := p.peek()
	// fmt.Println("Peeked token[4]: ", token.Text, token.Type, token.SubType)
	if token.Type == Punctuation {
		return nil, nil
	}
	if token.Type == CloseBracket {
		return nil, nil
	}
	if token.Type == Identifier && token.SubType == IdentifierFormEnd {
		return nil, nil
	}
	if token.Type == Identifier && token.SubType == IdentifierVariable && token.IsBreaker(formStart) {
		return nil, nil
	}
	if token.Type == Sign && token.SubType == SignLabel {
		return nil, nil
	}
	if context.AcceptNewline && token.PrecededByNewline {
		return nil, nil
	}
	return p.readExprPrec(outer_prec, context)
}

func (p *Parser) readExprPrec(outer_prec int, context Context) (*Node, error) {
	// fmt.Println(">>> READ EXPR PREC")
	lhs, err := p.readPrimaryExpr(context)
	if err != nil {
		return nil, err
	}
	for p.hasNext() {
		token1 := p.peek()
		// fmt.Println("Peeked token[3]: ", token1.Text, token1.Type, token1.SubType)
		if context.AcceptNewline && token1.PrecededByNewline {
			break
		}
		if context.InsideForm && token1.Type == Sign && token1.SubType == SignLabel {
			break
		}
		prec, ok := token1.Precedence()
		if !ok || prec > outer_prec {
			break
		}
		// fmt.Println("ok", ok, "Precedence", prec, "Outer precedence", outer_prec)
		token2 := p.next()
		c := context
		c.AcceptNewline = false
		if token2.Type == OpenBracket {
			// lvars (sep, args) = read_arguments( close_bracket );
			// lvars dname = delimiter_name( item1 );
			// consNode( "apply", $(kind=dname, separator=sep.sep_as_word), [^lhs ^args] ) -> lhs;
			sep_text, args, err := p.readArguments(token2.SubType, c)
			if err != nil {
				return nil, err
			}
			dname := token2.DelimiterName()
			lhs = &Node{
				Name: "apply",
				Options: map[string]string{
					"kind":      dname,
					"separator": sep_text,
				},
				Children: []*Node{lhs, args},
			}
		} else if token2.Type == Sign && token2.SubType == SignDot && p.hasNext() {
			property := p.next()
			if p.hasNext() && p.peek().Type == OpenBracket {
				token3 := p.next()
				sep_text, rhs, err := p.readArguments(token3.SubType, c)
				if err != nil {
					return nil, err
				}
				lhs = &Node{
					Name: "invoke",
					Options: map[string]string{
						"kind":      token3.DelimiterName(),
						"separator": sep_text,
						"name":      property.Text,
					},
					Children: []*Node{lhs, rhs},
				}
			} else {
				lhs = &Node{
					Name: "get",
					Options: map[string]string{
						"name": property.Text,
					},
					Children: []*Node{lhs},
				}
			}
		} else {
			rhs, err := p.readExprPrec(prec, c)
			if err != nil {
				return nil, err
			}
			lhs = &Node{
				Name: "operator",
				Options: map[string]string{
					"name":   token1.Text,
					"syntax": "infix",
				},
				Children: []*Node{lhs, rhs}, // lhs and rhs are the children of the operator node
			}
		}
	}
	// fmt.Println("<<< READ EXPR PREC")
	return lhs, nil
}

func (p *Parser) readExprSeqTo(closingSubtype uint8, allowComma bool, context Context) (string, []*Node, error) {
	seq := []*Node{}
	allowSemicolon := true
	separatorDecided := !allowComma
	// fmt.Println(">>> READ EXPR SEQ TO", closingSubtype, "allowComma", allowComma, "separatorDecided", separatorDecided)
	for p.hasNext() {
		t := p.peek()
		// fmt.Println("Peeked token[1]: ", t.Text, t.Type, t.SubType, closingSubtype)
		if t.Type == CloseBracket {
			if t.SubType == closingSubtype {
				p.next()
				break
			}
			return "", nil, fmt.Errorf("unexpected closing bracket")
		}
		expr, err := p.readExpr(context)
		if err != nil {
			return "", nil, err
		}
		seq = append(seq, expr)
		t = p.peek()
		// fmt.Println("Peeked token[2]: ", t.Text, t.Type, t.SubType, closingSubtype)
		if t.Type == Punctuation {
			// fmt.Println("Punctuation", t.SubType)
			if separatorDecided {
				if t.SubType == PunctuationComma && !allowComma {
					return "", nil, fmt.Errorf("unexpected comma")
				}
				if t.SubType == PunctuationSemicolon && !allowSemicolon {
					return "", nil, fmt.Errorf("unexpected semicolon")
				}
			} else if t.SubType == PunctuationComma {
				allowSemicolon = false
				separatorDecided = true
			} else if t.SubType == PunctuationSemicolon {
				allowComma = false
				separatorDecided = true
			}
			p.next()
			continue
		}
		if t.Type == CloseBracket {
			// fmt.Println("CloseBracket", t.SubType, closingSubtype)
			if t.SubType == closingSubtype {
				p.next()
				break
			}
			fmt.Println("Unexpected closing bracket", t.SubType, closingSubtype)
			return "", nil, fmt.Errorf("unexpected closing bracket")
		} else {
			fmt.Println("Unexpected token", t.Text, t.Type, t.SubType)
			return "", nil, fmt.Errorf("Unexpected token: %s", t.Text)
		}
	}
	// fmt.Println("<<< READ EXPR SEQ TO", "allowComma", allowComma, "separatorDecided", separatorDecided)
	sep_text := chooseSeparator(separatorDecided, allowComma, allowSemicolon)
	return sep_text, seq, nil
}

func chooseSeparator(separatorDecided bool, allowComma bool, allowSemicolon bool) string {
	if separatorDecided {
		if allowComma {
			return "comma"
		}
		if allowSemicolon {
			return "semicolon"
		}
	}
	return "undefined"
}

func (p *Parser) readFormExpr(formStart *Token, context Context) (*Node, error) {
	// fmt.Println(">>> READ FORM EXPR: ", formStart.Text)
	closingTokenText := "end" + formStart.Text
	c := context
	c.InsideForm = true
	c.AcceptNewline = true
	var currentPart []*Node
	currentKeyword := formStart
	content := []*Node{}
	first_expr_in_part := true
	prev_expr_terminated := true
	for {
		if !p.hasNext() {
			return nil, fmt.Errorf("unexpected end of tokens")
		}
		token := p.peek()
		if token.Type == Identifier && token.SubType == IdentifierFormEnd && token.Text == closingTokenText {
			p.next()
			break
		}

		if first_expr_in_part {
			// fmt.Println("::: First expr in part")
			n, err := p.readExpr(c)
			if err != nil {
				return nil, err
			}
			if !p.hasNext() {
				return nil, fmt.Errorf("unexpected end of input in form: %s", formStart.Text)
			}
			t := p.peek()
			if t.IsLabel() {
				// fmt.Println("::: Found label", t.Text)
				p.next()
				currentPart = append(currentPart, n)
				content = append(content, &Node{
					Name: "part",
					Options: map[string]string{
						"keyword": currentKeyword.Text,
					},
					Children: currentPart,
				})
				currentKeyword = p.UnglueOption
				currentPart = []*Node{}
				first_expr_in_part = false
				prev_expr_terminated = true
			} else {
				currentPart = append(currentPart, n)
				first_expr_in_part = false
			}
		} else if token.IsSimpleBreaker() {
			// fmt.Println("::: Simple breaker")
			p.next() // skip the breaker
			p.next() // remove the ':'
			new_currentKeyword := token

			content = append(content, &Node{
				Name: "part",
				Options: map[string]string{
					"keyword": currentKeyword.Text,
				},
				Children: currentPart,
			})

			currentKeyword = new_currentKeyword
			currentPart = []*Node{}
			first_expr_in_part = false
			prev_expr_terminated = true
		} else if token.IsCompoundBreaker(formStart) {
			// fmt.Println("::: Compound breaker")
			t1 := p.next() // skip the breaker
			t2 := p.next() // remove the '-
			t3 := p.next() // remove the form-start

			content = append(content, &Node{
				Name: "part",
				Options: map[string]string{
					"keyword": currentKeyword.Text,
				},
				Children: currentPart,
			})

			currentKeyword = &Token{Text: t1.Text + t2.Text + t3.Text}
			currentPart = []*Node{}
			first_expr_in_part = true
			prev_expr_terminated = true
		} else {
			// fmt.Println("::: Normal expr")
			if !prev_expr_terminated {
				return nil, fmt.Errorf("semi-colon or line-break expected")
			}
			n, err := p.readExpr(c)
			if err != nil {
				return nil, err
			}
			currentPart = append(currentPart, n)
			if !p.hasNext() {
				return nil, fmt.Errorf("unexpected end of input in form: %s", formStart.Text)
			}
			if p.peek().Type == Punctuation && p.peek().SubType == PunctuationSemicolon {
				p.next()
				prev_expr_terminated = true
			} else {
				prev_expr_terminated = p.peek().PrecededByNewline
			}
			first_expr_in_part = false
		}
	}
	if len(currentPart) > 0 {
		content = append(content, &Node{
			Name: "part",
			Options: map[string]string{
				"keyword": currentKeyword.Text,
			},
			Children: currentPart,
		})
	}
	// fmt.Println("<<< READ FORM EXPR")
	return &Node{
		Name:     "form",
		Options:  map[string]string{"syntax": "surround"},
		Children: content,
	}, nil
}

// readDelimitedExpr reads a delimited expression.
func (p *Parser) readDelimitedExpr(open *Token, context Context) (*Node, error) {
	sep, seq, err := p.readExprSeqTo(open.SubType, true, context)
	if err != nil {
		return nil, err
	}
	dname := open.DelimiterName()
	return &Node{
		Name:     "delimited",
		Options:  map[string]string{"kind": dname, "separator": sep},
		Children: seq,
	}, nil
}

func (p *Parser) readPrimaryExpr(context Context) (*Node, error) {
	// fmt.Println(">>> READ PRIMARY EXPR")
	n, e := p.doReadPrimaryExpr(context)
	// fmt.Println("<<< READ PRIMARY EXPR", n)
	return n, e
}

func (p *Parser) doReadPrimaryExpr(context Context) (*Node, error) {
	if !p.hasNext() {
		return nil, fmt.Errorf("unexpected end of tokens")
	}
	token := p.next()

	switch token.Type {
	case Literal:
		switch token.SubType {
		case LiteralString:
			return &Node{
				Name:    "string",
				Options: map[string]string{"quote": token.QuoteWord(), "value": token.Text},
			}, nil
		case LiteralNumber:
			return &Node{
				Name:    "number",
				Options: map[string]string{"value": token.Text},
			}, nil
		}
	case Identifier:
		// fmt.Println("Identifier", token.Text, token.SubType, token.IsMacro())
		if token.IsMacro() {
			p.next()
			// fmt.Println("Label", label.Text, label.SubType)
			n, e := p.readOptExprPrec(token, maxPrecedence, context)
			if e != nil {
				return nil, e
			}

			outer_node := &Node{
				Name: "form",
				Options: map[string]string{
					"syntax": "prefix",
				},
				Children: []*Node{
					{
						Name: "part",
						Options: map[string]string{
							"keyword": token.Text,
						},
					},
				},
			}
			if n != nil {
				outer_node.Children[0].Children = []*Node{n}
			}
			return outer_node, nil
		} else {
			switch token.SubType {
			case IdentifierVariable:
				return &Node{
					Name:    "identifier",
					Options: map[string]string{"name": token.Text},
				}, nil
			case IdentifierFormStart:
				return p.readFormExpr(token, context)
			default:
				return nil, fmt.Errorf("unexpected identifier: %s", token.Text)
			}
		}
	case OpenBracket:
		return p.readDelimitedExpr(token, context)
	case Sign:
		if token.SubType != SignLabel && token.SubType != SignDot {
			prec, valid := token.Precedence()
			if valid && prec > 0 {
				c := context
				c.AcceptNewline = false
				expr, err := p.readExprPrec(prefixPrecedence, c)
				if err != nil {
					return nil, err
				}
				return &Node{
					Name: "operator",
					Options: map[string]string{
						"name":   token.Text,
						"syntax": "prefix",
					},
					Children: []*Node{expr},
				}, nil
			}
		} else {
			return nil, fmt.Errorf("misplace sign token: %s", token.Text)
		}
	default:
		return nil, fmt.Errorf("unexpected token: %s", token.Text)
	}
	return nil, fmt.Errorf("unexpected token: %s", token.Text)
}

func parseTokensToNodes(tokens []*Token, limit bool, breaker string) []*Node {
	parser := &Parser{tokens: tokens, UnglueOption: &Token{Type: Identifier, SubType: IdentifierVariable, Text: breaker}}
	nodes := []*Node{}
	for parser.hasNext() {
		node, err := parser.readExpr(Context{})
		if err != nil {
			// TODO: For the moment we force continuation but we will need
			// to come back nd fix this sooner or later
			// fmt.Println("Error reading primary expression:", err)
		} else {
			nodes = append(nodes, node)
		}
		if limit {
			break
		}
	}
	return nodes
}

func parseToASTArray(input string, limit bool, breaker string) []*Node {
	// Step 1: Tokenize the input
	tokens := tokenizeInput(input)

	// Step 2: Parse the tokens into nodes
	nodes := parseTokensToNodes(tokens, limit, breaker)

	return nodes
}

func parseToAST(input string, foptions *FormatOptions) *Node {
	// Get the array of nodes
	nodes := parseToASTArray(input, foptions.Limit, foptions.UnglueOption)

	var options map[string]string = map[string]string{}
	if foptions.Input != "" {
		options["src"] = foptions.Input
	}

	// Wrap the array in a "unit" node
	var unitNode *Node
	if foptions.Limit && len(nodes) == 1 {
		unitNode = nodes[0]
	} else {
		unitNode = &Node{
			Name:     "unit",
			Options:  options,
			Children: nodes,
		}
	}

	return unitNode
}

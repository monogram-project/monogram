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
	InsideForm   bool
	AllowNewline bool
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
	// TODO: Implement this
	return p.readPrimaryExpr(context)
}

func (p *Parser) readExprSeqTo(closingSubtype uint8, allowComma bool, context Context) (string, []*Node, error) {
	seq := []*Node{}
	allowSemicolon := true
	separatorDecided := !allowComma
	fmt.Println(">>> READ EXPR SEQ TO", closingSubtype, "allowComma", allowComma, "separatorDecided", separatorDecided)
	for p.hasNext() {
		t := p.peek()
		fmt.Println("Peeked token A", t.Text, t.Type, t.SubType, closingSubtype)
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
		fmt.Println("Peeked token B", t.Text, t.Type, t.SubType, closingSubtype)
		if t.Type == Punctuation {
			fmt.Println("Punctuation", t.SubType)
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
			fmt.Println("CloseBracket", t.SubType, closingSubtype)
			if t.SubType == closingSubtype {
				p.next()
				break
			}
			fmt.Println("Unexpected closing bracket", t.SubType, closingSubtype)
			return "", nil, fmt.Errorf("unexpected closing bracket")
		} else {
			fmt.Println("Unexpected token", t.Text, t.Type, t.SubType)
		}
	}
	fmt.Println("<<< READ EXPR SEQ TO", "allowComma", allowComma, "separatorDecided", separatorDecided)
	sep := "unknown"
	if separatorDecided {
		if allowComma {
			sep = "comma"
		} else {
			sep = "semicolon"
		}
	}
	return sep, seq, nil
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
	if !p.hasNext() {
		return nil, fmt.Errorf("unexpected end of tokens")
	}
	token := p.next()

	switch token.Type {
	case Literal:
		switch token.SubType {
		case LiteralString:
			quote := `"` // Assuming double quotes; could be adjusted if necessary.
			return &Node{
				Name:    "string",
				Options: map[string]string{"quote": quote, "value": token.Text},
			}, nil
		case LiteralNumber:
			return &Node{
				Name:    "number",
				Options: map[string]string{"value": token.Text},
			}, nil
		}
	// case Identifier:
	// 	switch token.SubType {
	// 	case IdentifierVariable:
	// 		return &Node{
	// 			Name:    "identifier",
	// 			Options: map[string]string{"name": token.Text},
	// 		}, nil
	// 	case IdentifierFormStart:
	// 		return p.readFormExpr(token)
	// 	default:
	// 		return nil, fmt.Errorf("unexpected identifier: %s", token.Text)
	// 	}
	case OpenBracket:
		return p.readDelimitedExpr(token, context)
	// case Sign:
	// 	if token.SubType == SignForce {
	// 		return nil, fmt.Errorf("Misplaced macro indicator (%s)", token.Text)
	// 	}
	// 	if token.SubType == SignOperator {
	// 		prec, valid := token.Precedence()
	// 		if valid && prec > 0 {
	// 			expr, err := p.readExprPrec(prec, false)
	// 			if err != nil {
	// 				return nil, err
	// 			}
	// 			return &Node{
	// 				Name: "operator",
	// 				Options: map[string]string{
	// 					"name":   token.Text,
	// 					"syntax": "prefix",
	// 				},
	// 				Children: []*Node{expr},
	// 			}, nil
	// 		}
	// 	}
	default:
		return nil, fmt.Errorf("unexpected token: %s", token.Text)
	}
	return nil, fmt.Errorf("unexpected token: %s", token.Text)
}

func parseTokensToNodes(tokens []*Token) []*Node {
	parser := &Parser{tokens: tokens}
	nodes := []*Node{}
	for parser.hasNext() {
		node, err := parser.readPrimaryExpr(Context{})
		if err != nil {
			// TODO: For the moment we force continuation but we will need
			// to come back nd fix this sooner or later
			fmt.Println("Error reading primary expression:", err)
		} else {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func parseToASTArray(input string) []*Node {
	// Step 1: Tokenize the input
	tokens := tokenizeInput(input)

	// Step 2: Parse the tokens into nodes
	nodes := parseTokensToNodes(tokens)

	return nodes
}

func parseToAST(input string, src string) *Node {
	// Get the array of nodes
	nodes := parseToASTArray(input)

	// Wrap the array in a "unit" node
	unitNode := &Node{
		Name:     "unit",
		Options:  map[string]string{"src": src},
		Children: nodes,
	}

	return unitNode
}

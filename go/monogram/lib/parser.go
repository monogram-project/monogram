package lib

import (
	"fmt"
	"strings"
)

type Node struct {
	Name     string            // The name of the node
	Options  map[string]string // Attributes (name-value pairs)
	Children []*Node           // Child nodes
}

const NameForm = "form"
const NamePart = "part"
const NameUnit = "unit"
const NameApply = "apply"
const NameArguments = "arguments"
const NameDelimited = "delimited"
const NameGet = "get"
const NameIdentifier = "identifier"
const NameInvoke = "invoke"
const NameNumber = "number"
const NameOperator = "operator"
const NameString = "string"
const NameJoin = "join"
const NameJoinLines = "joinlines"
const NameInterpolate = "interpolation"

const OptionValue = "value"
const OptionName = "name"
const OptionKind = "kind"
const OptionSeparator = "separator"
const OptionKeyword = "keyword"
const OptionSpan = "span"
const OptionSyntax = "syntax"
const OptionQuote = "quote"
const OptionSrc = "src"

const ValueInfix = "infix"
const ValuePrefix = "prefix"
const ValueSurround = "surround"

// Parser holds the list of tokens and our current reading position.
type Parser struct {
	tokens       []*Token
	pos          int
	UnglueOption *Token
	IncludeSpans bool
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

func (p *Parser) startSpan() (int, int) {
	t := p.peek()
	if t == nil {
		return 0, 0
	}
	return t.Span.StartLine, t.Span.StartColumn
}

func (p *Parser) endSpan() (int, int) {
	if p.pos > 0 {
		t := p.tokens[p.pos-1]
		return t.Span.EndLine, t.Span.EndColumn
	}
	return 0, 0
}

// peek returns the current token without advancing.
func (p *Parser) peek() *Token {
	if p.hasNext() {
		return p.tokens[p.pos]
	}
	return nil
}

func (p *Parser) readExpr(context Context) (*Node, error) {
	n, e := p.readExprPrec(maxPrecedence, context)
	return n, e
}

func (p *Parser) readArguments(subType uint8, context Context) (string, *Node, error) {
	span1, span2 := p.startSpan()
	c := context
	c.AcceptNewline = false
	sep, seq, err := p.readExprSeqTo(subType, true, c)
	if err != nil {
		return "", nil, err
	}
	node := &Node{
		Name:     NameArguments,
		Children: seq,
	}
	if p.IncludeSpans {
		span3, span4 := p.endSpan()
		node.Options = map[string]string{
			OptionSpan: fmt.Sprintf("%d %d %d %d", span1, span2, span3, span4),
		}
	}
	return sep, node, nil
}

func (p *Parser) readOptExprPrec(formStart *Token, outer_prec int, context Context) (*Node, error) {
	if !p.hasNext() {
		return nil, nil
	}
	token := p.peek()
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
	span1, span2 := p.startSpan()
	lhs, err := p.readPrimaryExpr(context)
	if err != nil {
		return nil, err
	}
	for p.hasNext() {
		token1 := p.peek()
		if context.AcceptNewline && token1.PrecededByNewline {
			break
		}
		if context.InsideForm && token1.Type == Sign && token1.SubType == SignLabel {
			break
		}
		if token1.Type == Literal && token1.SubType == LiteralNumber && token1.Text[0] == '-' {
			// Begin the classic hack to handle negative numbers.
			// TODO: Eliminate the duplication of code here.
			fake_minus_token := Token{
				Type:                 Sign,
				SubType:              SignMinus,
				Text:                 "-",
				Span:                 Span{StartLine: token1.Span.StartLine, StartColumn: token1.Span.StartColumn, EndLine: token1.Span.StartLine, EndColumn: token1.Span.StartColumn + 1},
				PrecededByNewline:    token1.PrecededByNewline,
				FollowedByWhitespace: false,
				NextToken:            token1,
			}
			prec, ok := fake_minus_token.Precedence()
			if !ok || prec > outer_prec {
				break
			}
			token1.Text = token1.Text[1:]
			token1.Span.StartColumn++
			token1.PrecededByNewline = false
			c := context
			c.AcceptNewline = false
			rhs, err := p.readExprPrec(prec, c)
			if err != nil {
				return nil, err
			}
			curr_lhs := lhs
			lhs = &Node{
				Name: NameOperator,
				Options: map[string]string{
					OptionName:   fake_minus_token.Text,
					OptionSyntax: "infix",
				},
				Children: []*Node{lhs, rhs}, // lhs and rhs are the children of the operator node
			}
			if p.IncludeSpans {
				curr_lhs_span := strings.Split(curr_lhs.Options[OptionSpan], " ")
				if len(curr_lhs_span) >= 2 {
					span3, span4 := p.endSpan()
					lhs.Options[OptionSpan] = fmt.Sprintf("%s %s %d %d", curr_lhs_span[0], curr_lhs_span[1], span3, span4)
				}
			}
			continue
		}
		prec, ok := token1.Precedence()
		if !ok || prec > outer_prec {
			break
		}
		token2 := p.next()
		c := context
		c.AcceptNewline = false
		if token2.Type == OpenBracket {
			sep_text, args, err := p.readArguments(token2.SubType, c)
			if err != nil {
				return nil, err
			}
			dname := token2.DelimiterName()
			lhs = &Node{
				Name: NameApply,
				Options: map[string]string{
					OptionKind:      dname,
					OptionSeparator: sep_text,
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
					Name: NameInvoke,
					Options: map[string]string{
						OptionKind:      token3.DelimiterName(),
						OptionSeparator: sep_text,
						OptionName:      property.Text,
					},
					Children: []*Node{lhs, rhs},
				}
			} else {
				lhs = &Node{
					Name: NameGet,
					Options: map[string]string{
						OptionName: property.Text,
					},
					Children: []*Node{lhs},
				}
			}
		} else {
			rhs, err := p.readExprPrec(prec, c)
			if err != nil {
				return nil, err
			}
			curr_lhs := lhs
			lhs = &Node{
				Name: NameOperator,
				Options: map[string]string{
					OptionName:   token1.Text,
					OptionSyntax: "infix",
				},
				Children: []*Node{lhs, rhs}, // lhs and rhs are the children of the operator node
			}
			if p.IncludeSpans {
				curr_lhs_span := strings.Split(curr_lhs.Options[OptionSpan], " ")
				if len(curr_lhs_span) >= 2 {
					span3, span4 := p.endSpan()
					lhs.Options[OptionSpan] = fmt.Sprintf("%s %s %d %d", curr_lhs_span[0], curr_lhs_span[1], span3, span4)
				}
			}
		}
	}
	if p.IncludeSpans {
		span3, span4 := p.endSpan()
		lhs.Options[OptionSpan] = fmt.Sprintf("%d %d %d %d", span1, span2, span3, span4)
	}
	return lhs, nil
}

func (p *Parser) readExprSeqTo(closingSubtype uint8, allowComma bool, context Context) (string, []*Node, error) {
	seq := []*Node{}
	allowSemicolon := true
	separatorDecided := !allowComma
	for p.hasNext() {
		t := p.peek()
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
		if t.Type == Punctuation {
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
			if t.SubType == closingSubtype {
				p.next()
				break
			}
			return "", nil, fmt.Errorf("unexpected closing bracket")
		} else {
			return "", nil, fmt.Errorf("unexpected token: %s", t.Text)
		}
	}
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
	closingTokenText := "end" + formStart.Text
	c := context
	c.InsideForm = true
	c.AcceptNewline = true
	var currentPart []*Node
	content := []*Node{}
	first_expr_in_part := true
	prev_expr_terminated := true
	currentKeyword := formStart
	span1, span2 := p.startSpan()
	span3, span4 := 0, 0
	for {
		if !p.hasNext() {
			return nil, fmt.Errorf("unexpected end of tokens")
		}
		token := p.peek()
		if token.Type == Identifier && token.SubType == IdentifierFormEnd && token.Text == closingTokenText {
			span3, span4 = p.endSpan()
			p.next()
			break
		}

		if first_expr_in_part {
			n, err := p.readExpr(c)
			if err != nil {
				return nil, err
			}
			if !p.hasNext() {
				return nil, fmt.Errorf("unexpected end of input in form: %s", formStart.Text)
			}
			t := p.peek()
			if t.IsLabel() {
				span3, span4 := p.endSpan()
				p.next()
				currentPart = append(currentPart, n)
				content = append(content, &Node{
					Name: NamePart,
					Options: map[string]string{
						OptionKeyword: currentKeyword.Text,
					},
					Children: currentPart,
				})
				currentKeyword = p.UnglueOption
				if p.IncludeSpans {
					content[len(content)-1].Options[OptionSpan] = fmt.Sprintf("%d %d %d %d", span1, span2, span3, span4)
					span1, span2 = p.startSpan()
				}
				currentPart = []*Node{}
				first_expr_in_part = false
				prev_expr_terminated = true
			} else {
				currentPart = append(currentPart, n)
				first_expr_in_part = false
			}
		} else if token.IsSimpleBreaker() {
			span3, span4 := p.endSpan()
			p.next() // skip the breaker
			p.next() // remove the ':'
			new_currentKeyword := token

			content = append(content, &Node{
				Name: NamePart,
				Options: map[string]string{
					OptionKeyword: currentKeyword.Text,
				},
				Children: currentPart,
			})

			currentKeyword = new_currentKeyword
			if p.IncludeSpans {
				content[len(content)-1].Options[OptionSpan] = fmt.Sprintf("%d %d %d %d", span1, span2, span3, span4)
				span1, span2 = p.startSpan()
			}
			currentPart = []*Node{}
			first_expr_in_part = false
			prev_expr_terminated = true
		} else if token.IsCompoundBreaker(formStart) {
			span3, span4 := p.endSpan()
			t1 := p.next() // skip the breaker
			t2 := p.next() // remove the '-
			t3 := p.next() // remove the form-start

			content = append(content, &Node{
				Name: NamePart,
				Options: map[string]string{
					OptionKeyword: currentKeyword.Text,
				},
				Children: currentPart,
			})

			currentKeyword = &Token{Text: t1.Text + t2.Text + t3.Text}
			if p.IncludeSpans {
				content[len(content)-1].Options[OptionSpan] = fmt.Sprintf("%d %d %d %d", span1, span2, span3, span4)
				span1, span2 = p.startSpan()
			}
			currentPart = []*Node{}
			first_expr_in_part = true
			prev_expr_terminated = true
		} else {
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
			Name: NamePart,
			Options: map[string]string{
				OptionKeyword: currentKeyword.Text,
			},
			Children: currentPart,
		})
		if p.IncludeSpans {
			content[len(content)-1].Options[OptionSpan] = fmt.Sprintf("%d %d %d %d", span1, span2, span3, span4)
		}
	}
	return &Node{
		Name:     NameForm,
		Options:  map[string]string{OptionSyntax: "surround"},
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
		Name:     NameDelimited,
		Options:  map[string]string{OptionKind: dname, OptionSeparator: sep},
		Children: seq,
	}, nil
}

func (p *Parser) readPrimaryExpr(context Context) (*Node, error) {
	span1, span2 := p.startSpan()
	n, e := p.doReadPrimaryExpr(context)
	if e != nil {
		return nil, e
	}
	if p.IncludeSpans {
		span3, span4 := p.endSpan()
		n.Options[OptionSpan] = fmt.Sprintf("%d %d %d %d", span1, span2, span3, span4)
	}
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
				Name:    NameString,
				Options: map[string]string{OptionQuote: token.QuoteWord(), OptionValue: token.Text},
			}, nil
		case LiteralNumber:
			return &Node{
				Name:    NameNumber,
				Options: map[string]string{OptionValue: token.Text},
			}, nil

		case LiteralInterpolatedString: // Handling interpolated strings
			return p.convertInterpolatedStringSubToken(token)

		case LiteralMultilineString:
			return p.convertMultilineStringSubToken(token)
		}
	case Identifier:
		if token.IsMacro() {
			p.next()
			n, e := p.readOptExprPrec(token, maxPrecedence, context)
			if e != nil {
				return nil, e
			}

			outer_node := &Node{
				Name: NameForm,
				Options: map[string]string{
					OptionSyntax: "prefix",
				},
				Children: []*Node{
					{
						Name: NamePart,
						Options: map[string]string{
							OptionKeyword: token.Text,
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
					Name:    NameIdentifier,
					Options: map[string]string{OptionName: token.Text},
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
					Name: NameOperator,
					Options: map[string]string{
						OptionName:   token.Text,
						OptionSyntax: "prefix",
					},
					Children: []*Node{expr},
				}, nil
			}
		} else {
			return nil, fmt.Errorf("misplace sign token: %s", token.Text)
		}
	default:
		return nil, fmt.Errorf("unexpected token (1): %s, %d, %d", token.Text, token.Type, token.SubType)
	}
	return nil, fmt.Errorf("unexpected token (2): %s, %d, %d", token.Text, token.Type, token.SubType)
}

func (p *Parser) convertMultilineStringSubToken(token *Token) (*Node, error) {
	multilineNode := &Node{
		Name: NameJoinLines,
		Options: map[string]string{
			OptionQuote: token.QuoteWord(),
		},
		Children: []*Node{},
	}

	// Process sub-tokens
	for _, subToken := range token.SubTokens {
		err := p.insertConvertedSubToken(subToken, multilineNode)
		if err != nil {
			return nil, err
		}
	}
	return multilineNode, nil
}

func (p *Parser) insertConvertedSubToken(subToken *Token, interpolationNode *Node) error {
	node, err := p.convertSubToken(subToken)
	if err != nil {
		return err
	}

	interpolationNode.Children = append(interpolationNode.Children, node)
	return nil
}

func (p *Parser) convertSubToken(subToken *Token) (*Node, error) {
	if subToken.SubType == LiteralExpressionString {
		// Recursively parse the expression string
		node, err := p.convertLiteralExpressionStringSubToken(subToken)
		return node, err
	} else if subToken.SubType == LiteralString {
		// Handle plain string parts
		n := &Node{
			Name:    NameString,
			Options: map[string]string{OptionQuote: subToken.QuoteWord(), OptionValue: subToken.Text},
		}
		if p.IncludeSpans {
			n.Options[OptionSpan] = subToken.SpanString()
		}
		return n, nil
	} else if subToken.SubType == LiteralInterpolatedString {
		node, err := p.convertInterpolatedStringSubToken(subToken)
		return node, err
	} else {
		return nil, fmt.Errorf("unexpected sub-token subtype: %v", subToken.SubType)
	}
}

func (p *Parser) convertInterpolatedStringSubToken(token *Token) (*Node, error) {
	interpolationNode := &Node{
		Name: NameJoin,
		Options: map[string]string{
			OptionQuote: token.QuoteWord(),
		},
		Children: []*Node{},
	}
	if p.IncludeSpans {
		interpolationNode.Options[OptionSpan] = token.SpanString()
	}

	// Process sub-tokens
	for _, subToken := range token.SubTokens {
		err := p.insertConvertedSubToken(subToken, interpolationNode)
		if err != nil {
			return nil, err
		}
	}
	return interpolationNode, nil
}

func (p *Parser) convertLiteralExpressionStringSubToken(subToken *Token) (*Node, error) {
	columnOffset := subToken.Span.StartColumn - 1
	expressionNode, err := ParseToAST(subToken.Text, "", true, p.UnglueOption.Text, p.IncludeSpans, columnOffset)
	expressionNode.Name = NameInterpolate // The outer brackets can be repurposed!
	if p.IncludeSpans {
		span := subToken.Span
		span.StartColumn-- // Now ensure we include the backslash.
		expressionNode.Options[OptionSpan] = span.SpanString()
	}
	delete(expressionNode.Options, OptionSeparator)
	if err != nil {
		return nil, fmt.Errorf("error parsing expression string: %v", err)
	}
	return expressionNode, nil
}

func parseTokensToNodes(tokens []*Token, limit bool, breaker string, include_spans bool) ([]*Node, error) {
	parser := &Parser{
		tokens:       tokens,
		UnglueOption: &Token{Type: Identifier, SubType: IdentifierVariable, Text: breaker},
		IncludeSpans: include_spans,
	}
	nodes := []*Node{}
	for parser.hasNext() {
		node, err := parser.readExpr(Context{})
		if err != nil {
			return nil, err
		} else {
			nodes = append(nodes, node)
		}
		if limit {
			break
		}
	}
	return nodes, nil
}

func parseToASTArray(input string, limit bool, breaker string, include_spans bool, colOffset int) ([]*Node, Span, error) {
	// Step 1: Tokenize the input
	tokens, span, terr := tokenizeInput(input, colOffset)
	if terr != nil {
		return nil, Span{}, fmt.Errorf(terr.Message + " (line" + fmt.Sprint(terr.Line) + ", column" + fmt.Sprint(terr.Column) + ")")
	}

	// Step 2: Parse the tokens into nodes
	nodes, err := parseTokensToNodes(tokens, limit, breaker, include_spans)
	if err != nil {
		return nil, Span{}, err
	}

	return nodes, span, nil
}

func ParseToAST(input string, src string, limit bool, unglue string, include_spans bool, colOffset int) (*Node, error) {
	// Get the array of nodes
	nodes, span, err := parseToASTArray(input, limit, unglue, include_spans, colOffset)
	if err != nil {
		return nil, err
	}

	var options map[string]string = map[string]string{}
	if src != "" {
		options[OptionSrc] = src
	}

	// Wrap the array in a "unit" node
	var unitNode *Node
	if limit && len(nodes) == 1 {
		unitNode = nodes[0]
		if include_spans {
			unitNode.Options[OptionSpan] = nodes[0].Options[OptionSpan]
		}
	} else {
		unitNode = &Node{
			Name:     NameUnit,
			Options:  options,
			Children: nodes,
		}
		if include_spans {
			unitNode.Options[OptionSpan] = span.SpanString()
		}
	}

	return unitNode, nil
}

func ParseToElement(input string, src string, limit bool, unglue string, include_spans bool) (Element, error) {
	node, err := ParseToAST(input, src, limit, unglue, include_spans, 0)
	if err != nil {
		return nil, err
	}
	e, err := node.ToElement()
	if err != nil {
		return nil, err
	}
	return e, nil
}

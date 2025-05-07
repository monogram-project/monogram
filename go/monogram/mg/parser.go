package mg

import (
	"fmt"
	"strings"
)

type HowIdentsAreUsed uint8

const (
	NotUsedYet HowIdentsAreUsed = iota
	UsedAsIdentifier
	UsedAsLabel
)

// Parser holds the list of tokens and our current reading position.
type Parser struct {
	tokens       []*Token
	pos          int
	UnglueOption *Token
	IncludeSpans bool
	Decimal      bool
	Idents       map[string]HowIdentsAreUsed
}

func NewParser(tokens []*Token, defaultLabel string, includeSpans bool, decimal bool) *Parser {
	return &Parser{
		tokens:       tokens,
		pos:          0,
		UnglueOption: &Token{Type: Identifier, SubType: IdentifierVariable, Text: defaultLabel},
		IncludeSpans: includeSpans,
		Decimal:      decimal,
		Idents:       make(map[string]HowIdentsAreUsed),
	}
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

func (p *Parser) startLineCol() LineCol {
	t := p.peek()
	if t == nil {
		return LineCol{0, 0}
	}
	return LineCol{t.Span.StartLine, t.Span.StartColumn}
}

func (p *Parser) endLineCol() LineCol {
	if p.pos > 0 {
		t := p.tokens[p.pos-1]
		return LineCol{t.Span.EndLine, t.Span.EndColumn}
	}
	return LineCol{0, 0}
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
	lineCol := p.startLineCol()
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
		node.Options = map[string]string{
			OptionSpan: lineCol.SpanString(p.endLineCol()),
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
	if token.Type == Identifier && token.SubType == IdentifierVariable && token.IsLabelToken(formStart) {
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
	startLineCol := p.startLineCol()
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
			endLineCol := LineCol{token1.Span.StartLine, token1.Span.StartColumn + 1}
			fake_minus_token := Token{
				Type:                 Sign,
				SubType:              SignMinus,
				Text:                 "-",
				Span:                 startLineCol.Span(endLineCol),
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
					OptionSyntax: ValueInfix,
				},
				Children: []*Node{lhs, rhs}, // lhs and rhs are the children of the operator node
			}
			if p.IncludeSpans {
				curr_lhs_span := strings.Split(curr_lhs.Options[OptionSpan], " ")
				if len(curr_lhs_span) >= 2 {
					lc := p.endLineCol()
					lhs.Options[OptionSpan] = fmt.Sprintf("%s %s %d %d", curr_lhs_span[0], curr_lhs_span[1], lc.LineNo, lc.ColNo)
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
		curr_lhs := lhs
		if token2.Type == OpenBracket && token2.SubType != BracketBrace {
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
			if p.hasNext() && p.peek().Type == OpenBracket && p.peek().SubType != BracketBrace {
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
				if p.IncludeSpans {

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
			lhs = &Node{
				Name: NameOperator,
				Options: map[string]string{
					OptionName:   token1.Text,
					OptionSyntax: "infix",
				},
				Children: []*Node{lhs, rhs}, // lhs and rhs are the children of the operator node
			}
		}
		if p.IncludeSpans {
			curr_lhs_span := strings.Split(curr_lhs.Options[OptionSpan], " ")
			if len(curr_lhs_span) >= 2 {
				lc := p.endLineCol()
				lhs.Options[OptionSpan] = fmt.Sprintf("%s %s %d %d", curr_lhs_span[0], curr_lhs_span[1], lc.LineNo, lc.ColNo)
			}
		}
	}
	if p.IncludeSpans {
		lhs.Options[OptionSpan] = startLineCol.SpanString(p.endLineCol())
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
			return ValueComma
		}
		if allowSemicolon {
			return ValueSemicolon
		}
	}
	return ValueUndefined
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
	startLineCol := p.startLineCol()
	var endLineCol LineCol
	for {
		if !p.hasNext() {
			return nil, fmt.Errorf("unexpected end of tokens")
		}
		token := p.peek()
		if token.Type == Identifier && token.SubType == IdentifierFormEnd && token.Text == closingTokenText {
			endLineCol = p.endLineCol()
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
				endLC := p.endLineCol()
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
					content[len(content)-1].Options[OptionSpan] = startLineCol.SpanString(endLC)
					startLineCol = p.startLineCol()
				}
				currentPart = []*Node{}
				first_expr_in_part = false
				prev_expr_terminated = true
			} else {
				currentPart = append(currentPart, n)
				first_expr_in_part = false
				prev_expr_terminated = p.tryReadSemi()
			}
		} else if token.IsSimpleLabelToken() {
			lc34 := p.endLineCol()
			t := p.next() // skip the label
			e := p.SetAsSimpleLabel(t)
			if e != nil {
				return nil, e
			}
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
				content[len(content)-1].Options[OptionSpan] = startLineCol.SpanString(lc34)
				startLineCol = p.startLineCol()
			}
			currentPart = []*Node{}
			first_expr_in_part = false
			prev_expr_terminated = true
		} else if token.IsCompoundLabelToken(formStart) {
			p.SetAsCompoundLabel(token)
			lc34 := p.endLineCol()
			t1 := p.next() // skip the label
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
				content[len(content)-1].Options[OptionSpan] = startLineCol.SpanString(lc34)
				startLineCol = p.startLineCol()
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
			prev_expr_terminated = p.tryReadSemi()
			first_expr_in_part = false
		}
	}
	// if len(currentPart) > 0 {
	content = append(content, &Node{
		Name: NamePart,
		Options: map[string]string{
			OptionKeyword: currentKeyword.Text,
		},
		Children: currentPart,
	})
	if p.IncludeSpans {
		content[len(content)-1].Options[OptionSpan] = startLineCol.SpanString(endLineCol)
	}
	// } else if len(content) == 0 {
	// 	content = append(content, &Node{
	// 		Name:     NamePart,
	// 		Options:  map[string]string{OptionKeyword: formStart.Text},
	// 		Children: []*Node{},
	// 	})
	// }
	return &Node{
		Name:     NameForm,
		Options:  map[string]string{OptionSyntax: ValueSurround},
		Children: content,
	}, nil
}

func (p *Parser) tryReadSemi() bool {
	token := p.peek()
	if token == nil {
		return false
	}
	if token.Type == Punctuation && token.SubType == PunctuationSemicolon {
		p.next()
		return true
	} else {
		return token.PrecededByNewline
	}
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
	lc12 := p.startLineCol()
	n, e := p.doReadPrimaryExpr(context)
	if e != nil {
		return nil, e
	}
	if p.IncludeSpans {
		n.Options[OptionSpan] = lc12.SpanString(p.endLineCol())
	}
	return n, e
}

func (p *Parser) numberOptions(text string) (map[string]string, error) {
	if !p.Decimal {
		return map[string]string{OptionValue: text}, nil
	}
	options := map[string]string{OptionValue: text}

	// Convert the text to a high-precision floating-point number
	// and store it in the options map. We rely on the fact that the value
	// is extremely constrained at this point.
	decimal_text, err := ConvertToDecimal(text)
	if err != nil {
		return nil, err
	}
	options[OptionsDecimalValue] = decimal_text
	return options, nil
}

func (p *Parser) SetAsSimpleLabel(token *Token) error {
	token.SubType = IdentifierSimpleLabel
	if p.Idents[token.Text] == UsedAsIdentifier {
		return fmt.Errorf("identifier used as label: %s", token.Text)
	}
	p.Idents[token.Text] = UsedAsLabel
	return nil
}

func (p *Parser) SetAsCompoundLabel(token *Token) error {
	token.SubType = IdentifierCompoundLabel
	return nil
}

func (p *Parser) SetAsIdentifier(token *Token) error {
	text := token.Text
	if p.Idents[text] == UsedAsLabel {
		return fmt.Errorf("labels used as identifier: %s", text)
	}
	p.Idents[text] = UsedAsIdentifier
	return nil
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
			opts, err := p.numberOptions(token.Text)
			if err != nil {
				return nil, err
			}
			return &Node{
				Name:    NameNumber,
				Options: opts,
			}, nil

		case LiteralInterpolatedString: // Handling interpolated strings
			return p.convertInterpolatedStringSubToken(token)

		case LiteralMultilineString:
			return p.convertMultilineStringSubToken(token)
		}
	case Identifier:
		if token.IsMacro() {
			p.next()
			cxt := context
			cxt.AcceptNewline = true
			startAgain := true

			formBuilder := NewFormBuilder(token.Text, p.startLineCol(), p.IncludeSpans)

			for p.hasNext() {
				if p.tryReadSemi() || p.peek().PrecededByNewline {
					break
				}

				if p.peek().IsSimpleLabelToken() {
					t := p.next()
					e := p.SetAsSimpleLabel(t)
					if e != nil {
						return nil, e
					}
					p.next() // Skip :
					formBuilder.BeginNextPart(t.Text, p.endLineCol(), p.startLineCol())
					startAgain = true
				} else if p.peek().IsCompoundLabelToken(token) {
					p.SetAsCompoundLabel(token)
					t1 := p.next() // skip the label
					t2 := p.next() // remove the '-
					t3 := p.next() // remove the form-start
					text := t1.Text + t2.Text + t3.Text
					formBuilder.BeginNextPart(text, p.endLineCol(), p.startLineCol())
					startAgain = true
				} else if startAgain {
					n, e := p.readOptExprPrec(token, maxPrecedence, cxt)
					if e != nil {
						return nil, e
					}
					startAgain = false
					if n == nil {
						break
					}
					formBuilder.AddChild(n)
				} else {
					n, e := p.readOptExprPrec(token, maxPrecedence, cxt)
					if e != nil {
						return nil, e
					}
					if n != nil {
						formBuilder.BeginNextPart(p.UnglueOption.Text, p.endLineCol(), p.startLineCol())
						formBuilder.AddChild(n)
					} else {
						break
					}
				}
			}

			return formBuilder.Build(p.endLineCol()), nil
		} else {
			switch token.SubType {
			case IdentifierVariable:
				if !token.EscapeSeen {
					if e := p.SetAsIdentifier(token); e != nil {
						return nil, e
					}
				}
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
						OptionSyntax: ValuePrefix,
					},
					Children: []*Node{expr},
				}, nil
			}
		} else {
			return nil, fmt.Errorf("misplace sign token: %s", token.Text)
		}
	default:
		return nil, fmt.Errorf("unexpected token (case #1): %s, %d, %d", token.Text, token.Type, token.SubType)
	}
	return nil, fmt.Errorf("unexpected token (case #2): %s, %d, %d", token.Text, token.Type, token.SubType)
}

func (p *Parser) convertMultilineStringSubToken(token *Token) (*Node, error) {
	multilineNode := &Node{
		Name: NameJoinLines,
		Options: map[string]string{
			OptionQuote:     token.QuoteWord(),
			OptionSpecifier: token.Specifier,
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
	p_opts := &ParserOptions{
		colOffset:    columnOffset,
		DefaultLabel: p.UnglueOption.Text,
		IncludeSpans: p.IncludeSpans,
		Decimal:      p.Decimal,
	}
	expressionNode, err := p_opts.ParseToAST(subToken.Text, "", true)
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

func parseTokensToNodes(tokens []*Token, limit bool, defaultLabel string, includeSpans bool, decimal bool) ([]*Node, error) {
	parser := NewParser(tokens, defaultLabel, includeSpans, decimal)
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

func parseToASTArray(input string, limit bool, defaultLabel string, include_spans bool, decodeNumbers bool, colOffset int) ([]*Node, Span, error) {
	// Step 1: Tokenize the input
	tokens, span, terr := tokenizeInput(input, colOffset)
	if terr != nil {
		return nil, Span{}, fmt.Errorf(terr.Message + " (line " + fmt.Sprint(terr.Line) + ", column " + fmt.Sprint(terr.Column) + ")")
	}

	// Step 2: Parse the tokens into nodes
	nodes, err := parseTokensToNodes(tokens, limit, defaultLabel, include_spans, decodeNumbers)
	if err != nil {
		return nil, Span{}, err
	}

	return nodes, span, nil
}

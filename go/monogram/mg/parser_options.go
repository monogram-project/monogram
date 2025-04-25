package mg

type ParserOptions struct {
	colOffset    int
	DefaultLabel string
	IncludeSpans bool
	Decimal      bool
}

func NewParserOptions() *ParserOptions {
	return &ParserOptions{
		colOffset:    0,
		DefaultLabel: "_",
		IncludeSpans: false,
		Decimal:      false,
	}
}

func (p_opts *ParserOptions) ParseToAST(input string, src string, limit bool) (*Node, error) {
	// Get the array of nodes
	nodes, span, err := parseToASTArray(input, limit, p_opts.DefaultLabel, p_opts.IncludeSpans, p_opts.Decimal, p_opts.colOffset)
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
		if p_opts.IncludeSpans {
			unitNode.Options[OptionSpan] = nodes[0].Options[OptionSpan]
		}
	} else {
		unitNode = &Node{
			Name:     NameUnit,
			Options:  options,
			Children: nodes,
		}
		if p_opts.IncludeSpans {
			unitNode.Options[OptionSpan] = span.SpanString()
		}
	}

	return unitNode, nil
}

func (p_opts *ParserOptions) ParseToElement(input string, src string, limit bool) (Element, error) {
	node, err := p_opts.ParseToAST(input, src, limit)
	if err != nil {
		return nil, err
	}
	e, err := node.ToElement()
	if err != nil {
		return nil, err
	}
	return e, nil
}

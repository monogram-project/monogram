package main

type Node struct {
	Name     string            // The name of the node
	Options  map[string]string // Attributes (name-value pairs)
	Children []*Node           // Child nodes
}

func parseTokensToNodes(tokens []*Token) []*Node {
	// Dummy implementation for now: creates a node array based on dummy tokens
	return []*Node{
		{
			Name:    "node1",
			Options: map[string]string{"key1": "value1"},
			Children: []*Node{
				{
					Name:     "child1",
					Options:  map[string]string{"attribute": "data"},
					Children: nil,
				},
			},
		},
		{
			Name:     "node2",
			Options:  map[string]string{"key2": "value2"},
			Children: nil,
		},
	}
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

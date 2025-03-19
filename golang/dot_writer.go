package main

import (
	"fmt"
	"io"
	"strings"
)

func translateDOT(input io.Reader, output io.Writer, options *FormatOptions) {
	translate(input, output, printASTDOT, options)
}

func printASTDOT(root *Node, indentDelta string, output io.Writer) {
	// Initialize the DOT graph
	fmt.Fprintln(output, `digraph G {`)
	fmt.Fprintln(output, `  bgcolor="transparent";`)
	fmt.Fprintln(output, `  node [shape="box", style="filled", fontname="Ubuntu Mono"];`)

	// Recursively print the nodes and edges
	printNodeDOT(root, "", output)

	// Close the graph
	fmt.Fprintln(output, `}`)
}

func printNodeDOT(node *Node, parentID string, output io.Writer) {
	// Generate a unique identifier for the current node
	nodeID := fmt.Sprintf("node_%p", node)

	// Create the node label
	label := node.Name
	if len(node.Options) == 1 {
		for _, value := range node.Options {
			label = fmt.Sprintf("%s: %s", node.Name, escapeDOTValue(value))
		}
	} else if value, exists := node.Options["value"]; exists {
		label = fmt.Sprintf("%s: %s", node.Name, escapeDOTValue(value))
	} else if name, exists := node.Options["name"]; exists {
		label = fmt.Sprintf("%s: %s", node.Name, escapeDOTValue(name))
	}

	// Determine the fill color based on the tag
	fillColor := tagColors[node.Name]
	if fillColor == "" {
		fillColor = "lightgray"
	}

	// Add the node definition to the DOT graph
	fmt.Fprintf(output, "  \"%s\" [label=\"%s\", shape=\"box\", fillcolor=\"%s\"];\n", nodeID, label, fillColor)

	// If there's a parent node, add an edge
	if parentID != "" {
		fmt.Fprintf(output, "  \"%s\" -> \"%s\";\n", parentID, nodeID)
	}

	// Recurse for child nodes
	for _, child := range node.Children {
		printNodeDOT(child, nodeID, output)
	}
}

func escapeDOTValue(value string) string {
	// Escape special characters for DOT format
	return strings.ReplaceAll(value, `"`, `\"`)
}

var tagColors = map[string]string{
	"form":       "lightpink",
	"part":       "#FFD8E1",
	"apply":      "lightgreen",
	"identifier": "Honeydew",
	"arguments":  "PaleTurquoise",
	"operator":   "#C0FFC0",
	"number":     "lightgoldenrodyellow",
}

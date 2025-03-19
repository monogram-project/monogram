package main

import (
	"fmt"
	"io"
	"strings"
)

func translateMermaid(input io.Reader, output io.Writer, options *FormatOptions) {
	translate(input, output, printASTMermaid, options)
}

func printASTMermaid(root *Node, indentDelta string, output io.Writer) {
	// Initialize the Mermaid graph
	fmt.Fprintln(output, "graph LR")

	// Recursively print the nodes and their relationships
	printNodeMermaid(root, "", output)
	addStyles(output)
}

func printNodeMermaid(node *Node, parentID string, output io.Writer) {
	// Generate a unique ID for the current node
	nodeID := fmt.Sprintf("node_%p", node)

	// Create the node label
	label := node.Name
	if len(node.Options) == 1 {
		for _, value := range node.Options {
			label = fmt.Sprintf("%s: %s", node.Name, escapeMermaidValue(value))
		}
	} else if value, exists := node.Options["value"]; exists {
		label = fmt.Sprintf("%s: %s", node.Name, escapeMermaidValue(value))
	} else if name, exists := node.Options["name"]; exists {
		label = fmt.Sprintf("%s: %s", node.Name, escapeMermaidValue(name))
	}

	// Format the label for Mermaid
	if len(label) > 20 {
		label = strings.Replace(label, ": ", ":<br/>", 1)
	}

	// Define the node's style based on its type
	nodeStyle := fmt.Sprintf("custom_%s", node.Name)

	// Add the node definition to the Mermaid graph
	fmt.Fprintf(output, "  %s[\"%s\"]:::%s;\n", nodeID, label, nodeStyle)

	// If there's a parent node, add an edge
	if parentID != "" {
		fmt.Fprintf(output, "  %s --> %s;\n", parentID, nodeID)
	}

	// Recurse for child nodes
	for _, child := range node.Children {
		printNodeMermaid(child, nodeID, output)
	}
}

func addStyles(output io.Writer) {
	// Define styles for specific tags (similar to your Python tag_colors dictionary)
	tagColors := map[string]string{
		"form":       "lightpink",
		"part":       "#FFD8E1",
		"apply":      "lightgreen",
		"identifier": "Honeydew",
		"arguments":  "PaleTurquoise",
		"operator":   "#C0FFC0",
		"number":     "lightgoldenrodyellow",
	}

	// Write the style definitions for each tag
	for tag, color := range tagColors {
		fmt.Fprintf(output, "classDef custom_%s fill:%s,stroke:#333,stroke-width:2px;\n", tag, color)
	}
}

func escapeMermaidValue(value string) string {
	// Escape special characters for Mermaid (minimal requirements)
	return strings.ReplaceAll(value, "\"", "\\\"")
}

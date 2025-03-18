package main

import (
	"fmt"
	"io"
	"strings"
)

func translateYAML(input io.Reader, output io.Writer, src string, indent int) {
	// fmt.Fprintln(output, "YAML Translation Output:")
	translate(input, output, printASTYAML, src, indent)
}

func printASTYAML(root *Node, indentDelta string, output io.Writer) {
	// Print the root node (which is the "unit" node)
	printNodeYAML(root, 0, indentDelta, output)
}

func n1print(indent int, indentDelta string, output io.Writer) {
	for i := 0; i < indent; i++ {
		fmt.Fprint(output, indentDelta) // Indent the node
	}
}

func printNodeYAML(node *Node, currentIndent int, indentDelta string, output io.Writer) {
	// Print the node name as a key at the current indentation
	if currentIndent > 0 {
		n1print(currentIndent-1, indentDelta, output)
		fmt.Fprintf(output, "- role: %s\n", node.Name)
	}

	// Print attributes (options)
	for key, value := range node.Options {
		escapedValue := escapeYAMLString(value) // Escape YAML special characters
		n1print(currentIndent, indentDelta, output)
		fmt.Fprintf(output, "%s: %s\n", key, escapedValue)
	}

	// Print children (if any)
	if len(node.Children) > 0 {
		n1print(currentIndent, indentDelta, output)
		fmt.Fprintf(output, "children:\n")
		childIndent := currentIndent + 1
		for _, child := range node.Children {
			printNodeYAML(child, childIndent, indentDelta, output) // Indent child nodes appropriately
		}
	}
}

func escapeYAMLString(value string) string {
	// Check if the string needs escaping (contains special characters)
	var needsQuotes bool
	for _, r := range value {
		if r == '\n' || r == '\t' || r == ':' || r == '-' || r == '\'' || r == '"' || r == '&' ||
			r == '*' || r == '!' || r == '|' || r == '>' || r == '%' || r == '@' ||
			r == '{' || r == '}' || r == '[' || r == ']' || r == '`' || r == '#' {
			needsQuotes = true
			break
		}
	}

	// If no special characters, return as-is
	if !needsQuotes {
		return value
	}

	// Escape the string using double quotes
	var sb strings.Builder
	sb.WriteRune('"') // Start double-quoted string

	for _, r := range value {
		switch r {
		case '"':
			sb.WriteString("\\\"") // Escape double quote
		case '\\':
			sb.WriteString("\\\\") // Escape backslash
		case '\n':
			sb.WriteString("\\n") // Escape newline
		case '\t':
			sb.WriteString("\\t") // Escape tab
		case '\r':
			sb.WriteString("\\r") // Escape carriage return
		default:
			sb.WriteRune(r) // Add normal characters
		}
	}

	sb.WriteRune('"') // End double-quoted string
	return sb.String()
}

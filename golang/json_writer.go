package main

import (
	"fmt"
	"io"
	"strings"
)

func translateJSON(input io.Reader, output io.Writer, src string, indent int) {
	fmt.Fprintln(output, "JSON Translation Output:")
	translate(input, output, printASTJSON, src, indent)
}

// escapeJSONString ensures all strings in JSON are properly escaped.
func escapeJSONString(value string) string {
	var sb strings.Builder

	for _, r := range value {
		switch r {
		case '"':
			sb.WriteString("\\\"") // Escape double quotes
		case '\\':
			sb.WriteString("\\\\") // Escape backslashes
		case '\b':
			sb.WriteString("\\b") // Escape backspace
		case '\f':
			sb.WriteString("\\f") // Escape form feed
		case '\n':
			sb.WriteString("\\n") // Escape newline
		case '\r':
			sb.WriteString("\\r") // Escape carriage return
		case '\t':
			sb.WriteString("\\t") // Escape tab
		default:
			if r >= 0x20 && r <= 0x7E {
				sb.WriteRune(r) // Printable ASCII characters
			} else {
				sb.WriteString(fmt.Sprintf("\\u%04X", r)) // Non-printable and Unicode characters
			}
		}
	}

	return sb.String()
}

// printASTJSON starts printing the AST structure as JSON.
func printASTJSON(nodes []*Node, src string, indentDelta string, output io.Writer) {
	// Start the currentIndent with the indentDelta
	currentIndent := indentDelta

	// Open the JSON array
	// Open the enclosing object
	fmt.Fprintf(output, "{\n")
	fmt.Fprintf(output, "%s\"role\": \"unit\",\n", currentIndent)
	fmt.Fprintf(output, "%s\"src\": \"%s\",\n", currentIndent, escapeJSONString(src))
	fmt.Fprintf(output, "%s\"children\": [\n", currentIndent)

	// Print the child nodes
	childIndent := currentIndent + indentDelta
	for i, node := range nodes {
		printNodeJSON(node, childIndent, indentDelta, output) // Adjust child indentation
		if i < len(nodes)-1 {
			fmt.Fprintln(output, ",") // Add a comma for all but the last node
		} else {
			fmt.Fprintln(output)
		}
	}

	// Close the children array and the enclosing object
	fmt.Fprintf(output, "%s]\n", currentIndent)
	fmt.Fprintf(output, "}\n")
}

// printNodeJSON recursively prints a single node and its children in JSON format.
func printNodeJSON(node *Node, currentIndent string, indentDelta string, output io.Writer) {
	// Precompute the next level of indentation
	nextIndent := currentIndent + indentDelta

	// Open the object
	fmt.Fprintf(output, "%s{\n", currentIndent)

	// Include the `role` field
	fmt.Fprintf(output, "%s\"role\": \"%s\",\n", nextIndent, node.Name)

	// Flatten the options map directly into string-valued fields
	optionCount := len(node.Options)
	current := 0
	for key, value := range node.Options {
		current++
		escapedValue := escapeJSONString(value)              // Escape the option value
		if current < optionCount || len(node.Children) > 0 { // Add a comma if there are more fields or children
			fmt.Fprintf(output, "%s\"%s\": \"%s\",\n", nextIndent, key, escapedValue)
		} else {
			fmt.Fprintf(output, "%s\"%s\": \"%s\"\n", nextIndent, key, escapedValue)
		}
	}

	// Print the children field
	if len(node.Children) > 0 {
		fmt.Fprintf(output, "%s\"children\": [\n", nextIndent)

		childIndent := nextIndent + indentDelta
		for i, child := range node.Children {
			printNodeJSON(child, childIndent, indentDelta, output)
			if i < len(node.Children)-1 {
				fmt.Fprintln(output, ",") // Add a comma for all but the last child
			} else {
				fmt.Fprintln(output)
			}
		}

		fmt.Fprintf(output, "%s]\n", nextIndent) // Close the JSON array for children
	}

	// Close the object
	fmt.Fprintf(output, "%s}", currentIndent)
}

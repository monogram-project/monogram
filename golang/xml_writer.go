package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

func translateXML(input io.Reader, output io.Writer, options *FormatOptions) {
	// fmt.Fprintln(output, "XML Translation Output:")
	translate(input, output, printASTXML, options)
}

func printASTXML(root *Node, indentDelta string, output io.Writer) {
	// Print the root node (which is the "unit" node)
	printNodeXML(root, "", indentDelta, output)
}

func printNodeXML(node *Node, currentIndent string, indentDelta string, output io.Writer) {
	// Open the XML tag
	fmt.Fprintf(output, "%s<%s", currentIndent, node.Name)

	// Sort the attributes (Options) by their keys
	sortedKeys := make([]string, 0, len(node.Options))
	for key := range node.Options {
		sortedKeys = append(sortedKeys, key)
	}
	// Sort the keys alphabetically
	sort.Strings(sortedKeys)

	// Print attributes in alphabetical order
	for _, key := range sortedKeys {
		value := node.Options[key]
		fmt.Fprintf(output, ` %s="%s"`, key, escapeXMLValue(value))
	}

	// Handle self-closing tag if no children are present
	if len(node.Children) == 0 {
		fmt.Fprintln(output, " />")
		return
	}

	// Otherwise, close the opening tag and iterate over children
	fmt.Fprintln(output, ">")
	newIndent := currentIndent + indentDelta
	for _, child := range node.Children {
		printNodeXML(child, newIndent, indentDelta, output)
	}

	// Close the XML tag
	fmt.Fprintf(output, "%s</%s>\n", currentIndent, node.Name)
}

func escapeXMLValue(value string) string {
	var sb strings.Builder

	for _, r := range value {
		switch r {
		// XML special characters
		case '&':
			sb.WriteString("&amp;") // Escape ampersand
		case '<':
			sb.WriteString("&lt;") // Escape less than
		case '>':
			sb.WriteString("&gt;") // Escape greater than
		case '"':
			sb.WriteString("&quot;") // Escape double quote
		case '\'':
			sb.WriteString("&apos;") // Escape single quote

		// Handle printable ASCII characters (no escaping required)
		default:
			if r >= 0x20 && r <= 0x7E {
				sb.WriteRune(r)
			} else {
				// Non-printable and Unicode characters
				sb.WriteString(fmt.Sprintf("&#x%X;", r)) // Use numeric character references
			}
		}
	}

	return sb.String()
}

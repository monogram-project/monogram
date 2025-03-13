package main

import (
	"fmt"
	"io"
	"strings"
)

func translateXML(input io.Reader, output io.Writer, indent int) {
	fmt.Fprintln(output, "XML Translation Output:")
	translate(input, output, printASTXML, indent)
}

func printASTXML(nodes []*Node, indentDelta string, output io.Writer) {
	for _, node := range nodes {
		printNodeXML(node, "", indentDelta, output)
	}
}

func printNodeXML(node *Node, currentIndent string, indentDelta string, output io.Writer) {
	// Open the XML tag
	fmt.Fprintf(output, "%s<%s", currentIndent, node.Name)

	// Print attributes (with escaping)
	for key, value := range node.Options {
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
	replacer := strings.NewReplacer(
		"&", "&amp;",
		"<", "&lt;",
		">", "&gt;",
		"\"", "&quot;",
		"'", "&apos;",
	)
	return replacer.Replace(value)
}

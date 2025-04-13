package lib

import (
	"fmt"
	"io"
	"sort"

	asciitree "github.com/thediveo/go-asciitree"
)

type AsciiNode struct {
	Label    string      `asciitree:"label"`
	Props    []string    `asciitree:"properties"`
	Children []AsciiNode `asciitree:"children"`
}

// convertToTree converts your Node structure to an asciitree.Tree using the custom label.
func convertToTree(n *Node) AsciiNode {
	label := n.Name

	// Extract and sort keys lexically
	sortedKeys := make([]string, 0, len(n.Options))
	for key := range n.Options {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys) // Sort keys alphabetically

	// Generate properties list in sorted order
	var props []string
	for _, key := range sortedKeys {
		props = append(props, fmt.Sprintf("%s: %s", key, n.Options[key]))
	}

	// Convert children recursively
	var children []AsciiNode
	for _, child := range n.Children {
		children = append(children, convertToTree(child))
	}
	return AsciiNode{
		Label:    label,
		Props:    props,
		Children: children,
	}
}

func PrintASTAsciiTree(root *Node, indentDelta string, output io.Writer) {
	fmt.Fprintln(output, asciitree.RenderFancy(convertToTree(root)))
}

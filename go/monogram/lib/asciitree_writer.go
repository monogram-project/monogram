package lib

import (
	"fmt"
	"io"

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
	var props []string
	for key, value := range n.Options {
		props = append(props, fmt.Sprintf("%s: %s", key, value))
	}
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

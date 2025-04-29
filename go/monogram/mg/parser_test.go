package mg

import (
	"fmt"
	"testing"
)

func getTokens(input string) ([]*Token, error) {
	tokens, _, err := tokenizeInput(input, 0)
	if err != nil {
		return nil, fmt.Errorf("tokenizeInput error: %s", err.Message)
	}
	return tokens, nil
}

func getParser(input string) (*Parser, error) {
	tokens, err := getTokens(input)
	if err != nil {
		return nil, fmt.Errorf("getTokens error: %w", err)
	}
	return NewParser(tokens, "_", false, false), nil
}

func TestParsePrefix0(t *testing.T) {
	// Arrange
	// Step 1: Tokenize the input
	parser, err := getParser("return!")
	if err != nil {
		t.Fatalf("getParser error: %v", err)
		return
	}
	cxt := Context{false, false}

	// Act
	node, err := parser.doReadPrimaryExpr(cxt)
	if err != nil {
		t.Fatalf("doReadPrimaryExpr error: %v", err)
		return
	}

	// Assert
	if node.Name != "form" {
		t.Errorf("Expected node name 'form', got '%s'", node.Name)
	}
	if len(node.Children) != 1 {
		t.Errorf("Expected 1 child, got %d", len(node.Children))
	}
	part := node.Children[0]
	if part.Name != "part" {
		t.Errorf("Expected part name 'part', got '%s'", part.Name)
	}
	if part.Options["keyword"] != "return" {
		t.Errorf("Expected keyword 'return', got '%s'", part.Options["keyword"])
	}
	if len(part.Children) != 0 {
		t.Errorf("Expected 0 children, got %d", len(part.Children))
	}
}

func TestParsePrefix1(t *testing.T) {
	// Arrange
	// Step 1: Tokenize the input
	parser, err := getParser("return! x")
	if err != nil {
		t.Fatalf("getParser error: %v", err)
		return
	}
	cxt := Context{false, false}

	// Act
	node, err := parser.doReadPrimaryExpr(cxt)
	if err != nil {
		t.Fatalf("doReadPrimaryExpr error: %v", err)
		return
	}

	// Assert
	if node.Name != "form" {
		t.Errorf("Expected node name 'form', got '%s'", node.Name)
	}
	if len(node.Children) != 1 {
		t.Errorf("Expected 1 child, got %d", len(node.Children))
	}
	part := node.Children[0]
	if part.Name != "part" {
		t.Errorf("Expected part name 'part', got '%s'", part.Name)
	}
	if part.Options["keyword"] != "return" {
		t.Errorf("Expected keyword 'return', got '%s'", part.Options["keyword"])
	}
	if len(part.Children) != 1 {
		t.Errorf("Expected 1 children, got %d", len(part.Children))
	}
	child := part.Children[0]
	if child.Name != "identifier" {
		t.Errorf("Expected child name 'identifier', got '%s'", child.Name)
	}
}

func TestParsePrefix2NoComma(t *testing.T) {
	// Arrange
	// Step 1: Tokenize the input
	parser, err := getParser("while! (x) { x = f(y)}")
	if err != nil {
		t.Fatalf("getParser error: %v", err)
		return
	}
	cxt := Context{false, false}

	// Act
	node, err := parser.doReadPrimaryExpr(cxt)
	if err != nil {
		t.Fatalf("doReadPrimaryExpr error: %v", err)
		return
	}

	// Assert
	if node.Name != "form" {
		t.Errorf("Expected node name 'form', got '%s'", node.Name)
	}
	if len(node.Children) != 2 {
		t.Errorf("Expected 1 child, got %d", len(node.Children))
	}
	part0 := node.Children[0]
	if part0.Name != "part" {
		t.Errorf("Expected part name 'part', got '%s'", part0.Name)
	}
	part1 := node.Children[1]
	if part1.Name != "part" {
		t.Errorf("Expected part name 'part', got '%s'", part1.Name)
	}
	if part0.Options["keyword"] != "while" {
		t.Errorf("Expected keyword 'while	', got '%s'", part0.Options["keyword"])
	}
	if len(part0.Children) != 1 {
		t.Errorf("Expected 1 child, got %d", len(part0.Children))
	}
	child0 := part0.Children[0]
	if child0.Name != "delimited" || child0.Options["kind"] != "parentheses" {
		t.Errorf("Expected parenthesized expression")
	}
	if part1.Options["keyword"] != "_" {
		t.Errorf("Expected keyword '_', got '%s'", part1.Options["keyword"])
	}
	if len(part1.Children) != 1 {
		t.Errorf("Expected 1 child, got %d", len(part1.Children))
	}
	child1 := part1.Children[0]
	if child1.Name != "delimited" || child1.Options["kind"] != "braces" {
		t.Errorf("Expected brace expression")
	}
}

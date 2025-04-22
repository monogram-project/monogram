package main

// Demonstrates the use of the strongly-typed abstract syntax tree.

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/chzyer/readline"
	"github.com/sfkleach/monogram/go/monogram/mg"
)

type State struct {
	variables map[string]float64
	stack     []float64
}

func (s *State) EvaluateChildren(element mg.Element) error {
	children := element.Children()
	for children.Next() {
		child := children.Value()
		err := s.Evaluate(child)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *State) Evaluate(element mg.Element) error {
	switch node := element.(type) {
	case *mg.NumberNode:
		s.stack = append(s.stack, node.Value)
	case *mg.IdentifierNode:
		s.stack = append(s.stack, s.variables[node.Id])
	case *mg.InfixOperatorNode:
		switch node.Op {
		case "=":
			id, ok := node.LHS.(*mg.IdentifierNode)
			if !ok {
				return fmt.Errorf("left-hand side of assignment must be an identifier")
			}
			vname := id.Id
			s.Evaluate(node.RHS)
			s.variables[vname] = s.stack[len(s.stack)-1]
		case "+":
			e := s.EvaluateChildren(element)
			if e != nil {
				return e
			}
			s.stack[len(s.stack)-2] += s.stack[len(s.stack)-1]
			s.stack = s.stack[:len(s.stack)-1]

		case "-":
			e := s.EvaluateChildren(element)
			if e != nil {
				return e
			}
			s.stack[len(s.stack)-2] -= s.stack[len(s.stack)-1]
			s.stack = s.stack[:len(s.stack)-1]

		case "*":
			e := s.EvaluateChildren(element)
			if e != nil {
				return e
			}
			s.stack[len(s.stack)-2] *= s.stack[len(s.stack)-1]
			s.stack = s.stack[:len(s.stack)-1]

		case "/":
			e := s.EvaluateChildren(element)
			if e != nil {
				return e
			}
			s.stack[len(s.stack)-2] /= s.stack[len(s.stack)-1]
			s.stack = s.stack[:len(s.stack)-1]

		default:
			return fmt.Errorf("unknown infix operator %s", element.Name())
		}
	case *mg.PrefixOperatorNode:
		if node.Op == "-" {
			err := s.Evaluate(node.Arg)
			if err != nil {
				s.stack[len(s.stack)-1] = -s.stack[len(s.stack)-1]
			}
			return err
		} else {
			return fmt.Errorf("unknown prefix operator %s", node.Op)
		}
	default:
		return fmt.Errorf("unknown expression (%s)", element.Name())
	}
	return nil
}

func (s *State) Show() {
	if len(s.stack) == 1 {
		fmt.Printf("Result: %v\n", s.stack[0])
	} else {
		fmt.Printf("Stack: %v\n", s.stack)
		for name, value := range s.variables {
			fmt.Printf("%s: %v\n", name, value)
		}
	}
}

func main() {

	fmt.Println("A simple calculator demonstration using monogram syntax")
	fmt.Println("Use expressions such as x = 3 + 4 or y = x * 2")
	fmt.Println()

	state := &State{
		variables: make(map[string]float64),
		stack:     []float64{},
	}

	scanner := bufio.NewScanner(os.Stdin) // Create a scanner to read lines from stdin

	// Create a readline instance
	rl, err := readline.New(">>> ") // Prompt for input
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing readline: %v\n", err)
		os.Exit(1)
	}
	defer rl.Close()

	for {
		// Read user input
		line, err := rl.Readline()
		if err != nil { // Handle EOF or other errors
			if err == readline.ErrInterrupt { // Handle Ctrl+C
				break
			}
			if err == io.EOF { // Handle Ctrl+D
				break
			}
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			continue
		}

		element, err := mg.ParseToElement(line, "", false, "_", true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
			continue
		}
		children := element.Children()
		for children.Next() {
			child := children.Value()
			err2 := state.Evaluate(child)
			if err2 == nil {
				state.Show()
			} else {
				fmt.Println("Error evaluating", err2)
			}
			state.stack = state.stack[:0]
		}
	}

	// Check for scanning errors (other than EOF)
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}

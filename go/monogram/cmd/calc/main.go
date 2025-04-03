package main

// Demonstrates the use of the mono-typed abstract syntax tree.

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/chzyer/readline"
	"github.com/sfkleach/monogram/go/monogram/lib"
)

type State struct {
	variables map[string]float64
	stack     []float64
}

func (s *State) EvaluateChildren(node *lib.Node) error {
	for _, child := range node.Children {
		err := s.Evaluate(child)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *State) Evaluate(node *lib.Node) error {
	switch node.Name {
	case lib.NameNumber:
		str := node.Options[lib.OptionValue]
		num, err := strconv.ParseFloat(str, 64) // 64-bit float
		if err != nil {
			return err
		}
		s.stack = append(s.stack, num)
	case lib.NameIdentifier:
		name := node.Options[lib.OptionName]
		s.stack = append(s.stack, s.variables[name])
	case lib.NameOperator:
		name := node.Options[lib.OptionName]
		syntax := node.Options[lib.OptionSyntax]
		if syntax == "infix" {
			switch name {
			case "=":
				lhs := node.Children[0]
				rhs := node.Children[1]
				if lhs.Name != lib.NameIdentifier {
					return fmt.Errorf("left-hand side of assignment must be an identifier")
				}
				vname := lhs.Options[lib.OptionName]
				s.Evaluate(rhs)
				s.variables[vname] = s.stack[len(s.stack)-1]

			case "+":
				e := s.EvaluateChildren(node)
				if e != nil {
					return e
				}
				s.stack[len(s.stack)-2] += s.stack[len(s.stack)-1]
				s.stack = s.stack[:len(s.stack)-1]

			case "-":
				e := s.EvaluateChildren(node)
				if e != nil {
					return e
				}
				s.stack[len(s.stack)-2] -= s.stack[len(s.stack)-1]
				s.stack = s.stack[:len(s.stack)-1]

			case "*":
				e := s.EvaluateChildren(node)
				if e != nil {
					return e
				}
				s.stack[len(s.stack)-2] *= s.stack[len(s.stack)-1]
				s.stack = s.stack[:len(s.stack)-1]

			case "/":
				e := s.EvaluateChildren(node)
				if e != nil {
					return e
				}
				s.stack[len(s.stack)-2] /= s.stack[len(s.stack)-1]
				s.stack = s.stack[:len(s.stack)-1]

			default:
				return fmt.Errorf("unknown infix operator %s", name)
			}
		} else if name == "-" {
			err := s.Evaluate(node.Children[0])
			if err != nil {
				s.stack[len(s.stack)-1] = -s.stack[len(s.stack)-1]
			}
			return err
		} else {
			return fmt.Errorf("unknown prefix operator %s", name)
		}
	default:
		return fmt.Errorf("unknown expression (%s)", node.Name)
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

		node, err := lib.ParseToAST(line, "", false, "_", true, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
			continue
		}
		for _, child := range node.Children {
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

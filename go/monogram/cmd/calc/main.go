package main

// This is the entry point of the program. It processes command-line arguments
// and performs file format translation based on user-specified flags. The
// program supports built-in formats (e.g., XML and JSON) as well as delegating
// to external subprograms for custom formats.
//
// Flags:
// - --help: Displays help information for the program and available flags.
// - --format (-f): Specifies the output format. Required for both built-in and external formats.
// - --input (-i): Specifies the input file. If omitted, standard input (stdin) is used.
// - --output (-o): Specifies the output file. If omitted, standard output (stdout) is used.
//
// Built-in Formats:
// - xml: The program processes input and outputs in XML format.
// - json: The program processes input and outputs in JSON format.
// Additional built-in formats can be added by updating the global formatHandlers map.
//
// For non-built-in formats, the program delegates processing to a subprogram named "monogram-to-{format}".
//
// Usage Example:
// To translate a file to JSON format:
//
//	monogram --format json --input input.txt --output output.json
//
// To delegate to a custom subprogram:
//
//	monogram --format custom --input input.txt --output output.custom

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
			fmt.Println("Unary minus")
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

		node, err := lib.ParseToAST(line, "", false, "_", true)
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

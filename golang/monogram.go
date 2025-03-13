package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"syscall"
)

// Define a type for the translation function
type translationFunc func(io.Reader, io.Writer)

// Global map for format-to-function associations
// Updated formatHandlers map
var formatHandlers = map[string]func(io.Reader, io.Writer, int){
	"xml":  translateXML,
	"json": translateJSON,
}

// main is the entry point of the program. It processes command-line arguments
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
func main() {
	// Define flags
	helpFlag := flag.Bool("help", false, "Display help information")
	formatFlag := flag.String("format", "", "Output format")
	formatShortFlag := flag.String("f", "", "Output format (short)")
	inputFlag := flag.String("input", "", "Input file (optional, defaults to stdin)")
	inputShortFlag := flag.String("i", "", "Input file (short, defaults to stdin)")
	outputFlag := flag.String("output", "", "Output file (optional, defaults to stdout)")
	outputShortFlag := flag.String("o", "", "Output file (short, defaults to stdout)")
	indentFlag := flag.Int("indent", 2, "Number of spaces for indentation (0 for no formatting)") // New flag
	flag.Parse()

	// Determine the effective format and input/output
	format := *formatFlag
	if format == "" {
		format = *formatShortFlag
	}

	input := *inputFlag
	if input == "" {
		input = *inputShortFlag
	}

	output := *outputFlag
	if output == "" {
		output = *outputShortFlag
	}

	// Check if the format is built-in
	translator, isBuiltInFormat := formatHandlers[format]

	// Handle --help option
	if *helpFlag && isBuiltInFormat {
		fmt.Println("Usage: monogram [OPTIONS] < stdin > stdout")
		flag.PrintDefaults()
	}

	// Open input (default to stdin if input is not provided)
	var inputReader io.Reader
	if input == "" {
		fmt.Println("No input file specified. Using standard input.")
		inputReader = os.Stdin
	} else {
		file, err := os.Open(input)
		if err != nil {
			log.Fatalf("Error: Cannot open input file '%s': %v", input, err)
		}
		defer file.Close()
		inputReader = file
	}

	// Open output (default to stdout if output is not provided)
	var outputWriter io.Writer
	if output == "" {
		fmt.Println("No output file specified. Using standard output.")
		outputWriter = os.Stdout
	} else {
		file, err := os.Create(output)
		if err != nil {
			log.Fatalf("Error: Cannot create output file '%s': %v", output, err)
		}
		defer file.Close()
		outputWriter = file
	}

	// Handle built-in formats
	if isBuiltInFormat {
		translator(inputReader, outputWriter, *indentFlag) // Pass the indent parameter
		return
	}

	// For non-built-in formats, exec into a subprogram
	if format == "" {
		log.Fatalf("Error: Format was not specified.")
	}

	execName := "monogram-to-" + format
	newArgs := make([]string, len(os.Args))
	newArgs[0] = execName
	copy(newArgs[1:], os.Args[1:])

	err := syscall.Exec(execName, newArgs, os.Environ())
	if err != nil {
		log.Fatalf("Failed to execute %s: %v", execName, err)
	}
}

type Node struct {
	Name     string            // The name of the node
	Options  map[string]string // Attributes (name-value pairs)
	Children []*Node           // Child nodes
}

func parseToAST(input string) []*Node {
	// Dummy implementation for now: returns a single root node with no children
	// In a real-world case, this would parse the string into MinXML or JSON AST structure.
	root := &Node{
		Name:    "root",
		Options: map[string]string{"example": "true"},
		Children: []*Node{
			{
				Name:     "child1",
				Options:  map[string]string{"attribute1": "value1"},
				Children: nil,
			},
			{
				Name:     "child2",
				Options:  map[string]string{},
				Children: nil,
			},
		},
	}
	return []*Node{root}
}

func translate(input io.Reader, output io.Writer, printAST func([]*Node, string, io.Writer), indentSpaces int) {
	// Read the entire input as a string
	data, err := io.ReadAll(input)
	if err != nil {
		log.Fatalf("Error: Failed to read input: %v", err)
	}

	// Convert the input string into an AST
	ast := parseToAST(string(data))

	// Determine the indentation string (spaces or none)
	indent := ""
	if indentSpaces > 0 {
		indent = strings.Repeat("#", indentSpaces)
	}

	// Use the provided print function to recursively print the AST
	printAST(ast, indent, output)
}

func translateXML(input io.Reader, output io.Writer, indent int) {
	fmt.Fprintln(output, "XML Translation Output:")
	translate(input, output, printASTXML, indent)
}

func translateJSON(input io.Reader, output io.Writer, indent int) {
	fmt.Fprintln(output, "JSON Translation Output:")
	translate(input, output, printASTJSON, indent)
}

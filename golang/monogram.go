package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"syscall"
)

// Define a type for the translation function
type translationFunc func(io.Reader, io.Writer)

// Global map for format-to-function associations
var formatHandlers = map[string]translationFunc{
	"xml":  translateXML,
	"json": translateJSON,
}

func main() {
	// Define flags
	helpFlag := flag.Bool("help", false, "Display help information")
	formatFlag := flag.String("format", "", "Output format")
	formatShortFlag := flag.String("f", "", "Output format (short)")
	inputFlag := flag.String("input", "", "Input file (optional, defaults to stdin)")
	inputShortFlag := flag.String("i", "", "Input file (short, defaults to stdin)")
	outputFlag := flag.String("output", "", "Output file (optional, defaults to stdout)")
	outputShortFlag := flag.String("o", "", "Output file (short, defaults to stdout)")
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
		translator(inputReader, outputWriter)
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

// translateXML handles XML-specific translation logic
func translateXML(input io.Reader, output io.Writer) {
	fmt.Fprintln(output, "Converting to XML format...")
	// Add XML-specific logic here
}

// translateJSON handles JSON-specific translation logic
func translateJSON(input io.Reader, output io.Writer) {
	fmt.Fprintln(output, "Converting to JSON format...")
	// Add JSON-specific logic here
}

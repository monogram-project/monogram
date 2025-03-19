package main

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

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/spf13/pflag"
)

type FormatOptions struct {
	Format       string
	Input        string
	Output       string
	Indent       int
	Limit        bool
	UnglueOption string
}

// setupFlags initializes a flag set with the common flag definitions.
func setupFlags(fs *pflag.FlagSet, options *FormatOptions, optionsFile *string, showHelp *bool, classifyTokens *bool) {
	fs.StringVarP(&options.Format, "format", "f", options.Format, "Output format xml|json|yaml|mermaid|dot")
	fs.StringVarP(&options.Input, "input", "i", options.Input, "Input file (optional, defaults to stdin)")
	fs.StringVarP(&options.Output, "output", "o", options.Output, "Output file (optional, defaults to stdout)")
	fs.IntVar(&options.Indent, "indent", options.Indent, "Number of spaces for indentation (0 for no formatting)")
	fs.BoolVar(&options.Limit, "one", options.Limit, "Process only one monogram value and do not wrap in a unit node")
	fs.StringVarP(&options.UnglueOption, "default-breaker", "b", options.UnglueOption, "Default breakers")
	if optionsFile != nil {
		fs.StringVar(optionsFile, "options-file", "", "File containing additional options")
	}
	if showHelp != nil {
		pflag.BoolVarP(showHelp, "help", "h", false, "Display help information")
	}
	if classifyTokens != nil {
		fs.BoolVar(classifyTokens, "classify-tokens", false, "Classify tokens")
	}
}

// Define a type for the translation function
type translationFunc func(io.Reader, io.Writer, *FormatOptions)

// Global map for format-to-function associations
var formatHandlers = map[string]translationFunc{
	"xml":     translateXML,
	"json":    translateJSON,
	"yaml":    translateYAML,
	"mermaid": translateMermaid,
	"dot":     translateDOT,
}

func main() {

	// Initialize the options struct
	options := FormatOptions{
		Format: "",
		Input:  "",
		Output: "",
		Indent: 2,
		Limit:  false,
	}

	var optionsFile string
	var showHelp bool
	var classifyTokens bool

	// Set up the main command-line flag set
	setupFlags(pflag.CommandLine, &options, &optionsFile, &showHelp, &classifyTokens)

	// Parse command-line flags first to check for `--options-file`
	pflag.Parse()

	// Process options file if specified
	if optionsFile != "" {
		fileArgs, err := readOptionsFile(optionsFile)
		if err != nil {
			log.Fatalf("Error reading options file: %v", err)
		}

		// Create a temporary FlagSet for file-based options
		fileFlagSet := pflag.NewFlagSet("file-flags", pflag.ContinueOnError)
		setupFlags(fileFlagSet, &options, nil, nil, nil) // Reuse the same setup logic
		if err := fileFlagSet.Parse(fileArgs); err != nil {
			log.Fatalf("Error parsing options from file: %v", err)
		}
	}

	// Re-parse the command-line arguments to ensure they override file-based options
	pflag.Parse()

	// Check for help flag
	if showHelp {
		fmt.Println("Monogram: converts program-like text in monogram notation to various other formats.")
		fmt.Println("\nUsage:")
		fmt.Println("  monogram [OPTIONS] < STDIN > STDOUT")
		fmt.Println("\nOptions:")
		pflag.PrintDefaults()
		os.Exit(0) // Exit after displaying the help message
	}

	// Check if the format is built-in
	translator, isBuiltInFormat := formatHandlers[options.Format]

	// Open input (default to stdin if input is not provided)
	var inputReader io.Reader
	if options.Input == "" {
		inputReader = os.Stdin
	} else {
		file, err := os.Open(options.Input)
		if err != nil {
			log.Fatalf("Error: Cannot open input file '%s': %v", options.Input, err)
		}
		defer file.Close()
		inputReader = file
	}

	// Open output (default to stdout if output is not provided)
	var outputWriter io.Writer
	if options.Output == "" {
		outputWriter = os.Stdout
	} else {
		file, err := os.Create(options.Output)
		if err != nil {
			log.Fatalf("Error: Cannot create output file '%s': %v", options.Output, err)
		}
		defer file.Close()
		outputWriter = file
	}

	// Handle built-in formats
	if isBuiltInFormat {
		translator(inputReader, outputWriter, &options)
		return
	} else if classifyTokens {
		VSCodeClassifyTokens(inputReader, outputWriter)
		return
	}

	// For non-built-in formats, exec into a subprogram
	if options.Format == "" {
		log.Fatalf("Error: Format was not specified.")
	}

	execName := "monogram-to-" + options.Format
	newArgs := make([]string, len(os.Args))
	newArgs[0] = execName
	copy(newArgs[1:], os.Args[1:])

	err := syscall.Exec(execName, newArgs, os.Environ())
	if err != nil {
		log.Fatalf("Failed to execute %s: %v", execName, err)
	}
}

// readOptionsFile reads the options from the specified file and returns them as a slice of strings
func readOptionsFile(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Split the file into individual arguments (by whitespace or newlines)
	content := string(data)
	args := strings.Fields(content) // Splits by any whitespace (including newlines)
	return args, nil
}

func translate(input io.Reader, output io.Writer, printAST func(*Node, string, io.Writer), options *FormatOptions) {
	// Read the entire input as a string
	data, err := io.ReadAll(input)
	if err != nil {
		log.Fatalf("Error: Failed to read input: %v", err)
	}

	// Convert the input string into an AST
	ast, err := parseToAST(string(data), options)
	if err != nil {
		log.Fatalf("Error: Failed to parse input: %v", err)
	}

	// Determine the indentation string (spaces or none)
	indent := ""
	if options.Indent > 0 {
		indent = strings.Repeat(" ", options.Indent)
	}

	// Use the provided print function to recursively print the AST
	printAST(ast, indent, output)
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
)

func main() {
	helpFlag := flag.Bool("help", false, "Display help information")
	formatFlag := flag.String("format", "", "Output format")
	formatShortFlag := flag.String("f", "", "Output format (short)")
	flag.String("input", "", "Input file")
	flag.String("i", "", "Input file (short)")
	outputFlag := flag.String("output", "", "Output file")
	outputShortFlag := flag.String("o", "", "Output file (short)")
	flag.Parse()

	// If help flag is provided, display help and exit
	if *helpFlag && *formatFlag == "" && *formatShortFlag == "" {
		fmt.Println("Usage: monogram [OPTIONS] < stdin > stdout")
		flag.PrintDefaults()
		return
	}

	// Determine which output flag to use (short or long).
	output := *outputFlag
	if output == "" {
		output = *outputShortFlag
	}

	format := *formatFlag
	if format == "" {
		format = *formatShortFlag
	}
	if format == "" {
		ext := filepath.Ext(output)
		if ext != "" {
			format = strings.TrimPrefix(ext, ".")
		}
	}
	if format == "" {
		log.Fatalf("error: unable to determine output format (no --format provided and no extension on output file)")
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

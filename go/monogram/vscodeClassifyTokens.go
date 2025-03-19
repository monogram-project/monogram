package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

type Position struct {
	Line   int `json:"line"`   // Line number (0-indexed)
	Column int `json:"column"` // Column number (0-indexed)
}

type TokenClassification struct {
	Type  string   `json:"type"`  // e.g., "keyword", "identifier", "literal"
	Value string   `json:"value"` // The actual token string
	Start Position `json:"start"` // Start position in the input
	End   Position `json:"end"`   // End position in the input
}

func VSCodeClassifyTokens(input io.Reader, output io.Writer) {
	// Read the entire input as a string
	data, err := io.ReadAll(input)
	if err != nil {
		log.Fatalf("Error: Failed to read input: %v", err)
	}
	// fmt.Println("VSCodeClassifyTokens: input data:", string(data))
	tokens, terr := tokenizeInput(string(data))
	if terr != nil {
		// fmt.Println("Error: Failed to tokenize input", terr)
		jsonOutput, err := json.MarshalIndent(map[string]interface{}{
			"error": terr, // Nest the TokenizerError under an "error" field
		}, "", "  ")
		if err != nil {
			log.Fatalf("Error serializing TokenizerError to JSON: %v", err)
		}
		fmt.Println(string(jsonOutput)) // Output JSON to stdout
		return
	}

	var classifications []TokenClassification = []TokenClassification{}

	for _, token := range tokens {
		classifications = append(classifications, TokenClassification{
			Type:  token.VSCodeTokenType(),
			Value: token.Text,
			Start: Position{Line: token.StartLine, Column: token.StartColumn},
			End:   Position{Line: token.EndLine, Column: token.EndColumn}, // TODO!
		})
	}

	// Serialize the classifications into JSON
	jsonOutput, err := json.MarshalIndent(classifications, "", "  ")
	if err != nil {
		log.Fatalf("Error serializing JSON: %v", err)
	}
	fmt.Println(string(jsonOutput)) // Output JSON to stdout
}

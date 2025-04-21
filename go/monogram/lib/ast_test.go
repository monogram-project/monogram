package lib

import (
	"os"
	"testing"

	"gopkg.in/yaml.v3"
)

// Test structure matching the YAML structure
type TestCase struct {
	Name               string `yaml:"name"`
	Command            string `yaml:"command"`
	Input              string `yaml:"input"`
	ExpectedOutput     string `yaml:"expected_output"`
	Normalize          string `yaml:"normalize,omitempty"`
	ExpectedExitStatus int    `yaml:"expected_exit_status"`
}

type TestFile struct {
	Normalize string     `yaml:"normalize"`
	Tests     []TestCase `yaml:"tests"`
}

// CheckTranslationToAST is the function to be tested
func CheckTranslationToAST(input string) error {
	node, err := ParseToAST(input, "", false, "_", false, 0)
	if err != nil {
		return err
	}
	_, err = node.ToElement()
	if err != nil {
		return err
	}
	return nil
}

func TestAST(t *testing.T) {
	// Load the YAML file - we run from go/monogram/lib
	data, err := os.ReadFile("../../../functests/xml-tests.yaml")
	if err != nil {
		t.Fatalf("Failed to read YAML file: %v", err)
	}

	// Parse the YAML file
	var testFile TestFile
	if err := yaml.Unmarshal(data, &testFile); err != nil {
		t.Fatalf("Failed to parse YAML file: %v", err)
	}

	// Iterate over test cases
	for _, testCase := range testFile.Tests {
		// Only run tests with expected exit status 0
		if testCase.ExpectedExitStatus == 0 {
			t.Run(testCase.Name, func(t *testing.T) {
				err := CheckTranslationToAST(testCase.Input)
				if err != nil {
					t.Errorf("Cannot parse test-case '%s': %v", testCase.Name, err)
				}
			})
		}
	}
}

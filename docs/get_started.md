# Get Started

## Prerequisites

- Go 1.23 or later installed on your machine.
- A working `GOPATH` or a Go workspace with module support.

## Installing the CLI

To install the Monogram CLI, run the following command:

```sh
go install github.com/sfkleach/monogram/go/monogram/cmd/monogram@latest
```

This command will compile the CLI and install the monogram executable in your
$GOBIN directory. Make sure $GOBIN is in your system's PATH so that you can
invoke monogram directly from the terminal.

## Using the Monogram CLI tool

The Monogram CLI tool allows you to process and analyze monogram notation directly from the command line. Below are some common use cases and examples to help you get started.

### Basic Usage

Once installed, you can invoke the Monogram CLI tool using the `monogram` command. For example:

```sh
monogram [flags] [arguments]
```

To see a list of available commands and options, run:

```sh
monogram --help
```

### Example Commands

1. **Processing a Monogram File**

   To process a monogram file and output the result in a specific format (e.g., JSON):

   ```sh
   monogram -format json < input.monogram
   ```

   Replace `input.monogram` with the path to your monogram file.

2. **Specifying Output File**

   To save the output to a file instead of printing it to the terminal:

   ```sh
   monogram --format xml --o output.xml -i input.monogram
   ```

   This will process the `input.monogram` file and save the result as `output.xml`.

3. **Using Indentation**

   To control the indentation level in the output (e.g., for JSON or XML):

   ```sh
   monogram -format json -indent 4 -i input.monogram
   ```

   This will format the output with 4 spaces of indentation.

4. **Including Spans**

   To include span information in the output (useful for debugging) and emit as YAML:

   ```sh
   monogram -include-spans -format yaml -i input.monogram
   ```

### Flags and Options

Here are some commonly used flags:

- `-format <format>`: Specify the output format. Supported formats include
  `json`, `xml`, `yaml`, `mermaid`, and `dot`.
- `-output <file>`: Specify the output file. If omitted, the output is printed to the terminal.
- `-indent <number>`: Set the number of spaces for indentation in the output.
- `-include-spans`: Include span information in the output.


This will display a comprehensive list of available options and their descriptions.



## Using the Library

If you want to integrate Monogram's functionality into your own project, you can import the library package:

```go
import "github.com/sfkleach/monogram/go/monogram/lib"
```

For example, you might call an exported function like this:

```go
package main

import (
    "os"
    "log"
    "github.com/sfkleach/monogram/go/monogram/lib"
)

func main() {
    // PrintASTXML is an exported function in lib/xml_writer.go
    if err := lib.PrintASTXML(myAST, os.Stdout); err != nil {
        log.Fatalf("Error printing AST to XML: %v", err)
    }
}
```

## Building from Source

If you'd like to build everything from source manually then follow these steps.

### Step 1: Clone the Repository

```sh
git clone https://github.com/sfkleach/monogram.git
cd monogram/go/monogram
```

### Step 2: Build the CLI

From the module root (go/monogram), run:
```sh
go build -o monogram ./cmd/monogram
```

### Step 3: Test the Library:

```
go test ./lib
```

### Step 4: Usage

Once installed, run the CLI tool with:
```sh
monogram [flags] [arguments]
```

For help and command-line options, use:
```sh
monogram -h
```
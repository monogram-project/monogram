# Get Started

This guide covers getting started with the reference implemention [Monogram-Go](https://github.com/monogram-project/monogram-go). 

You can install the `monogram` command-line tool or the Go library, if you want
to write a Go application that uses `Monogram` notation, or both! 

**Command-line**:
- [Get started with the Monogram command-line tool](get_started_cli.md)
- [Get started with the Monogram docker container](get_started_container.md)

**Go Library**
- [Get started with the Monogram Go module](get_started_go_mod.md)


## Using the Monogram CLI tool

Once you have installed the `monogram` tool, you can use it to process and
analyze monogram notation directly from the command line. Listed below are some
common use cases and examples to help you get started.

### Basic Usage

Once installed, you can invoke the Monogram CLI tool using the `monogram`
command. For example:

```sh
monogram [flags] [arguments]
```

To see a list of available commands and options, run:

```sh
monogram --help
```

### Example Commands

1. **Processing a Monogram File**

   To process a monogram file and output the result in a specific format (e.g.,
   JSON):

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


This will display a comprehensive list of available options and their
descriptions.

### Exploring with the --test option

You can also start up with `monogram --test`. This runs a local web-server and
opens a web page that allows you to experiment with Monogram. It's a neat way to
learn more about the notation and to find problems.

Follow [this link](test_option.md) to learn more about this option.





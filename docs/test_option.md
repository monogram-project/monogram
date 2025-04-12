# Monogram option: --test

This section looks at the `--test` option available with the
`monogram` command. Using this option will launch an HTTP test server that
serves as an exploratory interface for Monogram. This feature is especially
useful during development and for learning about monogram, as it enables you to
interact with Monogram through a web interface. 


Here's a screenshot showing it in use:

![screenshot](images/test-option-screenshot.png)

n.b. This option is not available in `monogram-mini`.

## Usage

The following options can be used with `--test` to set the initial defaults.

```
  -b, --default-breaker string   Default breakers (default "_")
  -f, --format string            Output format xml|json|yaml|mermaid|dot
  --include-spans                Include start/end of expressions in the output
  --indent int                   Number of spaces for indentation, 0 for 
                                 no formatting, default 2
```

You can simply start the test-server on port 3000 with this command. It will
immediately attempt to open a browser window on that port, ready to go.

```bash
# Start the HTTP test server on port 3000 using the full build (with web support).
monogram --test 3000
```

If you omit the port number it will default to port 8080. And if you omit the --test option entirely, Monogram runs in the standard CLI mode.


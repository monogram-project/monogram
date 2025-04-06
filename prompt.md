We are collaborating on a Go program, codename `monogram`. The `monogram` tool
is a command-line tool that translates from `monogram` notation into XML, JSON
and other formats. The notation is designed to represent program-like texts.
However it is just a notation and not a programming language, although it does
have an opinionated grammar. Consequently it has no built-in variables, no
built-in operators and even the reserved words are dynamically discovered during
the parse.

We have completed a good first version of that tool. The repo that it resides in
is a monorepo with a parallel implementation, written in some other programming
language (Pop-11). As a consequence the go folder is not at the top-level of the
repo but in `~/go/monogram/`. The application itself is in
`~/go/monogram/cmd/monogram/main.go`.

I am interested in making it possible to run the monogram executable from a
docker image. This is the work-in-progress Dockerfile. 

```Dockerfile
# ---------
# Build stage
# ---------
FROM golang:1.23.8-alpine AS builder
WORKDIR /go/monogram

# Copy the Go project directory (adjust the context below if needed)
COPY go/monogram/ .

# Install essential tools (curl is used to install just)
RUN apk update && apk add --no-cache curl

# Install Just from its install script
RUN curl --proto '=https' --tlsv1.2 -sSf https://just.systems/install.sh | sh -s -- --to /usr/local/bin

# Run the Justfile recipe to build the monogram executable
RUN just build

# ----------
# Runtime stage
# ----------
FROM alpine:3.18
WORKDIR /app

# Copy the monogram binary from the builder stage
COPY --from=builder /go/monogram/monogram /app/monogram

# Ensure the binary is executable
RUN chmod +x /app/monogram

# Set the entrypoint to run the executable
ENTRYPOINT ["/app/monogram"]

```


And this is the output of `monogram --help`, indicating the command-line options.

```
Monogram: converts program-like text in monogram notation to various other formats.

Usage:
  monogram [OPTIONS] < STDIN > STDOUT

Options:
      --classify-tokens          Classify tokens
  -b, --default-breaker string   Default breakers (default "_")
  -f, --format string            Output format xml|json|yaml|mermaid|dot
  -h, --help                     Display help information
      --include-spans            Include start/end of expressions in the output
      --indent int               Number of spaces for indentation (0 for no formatting) (default 2)
  -i, --input string             Input file (optional, defaults to stdin)
      --one                      Process only one monogram value and do not wrap in a unit node
      --options-file string      File containing additional options
  -o, --output string            Output file (optional, defaults to stdout)
      --version                  Display the version information
```

Modify the Dockerfile so that it passes the options through and makes it 
possible to use the -i and -o options.
# Get Started developing with the Monogram Go module

## Prerequisites

- Go 1.23 or later installed on your machine.
- A Go project that you would like to use Monogram in.

## Installing the library

To add the library to your project use this command:

```sh
cd {$YOUR_PROJECT_HOME}
go get github.com/monogram-project/monogram-go@latest
```

This will add it to your projects `go.mod` file.


## Using the Library

If you want to integrate Monogram's functionality into your own project, you can import the library package. We suggest using an alias "monogram".

```go
import monogram "github.com/monogram-project/monogram-go/mg"
```

For example, you might call an exported function like this:

```go
package main

import (
    "os"
    "log"
    monogram "github.com/monogram-project/monogram-go/mg"
)

func main() {
    // Code not shown for brevity but assigns to `myAST`
    ...
    // PrintASTXML is an exported function in lib/xml_writer.go
    if err := monogram.PrintASTXML(myAST, os.Stdout); err != nil {
        log.Fatalf("Error printing AST to XML: %v", err)
    }
}
```

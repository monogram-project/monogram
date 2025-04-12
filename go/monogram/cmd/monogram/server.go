//go:build withweb
// +build withweb

package main

import (
	// ... existing imports ...
	"html/template"
	"log"
	"net/http"

	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func withWeb() bool {
	return true
}

var formTemplate = template.Must(template.New("form").Parse(`
<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<title>Monogram Test Interface</title>
	<script src="https://cdn.jsdelivr.net/npm/mermaid@10/dist/mermaid.min.js"></script>
	<script>mermaid.initialize({ startOnLoad: true });</script>
	<style>
        body {
            font-family: sans-serif; /* Set font to sans-serif */
            background-color: #f7f7ff; /* Light pastel background */
            color: #333; /* Dark text for readability */
            margin: 0;
            padding: 20px;
        }
        h1 {
            color: #4a4a8c; /* Pastel blue for the header */
        }
        textarea, input[type="text"], input[type="number"], select {
            border-radius: 8px; /* Add rounded corners */
            padding: 8px; /* Add some padding for better appearance */
            border: 1px solid #ccc; /* Add a light border */
            font-size: 14px; /* Adjust font size for consistency */
        	font-family: 'Fira Code', "Courier New", Courier, monospace; /* Use a font that renders underscores clearly */
            background-color: #ffffff; /* High-contrast white background */
            color: #000; /* Black text for readability */
        }
		textarea#monogramInput {
			width: 100%; /* Make the text box span the full width */
			box-sizing: border-box; /* Include padding and border in the width calculation */
		}
	    input[type="number"]#indent {
            width: 3ch; /* Set width to fit approximately two characters */
        }
        textarea:focus, input:focus, select:focus {
            outline: none;
            border-color: #4a4a8c; /* Highlight border on focus */
            box-shadow: 0 0 5px #4a4a8c;
        }
        .output {
            border: 1px solid #000; /* Add a light border */
            border-radius: 8px; /* Add rounded corners */
            padding: 8px; /* Add some padding for better appearance */
            background-color: #e8f0ff; /* Light pastel blue background */
            font-size: 14px; /* Adjust font size for consistency */
        	font-family: 'Fira Code', "Courier New", Courier, monospace; /* Use a font that renders underscores clearly */

        }
        input[type="submit"] {
            background-color: #ff6f61; /* Bold coral color for the button */
            color: #fff; /* White text for contrast */
            border: none;
            padding: 10px 20px;
            font-size: 16px;
            font-weight: bold;
            border-radius: 8px;
            cursor: pointer;
        }
        input[type="submit"]:hover {
            background-color: #e65b50; /* Slightly darker coral on hover */
        }
		.form-row {
			display: flex;
			flex-wrap: wrap; /* Allow wrapping when the container is compressed */
       		justify-content: flex-start; /* Align items to the left */
			gap: 2%; /* Use a percentage-based gap for flexibility */
			row-gap: 1em; /* Add vertical spacing between rows */
			margin-bottom: 20px; /* Add spacing below the row */
		}
		.form-row label {
			margin-right: 0.5em; /* Add spacing between label and input */
		}
		textarea#monogramInput {
			flex: 1; /* Allow the text area to take up available space */
			box-sizing: border-box; /* Include padding and border in the width calculation */
		}
		select, input[type="number"], input[type="text"] {
			flex: 0; /* Keep other inputs compact */
		}
		.form-row > div {
			display: flex; /* Use flexbox for alignment */
			align-items: center; /* Vertically align items to the center */
			gap: 0.5em; /* Add spacing between the label and the input */
		}
		input[type="checkbox"] {
			margin: 0; /* Remove default margins */
		}
		.raw-output {
			width: 100%; /* Make the text area span the full width */
			box-sizing: border-box; /* Include padding and border in the width calculation */
			margin-top: 1em; /* Add spacing above the raw output */
			background-color: #f9f9f9; /* Light gray background for distinction */
			color: #333; /* Dark text for readability */
		}
		.copy-link {
			display: inline-block; /* Ensure proper spacing and alignment */
			margin-top: 1em; /* Add vertical separation from the div above */
			font-size: 12px; /* Set a smaller font size */
			color: #888; /* Light grey color for the link */
			text-decoration: none; /* Remove underline */
			cursor: pointer; /* Show pointer cursor on hover */
		}
		.copy-link:hover {
			color: #555; /* Slightly darker grey on hover */
			text-decoration: underline; /* Add underline on hover for clarity */
		}
		.toaster {
			display: none; /* Hidden by default */
			position: fixed; /* Fixed position on the screen */
			top: 20px; /* 20px from the top */
			right: 20px; /* 20px from the right */
			background-color: #4caf50; /* Green background for success */
			color: white; /* White text for contrast */
			padding: 10px 20px; /* Padding for better appearance */
			border-radius: 5px; /* Rounded corners */
			box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2); /* Subtle shadow */
			font-size: 14px; /* Font size */
			z-index: 1000; /* Ensure it appears above other elements */
		}
	</style>
</head>
<body>
<h1>Monogram Test Interface</h1>
	<form action="/translate" method="post">

	    <div class="form-row">

			<div>
				<label for="format">Format:</label>
				<select name="format" id="format">
				<option value="xml" {{if eq .Format "xml"}}selected{{end}}>XML</option>
				<option value="json" {{if eq .Format "json"}}selected{{end}}>JSON</option>
				<option value="yaml" {{if eq .Format "yaml"}}selected{{end}}>YAML</option>
				<option value="mermaid" {{if eq .Format "mermaid"}}selected{{end}}>Mermaid</option>
				<option value="dot" {{if eq .Format "dot"}}selected{{end}}>DOT</option>
				</select>
			</div>
			
			<div>
				<label for="indent">Indentation:</label>
				<input type="number" id="indent" name="indent" value="{{.Indent}}" min="0" size="2">
			</div>

			<div>
				<label for="defaultBreaker">Default Breaker:</label>
				<input type="text" id="defaultBreaker" name="defaultBreaker" value="{{.Breaker}}">
			</div>
			
			<div>
				<label for="includeSpans">Include Spans:</label>
				<input type="checkbox" id="includeSpans" name="includeSpans" {{if .IncludeSpans}}checked{{end}}>
			</div>
		</div>

		<textarea name="monogramInput" id="monogramInput" rows="10" placeholder="Enter monogram text here">{{.MonogramInput}}</textarea><br><br>

		<input type="submit" value="Translate">
	</form>
	<br>
	{{if .Output}}
		<div class="output">
			{{if eq .Format "mermaid"}}
				<div class="mermaid">{{.Output}}</div> <!-- Render diagram only -->
			{{else}}
				<pre id="outputText">{{.Output}}</pre> <!-- Keep normal output format -->
			{{end}}
		</div>
		{{if eq .Format "mermaid"}}
			<textarea id="outputText" readonly class="raw-output" rows="10">{{.Output}}</textarea> <!-- Raw output -->
		{{end}}
    	<a href="#" class="copy-link" onclick="copyToClipboard()">Copy to clipboard</a>
		<div id="toaster" class="toaster">Output copied to clipboard!</div>

		<script>
			function copyToClipboard() {
				const outputElement = document.getElementById("outputText");
				if (outputElement) {
					const range = document.createRange();
					range.selectNode(outputElement);
					window.getSelection().removeAllRanges();
					window.getSelection().addRange(range);
					document.execCommand("copy");
					window.getSelection().removeAllRanges();
					// Show toaster message
					const toaster = document.getElementById("toaster");
					toaster.style.display = "block";
					setTimeout(() => {
						toaster.style.display = "none";
					}, 3000); // Hide after 3 seconds				
				} else {
					alert("No output to copy!");
				}
			}
		</script>
	{{end}}
</body>
</html>
`))

// startTestServer starts an HTTP listener on the specified port and opens the browser.
func startTestServer(port string, options *FormatOptions) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		indexHandler(w, r, options)
	})
	http.HandleFunc("/translate", translateHandler)

	if port == "" {
		port = "3000"
	}
	addr := "localhost:" + port
	go openBrowser("http://" + addr)
	log.Printf("Starting test server on %s...", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Failed to start test server: %v", err)
	}
}

func indexHandler(w http.ResponseWriter, _ *http.Request, options *FormatOptions) {
	format := "xml" // Default format
	if options.Format != "" {
		format = options.Format
	}
	formTemplate.Execute(w, struct {
		Output        string
		MonogramInput string
		Format        string
		IncludeSpans  bool
		Indent        int
		Breaker       string
	}{
		Output:        "",
		MonogramInput: "",
		Format:        format, // Default format
		IncludeSpans:  options.IncludeSpans,
		Indent:        options.Indent,
		Breaker:       options.UnglueOption,
	})
}

// translateHandler processes the form and renders the translation output.
func translateHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form: "+err.Error(), http.StatusBadRequest)
		return
	}
	monogramInput := r.FormValue("monogramInput")
	format := r.FormValue("format")
	indentVal := r.FormValue("indent")
	defaultBreaker := r.FormValue("defaultBreaker")
	includeSpans := r.FormValue("includeSpans") == "on"

	// Convert indent value to integer:
	indent := 2
	if indentParsed, err := strconv.Atoi(indentVal); err == nil {
		indent = indentParsed
	}

	// Set up FormatOptions based on the form values:
	options := FormatOptions{
		Format:       format,
		Input:        "", // Not used in test mode — we’re using form data.
		Output:       "", // Output will be captured in a buffer.
		Indent:       indent,
		Limit:        false,
		UnglueOption: defaultBreaker,
		IncludeSpans: includeSpans,
	}

	// Look up the translator function:
	translator, ok := formatHandlers[format]
	if !ok {
		http.Error(w, "Unknown format: "+format, http.StatusBadRequest)
		return
	}

	// Create reader from input text and a bytes.Buffer for capturing output:
	inputReader := strings.NewReader(monogramInput)
	var outputBuffer bytes.Buffer

	// Perform the translation.
	translator(inputReader, &outputBuffer, &options)

	// Render the same form with the translation output shown:
	formTemplate.Execute(w, struct {
		Output        string
		MonogramInput string
		Format        string
		IncludeSpans  bool
		Indent        int
		Breaker       string
	}{
		Output:        outputBuffer.String(),
		MonogramInput: monogramInput,
		Format:        format, // Pass format to the template
		IncludeSpans:  includeSpans,
		Indent:        indent,
		Breaker:       defaultBreaker,
	})
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Printf("Failed to open browser: %v", err)
	}
}

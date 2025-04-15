//go:build !withweb
// +build !withweb

package main

// Controlled by a build tag (withweb) to include or exclude web server functionality.
var WithWeb bool = false

// startTestServer starts an HTTP listener on the specified port and opens the browser.
func startTestServer(port string, _ bool, options *FormatOptions) {
	// A stub
}

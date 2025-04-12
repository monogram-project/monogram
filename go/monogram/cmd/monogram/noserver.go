//go:build !withweb
// +build !withweb

package main

func withWeb() bool {
	return false
}

// startTestServer starts an HTTP listener on the specified port and opens the browser.
func startTestServer(port string, options *FormatOptions) {
	// A stub
}

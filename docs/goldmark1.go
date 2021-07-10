// Demonstrates the goldmark Markdown to HTML converter.
// Creates a simple Markdown source file named "foo.md".
// Calls goldmark to convert it to HTML.
// Displays results to standard output
package main

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	"io/ioutil"
	"os"
)

const sampleFile = `
# Header 1
Standard HTML paragraph
`

func main() {
	// Holds contents of the Markdown file
	var input []byte
	var err error
	// Create a sample Markdown source file
	var filename = "foo.md"
	err = ioutil.WriteFile("foo.md", []byte(sampleFile), 0644)
	if err != nil {
		// TODO: More specific error handling
		quit(err.Error(), 1)
	}
	// Read the whole file ito memory.
	input, err = ioutil.ReadFile(filename)
	if err != nil {
		// TODO: More specific error handling
		quit(err.Error(), 1)
	}
	// Convert to Markdown
	fmt.Println(string(mdFileToHTML(filename, input)))

}

func mdFileToHTML(filename string, input []byte) []byte {
	var buf bytes.Buffer
	if err := goldmark.Convert(input, &buf); err != nil {
		return []byte("")
	}
	return buf.Bytes()
}

// quit() displays the error message and exits to the operating system.
func quit(errorMsg string, exitCode int) {
	fmt.Println(errorMsg)
	os.Exit(exitCode)
}

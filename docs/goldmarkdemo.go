/* Demonstration of using Goldmark for Markdown-> HTML conversion, but with some unnecessary code */
package main

import (
	/*
	"html/template"
	"log"
	"path/filepath"
	"strings"
	*/
	"io/ioutil"
	"fmt"
	"os"
	"flag"
	"bytes"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

// Options used when converting markdown to HTML.
type markdownOptions struct {
	// Name of color scheme used for code highlighting,
	// for example, "monokai"
	highlightStyle string
}

// Compound data structure for config example at
// https://gist.github.com/alexedwards/5cd712192b4831058b21
type Env struct {
	MarkdownOptions *markdownOptions
}

func main() {
	// Obtain command-line options
	// Filenames are left in flag.Args()
	flag.Parse()
	//pHighlightStyle := flag.String("highlight", "monokai", "color theme for code highlighting")
	if len(flag.Args()) < 1 {
		quit("Please specify a Markdown file", 1)
	}

	filename := flag.Arg(0)

	// Read the whole file ito memory.
	var input []byte
	var err error
	input, err = ioutil.ReadFile(filename)
	if err != nil {
		// TODO: More specific error handling
		quit(err.Error(), 1)
	}
	fmt.Println(string(mdFileToHTML(filename, input)))

}


func mdFileToHTML(filename string, input []byte) []byte {
	// Resolve any Go template variables before conversion to HTML.
	//interp := interps(filename, string(input))

	markdown := goldmark.New(
		goldmark.WithExtensions(extension.GFM, extension.DefinitionList,
			highlighting.NewHighlighting(
				highlighting.WithStyle("monokai"),
				highlighting.WithFormatOptions(
				//highlighting.WithLineNumbers(),
				),
			)),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			/* html.WithHardWraps(), */
			html.WithXHTML(),
		),
	)

	var buf bytes.Buffer
	// TEMPLATE */
	//if err := markdown.Convert([]byte(interp), &buf); err != nil {
	if err := markdown.Convert(input, &buf); err != nil {
		return []byte("")
	}

	return buf.Bytes()
}

// quit() displays the error message and exits to the operating system.
// Only to be used in main()
func quit(errorMsg string, exitCode int) {
	fmt.Println(errorMsg)
	os.Exit(exitCode)
}



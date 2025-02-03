// md2.go Parse YAML front matter and convert Markdown to HTML.
package main

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"log"
)

var source = `---
Title: hello, world.
Description: Show a few document types 
array:
    - Homer
    - Marge 
    - Bart
    - Lisa
---
# hello, world

Redirect this output to an HTML file and open it.
`

func main() {
	mdParser := goldmark.New(
		goldmark.WithExtensions(
			meta.New(
				meta.WithStoresInDocument(),
			),
		),
	)
	mdParserCtx := parser.NewContext()
	document := mdParser.Parser().Parse(text.NewReader([]byte(source)))
	metaData := document.OwnerDocument().Meta()
	var buf bytes.Buffer
	// Convert Markdown source to HTML and deposit in buf.Bytes().
	if err := mdParser.Convert([]byte(source), &buf, parser.WithContext(mdParserCtx)); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("HTML:\n%v\n", buf.String())
	fmt.Printf("YAML front matter: %v\n", metaData)
	fmt.Printf("Title: %v\n", metaData["Title"])
	fmt.Printf("Array: %+v\n", metaData["array"])
}

/*

Run from the command line:

$ go run md2.go

Output:

<h1>hello, world</h1>
<p>Redirect this output to an HTML file and open it.</p>

Try redirecting like so:

$ go run md2.go > foo.html

Then open foo.html in web browser.

*/
ß

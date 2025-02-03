// md3.go Parse YAML front matter. Convert Markdown to HTML. Add custom function.
package main

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/text"
	"html/template"
	"log"
	"time"
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

Month: {{ ftime "Jan" }}
Default time format: {{ ftime }}
`


func ftime(param ...string) string {
    	var ref = "Mon Jan 2 15:04:05 -0700 MST 2006"
    	var format string
    	if len(param) < 1 {
    		format = ref
    	} else {
    		format = param[0]
    	}
    	t := time.Now()
    	return t.Format(format)
    }

func main() {
	var CustomFuncs template.FuncMap
	// All built-in functions must appear here to be publicly available
	CustomFuncs = template.FuncMap{
		"ftime": ftime,
	}

	mdParser := goldmark.New(
		goldmark.WithExtensions(
			meta.New(
				meta.WithStoresInDocument(),
			),
		),
	)
	document := mdParser.Parser().Parse(text.NewReader([]byte(source)))
	metaData := document.OwnerDocument().Meta()
	var buf bytes.Buffer
  t := template.Must(template.New("").Funcs(CustomFuncs).Parse(source))
  if err := t.Execute(&buf, ""); err != nil {
    log.Fatal(err)
  }


	fmt.Printf("HTML:\n%v\n", buf.String())
	fmt.Printf("YAML front matter: %v\n", metaData)
	fmt.Printf("Title: %v\n", metaData["Title"])
	fmt.Printf("Array: %+v\n", metaData["array"])
}

/*

Run from the command line:

$ go run md3.go

Output:

<h1>hello, world</h1>
<p>Redirect this output to an HTML file and open it.</p>

Try redirecting like so:

$ go run md3.go > foo.html

Then open foo.html in web browser.

*/

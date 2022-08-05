package  main // General purpose routine to convert Markdown file to HTML.

import (
	"fmt"
	"os"
	"bytes"
  "io/ioutil"
	"github.com/yuin/goldmark"
)

var defaultExample = `
# CMS example
hello, world.
`

func main() {
  if len(os.Args) < 2 {
    // No file was provided on the command line. Use defaultExample
    if HTML, err := mdToHTML([]byte(defaultExample)); err != nil {
      //quit(err.Error(), 1)
      quit(err, 1)
    } else {
      fmt.Println(string(HTML))
      quit(err, 0)
    }
  }

  filename := os.Args[1]
	if HTML, err := mdFileToHTML(filename); err != nil {
		quit(err, 1)
	} else {
	  fmt.Println(HTML)
		quit(err, 0)
	}
}
/*
package md2html

import (
	"bytes"
	"fmt"
	"os"
	"github.com/yuin/goldmark"
)
*/

// mdToHTML takes Markdown source as a byte slice and converts it to HTML
// using Goldmark's default settings.
func mdToHTML(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert(input, &buf); err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

// mdFileToHTML converts a source file to an HTML string
// using Goldmark's default settings.
func mdFileToHTML(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
  if HTML, err := mdToHTML(bytes); err != nil {
    return "", err
  } else {
    return string(HTML), nil
  }
}

func quit(err error, exitCode int) {
	if err != nil {
    fmt.Printf("%v ", err.Error())
  }
	os.Exit(exitCode)
}



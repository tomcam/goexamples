// cli2 - Minimal command-line Black Friday program
// Example front-end for command-line use.
// Give it a
// markdown file and the name of an output HTML file,
// and it generates that HTML file from the markup.
// Like cli1 but adds style sheet support
// Example:
//  ./cli2 test1.md foo.html styles1.css
//
package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Given the name of an input Markdown file and the name of
// an output HTML file, generate HTML from the input file
// and create the HTML file. Does not check to see if an existing
// HTML file exists. options are things like HTML_SKIP_STYLE (to
// skip embedded style elements, and can be found here:
// https://github.com/russross/blackfriday/blob/master/html.go
// title gets converted to the title tag,
// css is a placeholder that doesn't get used in this
// program.
//
func generateHTML(infile string, outfile string, options int, title string, css string) (err error) {
	var input []byte
	// Read the markdown file into a byte slice.
	if input, err = ioutil.ReadFile(infile); err != nil {
		return err
	}
	// Create an object to do the rendering.
	// Pass it the contents of a title tag, and CSS
	// (which in this case isn't used)
	renderer := blackfriday.HtmlRenderer(options,
		title, css)

	// Read the markdown file from the byte slice named
	// input, render to HTML, and write
	// to a byte slice named output.
	var output []byte
	output = blackfriday.Markdown(input, renderer, 0)

	// Take the rendered HTML and open up a file object.
	var out *os.File
	if out, err = os.Create(outfile); err != nil {
		return err
	}
	defer out.Close()

	// Take the generated byte slice in output and
	// create an HTML file.
	if _, err = out.Write(output); err != nil {
		return err
	}

	// Success.
	return nil
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr,
			"Please specify at 2 filenames: the first is a markdown file, the second is an HTML output file. A style sheet is optional.")
		os.Exit(-1)
	}

	// For optional style sheet
	var cssfilename string
	// Loop through command line arguments
	// If a CSS file is included, add it to the call.
	for _, arg := range os.Args {
		if filepath.Ext(arg) == ".css" {
			cssfilename = arg
		}
	}

	if err := generateHTML(os.Args[1], os.Args[2], blackfriday.HTML_COMPLETE_PAGE, "Markdown!", cssfilename); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err.Error())
	}
}

package main
/* See the go playground version at https://play.golang.org/p/xzlLy5-Pw2s 
/* Trying to get WordPress-like shortcodes working in my static site generator,
 * which, of course, converts Markdown to HTML. Since there are many things
 * Markdown doesn't understand, you can use Go template techniques.
 *
 * For example, {{now}} is a demo custom template function that returns a
 * string containing the current time and date.
 *
 * Usin gthe built-in Go template language, you can create variables
 * at runtime like this: {{$h:="hello"}}  {{$h}}, world, which would of
 * course return "hello, world".
 *
 * The shortcode custom templatel function I'm trying to implement here
 * is intended to allow you to insert arbitrary HTML into the generated
 * output, and to be able to inject parameters using Go template variables.
 * In the example below, you'd be able to pass the ID of the YouTube video
 * into the youtube.html file... if I could figure out how to parse
 * these templates at runtime. The goal is to be able to do something like
 * this to get a youtube player inserted into your HTML stream as it's
 * compiling the markdown to HTML:
 *
 *   {{$v:="tcrTQUVkUe0"}}
 *   https://youtube.com/embed/{{$v}}
 *
 * This demo runs 2 similar source files. The first shows the
 * use of the {{now}} custom function and creation of a template
 * variable at runtime. It executes properly with the expected output.
 *
 *	{{$v:="tcrTQUVkUe0"}}
 *	Video ID is: {{$v}}.
 *	Time is: {{now}}. Site Name is: {{.Site.Name}}
 *
 *  Result:
 *
 *      Video ID is: tcrTQUVkUe0.
 *      Time is: Nov 26, 2019 3:01am. Site Name is: foo
 *
 *  The second demo source is:
 *
 *      {{$v:="tcrTQUVkUe0" }}Video ID is: {{$v}}.
 *      Time is: {{now}}. Site Name is: {{.Site.Name}}
 *      {{shortcode "" }}
 *
 * The the second demo runs a similar template that adds a call to
 * shortcode, which accepts as its parameter a filename, which it opens,
 * parses, and executes. Its job is to insert the content
 * of an html file (or the youtubeHTML constant if not passed a
 * filename). To do this I figure I have to reparse the the exact same
 * source code so it can determine the value of variables and pass them
 * through to the shortcode function, but it's not working.
 * Remember that for this demo passing "" as the filename hardcodes
 * it to insert the equivalent HTML file, shown above in the comments.
 * The output on this one is simply:
 *
 *      test:4: undefined variable "$v"
 */

/* Contents of youtube.html:
<div>
<iframe
src="https://www.youtube.com/embed/{{$v}}"
allowfullscreen>
</iframe>
</div>
*/

import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"html/template"
	"io/ioutil"
	"os"
	"time"
)

type SiteConfigs struct {
	Name string
}

type Data struct {
	Site *SiteConfigs
}

var (
	// TOML file is read into this
	siteConfig SiteConfigs

	// This is the TOML file.
	config = `name="foo"`

	// Simulates a file being read in at runtime
	youtubeHTML = `
<div>       
<iframe
src="https://www.youtube.com/embed/{{$v}}"
allowfullscreen>                       
</iframe>                              
</div>                                   
`
	// A template demonstrating that variables
	// can be assigned at runtime, and that
	// custom template functions ("now") work
	// properly.
	tpl = `
	{{$v:="tcrTQUVkUe0" }}Video ID is: {{$v}}. 
	Time is: {{now}}. Site Name is: {{.Site.Name}}
	`
	// A similar template file, but this one uses
	// the shortcode custom template function, which
	// is designed to read in an HTML file at
	// runtime, parse any
	// variables or custom functions, and return
	// the parsed file.
	// In this kludged example designed to work
	// without external files so it runs on the
	// Go playground, if no filename is
	// supplied, it hardcodes the contents
	// of an html file shown as the constant
	// youtubeHTML. In real life this could be
	// any file, say a twitter.html that passes
	// the tweet via a URL variable.
	tpl2 = `
	{{$v:="tcrTQUVkUe0" }}Video ID is: {{$v}}. 
	Time is: {{now}}. Site Name is: {{.Site.Name}}
	{{shortcode "" }} 
	`

	// The compound data structure holding the site configuration
	// object. This looks like overkill but the real version
	// has several data structures that need to be accessible through
	// a template: not just the site as in this exmaple, but also
	// the page itself, and the page front matter.
	data = Data{
		Site: &siteConfig,
	}
)

// List of custom functions this app knows about.
var funcs = template.FuncMap{"now": now, "shortcode": shortcode}

// List of custom functions this app knows about except for
// shortcode, because including it causes a cycle condition and
// it won't compile. So when the shortcode function is evaluated
// at runtime, it re-reads the same template but without
// shortcode in the function map.
var fewerFuncs = template.FuncMap{"now": now}

// Return contents of an HTML file,
// but first parse it as a Golang template
// so it can receive values of variables or
// custom functions.
func shortcode(filename string) template.HTML {
	// Read the HTML file into a byte slice.
	var input []byte
	var err error
	if filename == "" {
		input = []byte(youtubeHTML)
	} else {
		input, err = ioutil.ReadFile(filename)
		if err != nil {
			return template.HTML("")
		}
	}
	// Apply the template to it.
	// The one function missing from fewerFuncs is shortcode() itself.
	if s, err := execute("test", string(input), data, fewerFuncs); err != nil {
		quit(err, 4)
	} else {
		return template.HTML(s)
	}
	return template.HTML(string(input))
}

// Demo simply to prove that custom functions work
// in templates.
func now() string {
	//Mon Jan 2 15:04
	return fmt.Sprintf("%v", time.Now().Format("Jan 2, 2006 3:04pm"))
}

// Display any error message and exit to OS.
func quit(err error, code int) {
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error: %v\n", err.Error())
	}
	fmt.Fprintf(os.Stdout, "Quitting with error code %v\n", code)
	os.Exit(code)
}

// Parse a template, then execute it against HTML/template source.
// Return a string containing the result.
func execute(templateName string, tpl string, theData Data, funcMap template.FuncMap) (buf string, err error) {
	if t, err := template.New(templateName).Funcs(funcMap).Parse(tpl); err != nil {
		return "", err
	} else {
		var b bytes.Buffer
		err := t.ExecuteTemplate(&b, templateName, theData)
		if err != nil {
			return "", err
		}
		return b.String(), nil
	}
}

func main() {
	// Read the TOML site configuration file. It has only
	// a single demo, which is the name of the site.
	if _, err := toml.Decode(config, &siteConfig); err != nil {
		quit(err, 1)
	}

	// Execute and display results from the first template,
	// which doesn't use shortcode.
	if s, err := execute("tpl", tpl, data, funcs); err != nil {
		quit(err, 4)
	} else {
		fmt.Println(s)
	}

	// Execute a similar template, but this one reads in a file
	// using shortcode and tries to apply variables against it.
	// This is done by reading in the exact same config info
	// again, but shortcode runs a FuncMap that omits the
	// shortcode function itself to avoid cycles.
	if s, err := execute("tpl2", tpl2, data, funcs); err != nil {
		quit(err, 4)
	} else {
		fmt.Println(s)
	}
	return
}

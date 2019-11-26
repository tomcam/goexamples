package main
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
	"os"
	"fmt"
	"github.com/BurntSushi/toml"
	"html/template"
	"time"
	"io/ioutil"
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
	tpl= `
	{{$v:="tcrTQUVkUe0"}}

	Video ID is: {{$v}}. 

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
	{{$v:="tcrTQUVkUe0"}}

	{{shortcode "" }} 
	
	Video ID is: {{$v}}. 

	Time is: {{now}}. Site Name is: {{.Site.Name}}
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
var funcs = template.FuncMap{"now": now, "shortcode": shortcode }

// List of custom functions this app knows about except for
// shortcode, because including it causes a cycle condition and
// it won't compile. So when the shortcode function is evaluated
// at runtime, it re-reads the same template but without
// shortcode in the function map.
var fewerFuncs = template.FuncMap{"now": now  }

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
        return fmt.Sprintf("%v", time.Now())
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


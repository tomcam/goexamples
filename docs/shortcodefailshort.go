package main
// https://play.golang.org/p/wl2bNTILWWL
// Identical to but with fewer comments
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
	// Template demonstration 1
	tpl = `
	{{$v:="tcrTQUVkUe0" }}Video ID is: {{$v}}. 
	Time is: {{now}}. Site Name is: {{.Site.Name}}
	`
	// Template demonstration 2
	tpl2 = `
	{{$v:="tcrTQUVkUe0" }}Video ID is: {{$v}}. 
	Time is: {{now}}. Site Name is: {{.Site.Name}}
	{{shortcode "" }} 
	`
	data = Data{
		Site: &siteConfig,
	}
)

// List of custom functions this app knows about.
var funcs = template.FuncMap{"now": now, "shortcode": shortcode}

// List of custom functions this app knows about except for
// shortcode
var fewerFuncs = template.FuncMap{"now": now}

// Return contents of an HTML file,
// but first parse it as a Golang template
// so it can receive values of variables or
// custom functions.
func shortcode(filename string) template.HTML {
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
	if s, err := execute("tpl2", tpl2, data, funcs); err != nil {
		quit(err, 4)
	} else {
		fmt.Println(s)
	}
	return
}

package main
/* tmplfunction.go shows how to add a custom function to a Go HTML template. */
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

var funcs = template.FuncMap{"now": now }

func shortcode(filename string) template.HTML {
	// Return contents of an HTML file
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		return template.HTML("")
	}
	return template.HTML(string(input))
}

func now() string {
        return fmt.Sprintf("%v", time.Now())
}

var (
	siteConfig SiteConfigs
	config = `name="foo"`
	tpl = `Time is: {{now}}. Site Name is: {{.Site.Name}}`
)

func quit(err error, code int) {
	if err != nil {
		fmt.Fprintf(os.Stdout, "Error: %v\n", err.Error())
	}
	fmt.Fprintf(os.Stdout, "Quitting with error code %v\n", code)
}

func main() {
	data := struct {
		Site *SiteConfigs
	}{&siteConfig}

	if _, err := toml.Decode(config, &siteConfig); err != nil {
		quit(err, 1)
	}

	if t, err := template.New("test").Funcs(funcs).Parse(tpl); err != nil {
		quit(err, 2)
	} else {
		var b bytes.Buffer
		err := t.ExecuteTemplate(&b, "test", data)
		if err != nil {
			fmt.Printf("ERROR b.String(): %v\n", b.String())
			quit(err, 3)
		}
		fmt.Printf("%v\n", b.String())
	}
}


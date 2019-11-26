package main

/* Go Playground: https://play.golang.org/p/HG8j17pVY0Q */
import (
	"bytes"
	"fmt"
	"github.com/BurntSushi/toml"
	"html/template"
)

type SiteConfigs struct {
	Name string
}

var (
	siteConfig SiteConfigs
	config     = `name="foo"`
	tpl        = `Site name is: {{.Site.Name}}`
)

func quit(err error, code int) {
	if err != nil {
		fmt.Printf("Error: %v. ", err.Error())
	}
	fmt.Printf("Exit code %v\n", code)
}

func main() {
	data := struct {
		Site *SiteConfigs
	}{&siteConfig}

	if _, err := toml.Decode(config, &siteConfig); err != nil {
		quit(err, 1)
	}

	if t, err := template.New("test").Parse(tpl); err != nil {
		quit(err, 2)
	} else {
		var b bytes.Buffer
		err := t.ExecuteTemplate(&b, "test", data)
		if err != nil {
			quit(err, 3)
		}
		fmt.Println(b.String())
	}
}

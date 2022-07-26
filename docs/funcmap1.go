// Shows how to add multiple user defined Go template functions using a function map
package main

import (
	"html/template"
	"os"
	"time"
)

var funcs = template.FuncMap{
	"ftime":    ftime,
	"hostname": hostname,
}

const templ = `Default ftime: {{ ftime }}

Month and date: {{ ftime "January 2" }}
 
Host name: {{ hostname }}
`

func main() {
	t := template.Must(template.New("").Funcs(funcs).Parse(templ))
	if err := t.Execute(os.Stdout, ""); err != nil {
		panic(err)
	}
}

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

func hostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return ""
	} else {
		return hostname
	}
}


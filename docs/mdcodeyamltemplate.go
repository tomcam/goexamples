// mdcodeyamltemplate.go:
// Goldmark
// Demonstrates
// 1. The goldmark Markdown to HTML converter using an App object
// 2. Code highlighting.
// 3. Extracting YAML front matter
// 4. Executing a template to interpolate front matter metadata with its evaluated result
// 5. Adding a custom template function

// $ mkdir ~/g
// $ cd ~/g
// $ go mod init example.com/g # example.com is OK to use in this quick & dirty example
// $ go fmt
// $ go mod tidy
// $ go run g.go
// $ go run g.go > foobar.html
// $ open foobar.html

package main

import (
	"bytes"
	"fmt"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark-highlighting"
	"github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer"
	"github.com/yuin/goldmark/renderer/html"
	"os"
	"text/template"
	"time"
)

const title = `
## Title: {{ .Title }}
`

const frontMatter = `---
Title: goldmark-meta
Month: January
Theme: wide
Summary: Add YAML metadata to the document
Tags:
    - markdown
    - goldmark
---
`

const codeFence = "\nhello, world.\n\n### Code fence with highlighting:\n" +
	"Code fence example:\n" +
	"```js\n" +
	"console.log('hi')\n" +
	"```\n" + `
  `

const ftimeExample = `
## User-defined time function test
The Month is: {{ ftime .Month }}

Fully formatted date: {{ ftime }}
`

type App struct {
	mdParser    goldmark.Markdown
	mdParserCtx parser.Context

	// YAML front matter
	metaData map[string]interface{}

	// All built-in functions must appear here to be publicly available
	funcs map[string]interface{}
}

func (app *App) addTemplateFunctions() {
	app.funcs = template.FuncMap{
		/*
		   "article":  a.articlefunc,
		   "dirnames": a.dirNames,
		   "files":    a.files,
		*/
		"ftime": app.ftime,
		"quote": app.quote,
		/*
		   "hostname": a.hostname,
		   "inc":      a.inc,
		   "path":     a.path,
		   "scode":    a.scode,
		   "toc":      a.toc,
		*/
	}
}

// newGoldmark returns the a goldmark object with a parser and renderer.
func (app *App) newGoldmark() goldmark.Markdown {
	exts := []goldmark.Extender{
		meta.New(
			meta.WithStoresInDocument(),
		),
		// Support GitHub tables & other extensions
		extension.Table,
		extension.GFM,
		extension.DefinitionList,
		extension.Footnote,
		highlighting.NewHighlighting(
			highlighting.WithStyle("github"),
			highlighting.WithFormatOptions()),
	}

	parserOpts := []parser.Option{
		parser.WithAttribute(),
		parser.WithAutoHeadingID()}

	renderOpts := []renderer.Option{
		// WithUnsafe is required for HTML templates to work properly
		html.WithUnsafe(),
		html.WithXHTML(),
	}
	return goldmark.New(
		goldmark.WithExtensions(exts...),
		goldmark.WithParserOptions(parserOpts...),
		goldmark.WithRendererOptions(renderOpts...),
	)
}

func NewApp() *App {
	app := App{}

	app.mdParser = app.newGoldmark()
	app.mdParserCtx = parser.NewContext()
	app.addTemplateFunctions()
	return &app
}

func mdTest() {
	var app = NewApp()
	if b, err := app.mdToHTML([]byte(frontMatter +
		"# Markdown to HTML. No front matter support\n" +
		title +
		codeFence)); err != nil {
		panic("mdTest()")
		quit(err, 1)
	} else {
		fmt.Println(string(b))
		fmt.Printf("Front matter as YAML: %v\n", app.metaData)
	}
}
func mdYAMLTest() {
	var app = NewApp()
	if b, err := app.mdYAMLToHTML([]byte(frontMatter +
		"# Markdown to HTML with front matter parsed\n" +
		title +
		codeFence)); err != nil {
		panic("mdYAMLToHTML()")
		quit(err, 1)
	} else {
		fmt.Println(string(b))
		fmt.Printf("Front matter as YAML: %v\n", app.metaData)
	}
}

func mdYAMLTemplateTest() {
	var app = NewApp()
	var err error
	var b []byte
	if b, err = app.mdYAMLToHTML([]byte(frontMatter +
		"# Markdown to HTML with front matter parsed and executed in template\n" +
		title +
		codeFence)); err != nil {
		panic("mdYAMLTemplateTest()")
	}
	var t string
	t = app.doTemplate("METABUZZ", string(b))
	fmt.Println(t)
}

func mdYAMLTemplateFuncTest() {
	var app = NewApp()
	var err error
	var b []byte
	if b, err = app.mdYAMLToHTML([]byte(frontMatter +
		"# Markdown to HTML with front matter parsed and executed in template, plus a custom template function\n" +
		title +
		ftimeExample +
		codeFence)); err != nil {
		panic("mdYAMLTemplateFuncTest()")
	}
	var t string
	if t, err = app.doTemplateFuncs("METABUZZ", string(b)); err != nil {
		panic("mdYAMLTemplateFuncTest()")
	}
	fmt.Println(t)
}

// mdYAMLtoHTML converts a Markdown document with optional
// YAML front matter to HTML. YAML is written to app.metaData
// Returns a byte slice containing the HTML source.
// Pre: parser.NewContext() has already been called on app.parserCtx
func (app *App) mdYAMLToHTML(source []byte) ([]byte, error) {
	var buf bytes.Buffer
	// Convert Markdown source to HTML and deposit in buf.Bytes().
	if err := app.mdParser.Convert(source, &buf, parser.WithContext(app.mdParserCtx)); err != nil {
		return []byte{}, err
	}
	// Obtain YAML front matter from document.
	app.metaData = meta.Get(app.mdParserCtx)
	return buf.Bytes(), nil
}

// mdYAMLtoHTMLStr converts a Markdown document with optional YAML front matter to HTML. YAML is written to app.metaData
// Returns a string containing the HTML source.
// Pre: parser.NewContext() has already been called on app.parserCtx
func (app *App) mdYAMLToHTMLStr(source []byte) (string, error) {
	var buf bytes.Buffer
	// Convert Markdown source to HTML and deposit in buf.Bytes().
	if err := app.mdParser.Convert(source, &buf, parser.WithContext(app.mdParserCtx)); err != nil {
		return "", err
	}
	// Obtain YAML front matter from document.
	app.metaData = meta.Get(app.mdParserCtx)
	return string(buf.Bytes()), nil
}

// mdtoHTML converts a Markdown document to HTML.
// YAML front matter should not be present.
// Returns a byte slice containing the HTML source.
// Pre: parser.NewContext() has already been called on app.parserCtx
func (app *App) mdToHTML(source []byte) ([]byte, error) {
	var buf bytes.Buffer
	// Convert Markdown source to HTML and deposit in buf.Bytes().
	if err := app.mdParser.Convert(source, &buf, parser.WithContext(app.mdParserCtx)); err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

// doTemplate takes HTML in source, expects parsed front
// matter in app.metaData, and executes Go templates
// against the source.
// Returns a string containing the HTML with the
// template values embedded.
func (app *App) doTemplate(templateName string, source string) string {
	if templateName == "" {
		templateName = "Metabuzz"
	}
	tmpl, err := template.New(templateName).Parse(source)
	if err != nil {
		quit(err, 1)
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, app.metaData)

	if err != nil {
		quit(err, 1)
	}
	return buf.String()

}

// doTemplateFuncs takes HTML in source, expects parsed front
// matter in app.metaData, and executes Go templates
// against the source. It also handles user-defined
// functions, expected in funcMap
// Returns a string containing the HTML with the
// template values embedded.
func (app *App) doTemplateFuncs(templateName string, source string) (string, error) {
	if templateName == "" {
		templateName = "Metabuzz"
	}
	var tmpl *template.Template
	var err error
	if tmpl, err = template.New(templateName).Funcs(app.funcs).Parse(source); err != nil {
		// TODO: Function should return error
		return "", err
	}
	buf := new(bytes.Buffer)
	err = tmpl.ExecuteTemplate(buf, templateName, app.metaData)

	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func main() {

	// Simple Markdown to HTML. No front matter support
	mdTest()

	// Markdown to HTML with YAML front matter parsed
	mdYAMLTest()

	// Markdown to HTML with YAML front matter parsed and executed in template
	mdYAMLTemplateTest()

  // Markdown to HTML with front matter parsed and executed in template, plus a custom template function
	mdYAMLTemplateFuncTest()

}

// ftime() returns the current, local, formatted time.
// Can pass in a formatting string
// https://golang.org/pkg/time/#Time.Format
func (app *App) ftime(param ...string) string {
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

// quote
func (app *App) quote(param string) string {
	return param
}

func quit(err error, exitCode int) {
	if err != nil {
		fmt.Printf("Quitting with error: %v\n", err)
	}
	os.Exit(exitCode)
}

# Example programs in Go

All of these example Go programs are complete. You can run most or all of them on the Go Playground.

## Command-line flags/CLI
* [flagbool.go](flagbool.go) Illustrates output of the simplest possible command-line boolean flag using the [flag](https://pkg.go.dev/flag) package
* [cmdline](cmdline.go) shows how to parse the command line, with optional subcommands like "init" or "init sitename=test", and also how to retrieve those values later

## Directory tree
* [dl.go](dl.go) - dl lists all directories in the specified path (defaults to current directory). Gist at https://gist.github.com/tomcam/db640cfc7b846e083f4f7dec19cf345d
* [dirtree.go](dirtree.go) - Show directory tree as string slice. Allow files & dirs to be excluded
* [lastmodified.go](lastmodified.go) - Reusable code (and demo) showing how to retrieve the last modified date of a file by filename
* [genfiletree.go](genfiletree.go) - Generate small tree of text files, e.g. for a website. Gist at https://gist.github.com/tomcam/9766ca1ea4d63eff804335c5f391ab09
* [filewatch1.go](filewatch1.go) - Short demo of radovskyb's recursive file watching package. Gist at https://gist.github.com/tomcam/32760a5049a00ec6ba82bcb42b6759fc
* [filewatch2.go](filewatch2.go) - Short demo of radovskyb's recursive file watching package with different
order of execution. Gist at [https://gist.github.com/tomcam/32760a5049a00ec6ba82bcb42b6759fc](https://gist.github.com/tomcam/97c75f7706d4763732018a3429a020d3)

## embed (embedding data in a Go executable)
* [embedlistdir.go](embedlistdir.go) creates an executable that embeds a directory, then displays the filenames in that directory at runtime

## File handling
* [lastmodified.go](lastmodified.go) - Reusable code (and demo) showing how to retrieve the last modified date of a file by filename
* [genfiletree.go](genfiletree.co) - Generate small tree of text files, e.g. for a website. Gist at https://gist.github.com/tomcam/9766ca1ea4d63eff804335c5f391ab09
 
## Reflection/runtime type identification
* [fieldisstringtype.go](fieldisstringtype.go) determines at runtime whether the struct passed in the argument has a field named in the second argument of type string. Playground version at https://play.golang.org/p/yAEXeeCvJMH
* [structfieldbynamestrmust.go](structfieldbynamestrmust.go) takes any struct and field name (as a string) at runtime and
returns the string value of that field. It returns an empty string if the 
object passed in isn't a struct, or if the named field isn't a struct. Playground version at https://play.golang.org/p/MiCn6NtEp5-
* [structinfo.go](structinfo.go): `structInfo()` takes any struct at runtime and displays its type name, field names and types, 
and contents of each field. `structHasField()` returns true if a struct passed to it at runtime contains a field name passed as a string. Playground version at https://play.golang.org/p/zeOTNfHEQlH

## JSON
* [jsonstruct.go](jsonstruct.go) - Initialize nested struct in Golang. Read and write nested structs to a JSON file in Go. See also [jconstruct on Go playground](https://play.golang.org/p/S7HbAOk0ZDb)


## Markdown: Goldmark
* [md1.go: Minimal Goldmark to convert from Markdown to HTML](md1.go) [Go playground](https://go.dev/play/p/5BhsP-osFKs), [gist](https://go.dev/play/p/5BhsP-osFKs)
* [md2.go][Go Playgrund)(https://go.dev/play/p/CKm5Ik-Ti0V)
* [md2htmltemplates.go](md2htmltemplates.go) Demonstrates using progressive, self-contained functions the goldmark Markdown to HTML converter using an App object, code highlighting. extracting YAML front matter, executing a template to interpolate front matter metadata with its evaluated result, and adding a custom template function. [Go Playground](https://go.dev/play/p/PQ6AxAb09kx) version, [Gist](https://gist.github.com/tomcam/9bc1d8637eb2e8ee59b0f7d2674efb7c)
* [Gist with simplest Goldmark demo](https://gist.github.com/tomcam/942342f301c78a20457c0b2e752bbb2b) Gist with simplest Goldmark demo.)
* [microcms](microcmsnoyaml.go) A one-file Markdown to HTML converter. No front matter support.
* [goldmark converter using an App object.](https://gist.github.com/tomcam/063430a32e40979736cf78bf172c42d9)  See [playground version](https://go.dev/play/p/5UpB0Z5L_EZ) or https://go.dev/play/p/XNsZD6bqIXJ
* [Goldmark demo with with App object, Markdown to HTML conversion, code highlighting, YAML front matter support, and template support with custom template functions](mdcodeyamltemplate.go), gist [here](https://gist.github.com/tomcam/70dd62c9fa36032506fc406db9b89062), go Playground version [here](https://go.dev/play/p/4c5PPHFG85C)
* [md2rawhtml](md2rawhtml.go) Smallest general-purpose micro CMS that converts a Markdown to a raw HTML file with no head, html tags, etc.
* [Goldmark demo with App object Markdown to HTML conversion, code highlighting, YAML support, simple template support](https://gist.github.com/tomcam/a1c8fbe27a335164add3bc2b1d92b204), playground version [here](https://go.dev/play/p/Xu1ELDgl4ec)
* [goldmark1.go](goldmark1.go) Simplest example showing how to convert Markdown file to HTML using Goldmark

## Markdown: BurntSushi and BlackFriday (out od date)
* [burntsushitest.go](burntsushitest.go) - Use burntsushi to read a TOML file
* [blackfriday](blackfriday/) - Simple command line programs shoing how to use the Blackfriday parser convert Markdown in to HTML. cli2 also shows how to append file extensions and parse flags

## Regex
* [regexreplace.go](regexreplace.go) - Replaces a Go template identifier like {{.Name}} with an arbitrary string.
* [Extracting HTML headers using go regexp](https://gist.github.com/tomcam/996e9e565fc8db4ca41484a369338993)

## Slices
* [dirtree.go](dirtree.go) - Show directory tree as string slice. Allow files & dirs to be excluded

## Strings
* [dirtree.go](dirtree.go) - Show directory tree as string slice. Allow files & dirs to be excluded

## Templates
* [tmplfunction.go](tmplfunction.go) shows how to add a custom function to a Go HTML template. Followup to [cfgfile.go](cfgfile.go).
* [funcmap1.go](funcmap1.go) Shows how to add multiple custom template functions (Playground version [here](https://go.dev/play/p/BqkWiQ2v7Tj)


## TOML
* [cfgfile.go](cfgfile.go) - Store config in TOML file, then show its value in an HTML go HTML template. [tmplfunction.go](tmplfunction.go) builds on it by adding a custom function.
* [tomlperplatform.go](tomlperplatform.go) creates a TOML config file that sniffs out the host OS so you can write values specific to it
* [tomlgen.go](tomlgen.go) creates a general purpose TOML file that lets you define sections/kv pairs at will
* General purpose TOML read/write using BurntSushi (gist): https://play.golang.org/p/LgZwOT363sZ
* Simple example of reading a TOML file using BurntSushi (playground): https://play.golang.org/p/klLI41DiwqC

## YAML
* [yamlreadwritestruct.go](yamlreadwritestruct.go), playground version [here](https://play.golang.org/p/KglpI_JmqSE)
* [structtomap1.go](https://gist.github.com/tomcam/3a0002119d60435505bff426b9345ae7) (gist) Go/Golang, playground version [here](https://play.golang.org/p/t8XaP2eMPPE): Writes a struct to a YAML file, then reads it back from the YAML as a map using the gopkg.in/yaml package

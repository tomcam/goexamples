# Example programs in Go

All of these example Go programs are complete. You can run most or all of them on the Go Playground.

## Command-line flags/CLI
* [flagbool.go](flagbool.go) Illustrates output of the simplest possible command-line boolean flag using the [flag](https://pkg.go.dev/flag) package
* [cmdline](cmdline.go) shows how to parse the command line, with optional subcommands like "init" or "init sitename=test", and also how to retrieve those values later

## Directory tree
* [dirtree.go](dirtree.go) - Show directory tree as string slice. Allow files & dirs to be excluded
* [lastmodified.go](lastmodified.go) - Reusable code (and demo) showing how to retrieve the last modified date of a file by filename

## embed (embedding data in a Go executable)
* [embedlistdir.go](embedlistdir.go) creates an executable that embeds a directory, then displays the filenames in that directory at runtime

## File handling
* [lastmodified.go](lastmodified.go) - Reusable code (and demo) showing how to retrieve the last modified date of a file by filename

## Reflection/runtime type identification
* [fieldisstringtype.go](fieldisstringtype.go) determines at runtime whether the struct passed in the argument has a field named in the second argument of type string. Playground version at https://play.golang.org/p/yAEXeeCvJMH
* [structfieldbynamestrmust.go](structfieldbynamestrmust.go) takes any struct and field name (as a string) at runtime and
returns the string value of that field. It returns an empty string if the 
object passed in isn't a struct, or if the named field isn't a struct. Playground version at https://play.golang.org/p/MiCn6NtEp5-
* [structinfo.go](structinfo.go): `structInfo()` takes any struct at runtime and displays its type name, field names and types, 
and contents of each field. `structHasField()` returns true if a struct passed to it at runtime contains a field name passed as a string. Playground version at https://play.golang.org/p/zeOTNfHEQlH

## JSON
* [jsonstruct.go](jsonstruct.go) - Initialize nested struct in Golang. Read and write nested structs to a JSON file in Go. See also [jconstruct on Go playground](https://play.golang.org/p/S7HbAOk0ZDb)


## Markdown
* [g.go](https://gist.github.com/tomcam/942342f301c78a20457c0b2e752bbb2b) Gist with simplest Goldmark demo. See also 
* [goldmark1.go](goldmark1.go) Simplest example showing how to convert Markdown file to HTML using Goldmark
* [goldmarkdemo.go](goldmarkdemo.go) Convert Markdown file to HTML using Goldmark, but with some unnecessary code
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


## TOML
* [cfgfile.go](cfgfile.go) - Store config in TOML file, then show its value in an HTML go HTML template. [tmplfunction.go](tmplfunction.go) builds on it by adding a custom function.
* [tomlperplatform.go](tomlperplatform.go) creates a TOML config file that sniffs out the host OS so you can write values specific to it
* [tomlgen.go](tomlgen.go) creates a general purpose TOML file that lets you define sections/kv pairs at will
* General purpose TOML read/write using BurntSushi (gist): https://play.golang.org/p/LgZwOT363sZ
* Simple example of reading a TOML file using BurntSushi (playground): https://play.golang.org/p/klLI41DiwqC

## YAML
* [yamlreadwritestruct.go](yamlreadwritestruct.go), playground version [here](https://play.golang.org/p/KglpI_JmqSE)
* [structtomap1.go](https://gist.github.com/tomcam/3a0002119d60435505bff426b9345ae7) (gist) Go/Golang, playground version [here](https://play.golang.org/p/t8XaP2eMPPE): Writes a struct to a YAML file, then reads it back from the YAML as a map using the gopkg.in/yaml package

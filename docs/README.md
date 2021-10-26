# Example programs in Go

* [structtomap1.go](https://gist.github.com/tomcam/3a0002119d60435505bff426b9345ae7) (gist) Go/Golang: Go/Golang: writes a struct to a YAML file, then reads it back from the YAML as a map using the gopkg.in/yaml package
* [embedlistdir.go](embedlistdir.go) creates an executable that embeds a directory, then displays the filenames in that directory at runtime
* [tomlperplatform.go](tomlperplatform.go) creates a TOML config file that sniffs out the host OS so you can write values specific to it
* [tomlgen.go](tomlgen.go) creates a general purpose TOML file that lets you define sections/kv pairs at will
* [suggestcfgdir.go](suggestcfgdir.go) sniffs the current operating system and makes an educated guess about where that system expects user configuration data to be stored
* [flagbool.go](flagbool.go) Illustrates output of the simples possible command-line boolean flag
* [tmplfunction.go](tmplfunction.go) shows how to add a custom function to a Go HTML template. Followup to [cfgfile.go](cfgfile.go).
* [cfgfile.go](cfgfile.go) - Store config in TOML file, then show its value in an HTML go HTML template. [tmplfunction.go](tmplfunction.go) builds on it by adding a custom function.
* [g.go](https://gist.github.com/tomcam/942342f301c78a20457c0b2e752bbb2b) Gist with simplest Goldmark demo. See also 
* [goldmark1.go](goldmark1.go) Simplest example showing how to convert Markdown file to HTML using Goldmark
* [goldmarkdemo.go](goldmarkdemo.go) Convert Markdown file to HTML using Goldmark, but with some unnecessary code
* [dirtree.go](dirtree.go) - Show directory tree as string slice. Allow files & dirs to be excluded
* [jsonstruct.go](jsonstruct.go) - Initialize nested struct in Golang. Read and write nested structs to a JSON file in Go. See also [jconstruct on Go playground](https://play.golang.org/p/S7HbAOk0ZDb)
* [burntsushitest.go](burntsushitest.go) - Use burntsushi to read a TOML file
* [blackfriday](blackfriday/) - Simple command line programs shoing how to use the Blackfriday parser convert Markdown in to HTML. cli2 also shows how to append file extensions and parse flags
* [cmdline](cmdline.go) shows how to parse the command line, with optional subcommands like "init" or "init sitename=test", and also how to retrieve those values later
* [lastmodified.go](lastmodified.go) - Reusable code (and demo) showing how to retrieve the last modified date of a file by filename
* [regexreplace.go](regexreplace.go) - Replaces a Go template identifier like {{.Name}} with an arbitrary string.
* [Extracting HTML headers using go regexp](https://gist.github.com/tomcam/996e9e565fc8db4ca41484a369338993)

## Golang Playgrounds and Gists
* General purpose TOML read/write using BurntSushi: https://play.golang.org/p/LgZwOT363sZ
* Simple example of reading a TOML file using BurntSushi: https://play.golang.org/p/klLI41DiwqC

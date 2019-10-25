package main
import (
	"fmt"
	"io/ioutil"
    "github.com/BurntSushi/toml"
	"os"
)


// docset looks for this file, to figure out what,
// for example, is the current theme.
type tomlDefaults struct {
  // What theme is in use? If empty, the Default theme of course
  CurrentTheme string

  // Base directory for blog, which may be diffferent
  // from its root. For example, GitHub Pages prefers
  // the blog to start in /docs instead of root, but
  // a URL would omit i.
  BaseDir string
  
}

// A theme is made up of 1 or more page types.
// The default page type is just page and will be supplied if dir is tempy.

// Theme directory:
// Can be empty.
// PageLayouts subdirectory contains subdirectories that name page
// types, say Blog, Article, or Home
// Each one contains standardized files named header.html, footer.html, aside.html, section.html, article.html, etc
// A TOML file entry determines if they're used
// [article]
// 


// * stylesheet(s) with no classes
// This includes the root, aka default theme directory

// Example:
// https://github.com/BurntSushi/toml/blob/master/_examples/example.go
// Overall file format
type tomlTheme struct {
  Stylesheets []string

}

// Partial could be, say, a header:
// html is inline html. filename would be a pathname containing the HTML.
type partial struct {
  Name string
  Html string
  Filename string
}

type layout struct {
  Name string
}

func readTOMLfile(infile string) (err error) {
  var input[]byte
  if input, err = ioutil.ReadFile(infile); err != nil {
    return err
  }

  var deflayout layout
  if _, err = toml.Decode(string(input),&deflayout); err != nil {
    return err
  }
  fmt.Fprintf(os.Stdout, "Default config file: %v", defconfig)
  fmt.Fprintf(os.Stdout, "Layout: %v", deflayout)
  return nil
}
func readConfigFile(infile string) (err error) {
  var input[]byte
  if input, err = ioutil.ReadFile(infile); err != nil {
    return err
  }

  var deflayout layout
  if _, err = toml.Decode(string(input),&deflayout); err != nil {
    return err
  }
  fmt.Fprintf(os.Stdout, "Default config file: %v", defconfig)
  fmt.Fprintf(os.Stdout, "Layout: %v", deflayout)
  return nil
}

func main() {
  if err := readTOMLfile("foo.txt"); err != nil {
    panic(err)
  }
}

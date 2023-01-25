// Shows how to embed a subdirectory of the package
// named .config into the executable at compile time, then get
// access to its contents at runtime.

package main
import (
	"embed"
	"io/fs"
	"fmt"
 	"io/ioutil"
 	"os"
)


// The following embeds all files and subdirectories
// from the themes subdirectory of this package into
// the executable. Can change directory name from .config to anything else.
//go:embed all:.config
var configFiles embed.FS

// embedListDir() displays the filenames in the embedded 
// directory passed to the files parameter.
func embedListDir(files embed.FS) error {
	fs.WalkDir(files, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		print(path)
		return nil
	})
	return nil
}

// Open file name and send its contents to stdout.
func print(filename string) {
  f, err := os.Open(filename); 
  if err != nil {
    fmt.Println("Error opening %s: %v", filename, err.Error())
    os.Exit(1)
  }
  bytes, err := ioutil.ReadAll(f); 
  if err != nil {
    fmt.Println("Error reading %s: %v", filename, err.Error())
  } else {
    fmt.Println(string(bytes))
    os.Exit(1)
  }
}

func main() {
	embedListDir(configFiles)
}

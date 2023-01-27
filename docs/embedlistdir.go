// Thanks to Cerise Limon (https://stackoverflow.com/questions/75251998/failing-to-understand-go-embed#75252129)
// for the essential fix: https://stackoverflow.com/questions/75251998/failing-to-understand-go-embed#75252129
package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

// A populated subdirectory directory named .config is required
//go:embed all:.config
var configFiles embed.FS

func main() {
	ls(configFiles, ".")
}

func ls(files embed.FS, dir string) error {
	fs.WalkDir(files, dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			show(files, path) // Display contents of file
			filename := filepath.Join(path, d.Name())
			fmt.Println(filename)
		}
		return nil
	})
	return nil
}

func show(files embed.FS, filename string) {
	f, err := files.Open(filename)
	if err != nil {
		quit(filename, err)
	}
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		quit(filename, err)
	}
	fmt.Println(string(bytes))
}

func quit(filename string, err error) {
	fmt.Printf("Error for file %s: %v\n", filename, err.Error())
	os.Exit(1)
}

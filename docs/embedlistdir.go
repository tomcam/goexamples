// embedlistdir() shows how to embed a subdirectory of the package
// named themes into the executable at compile time, then get
// access to its contents at runtime.

package main

import (
	"embed"
	"fmt"
	"io/fs"
)


// The following embeds all files and subdirectories
// from the themes subdirectory of this package into
// the executable.

//go:embed themes/*
var themeFiles embed.FS

// embedListDir() displays the filenames in the embedded 
// directory named theme.
func embedListDir(files embed.FS) error {
	fs.WalkDir(files, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(path)
		return nil
	})
	return nil
}
func main() {
	embedListDir(themeFiles)
}

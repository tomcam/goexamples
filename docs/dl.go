// dl displays a list of directories. Starts at current dir unless
// another is given on the command line. Counts theme.
// Naively written; doesn't do a good job when you don't have 
// the right permissions.
// Usage:
//   dl [directory]

package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func visit(files *[]string, count *int) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
      fmt.Printf("Error at %v: %v", path,  err.Error())
			return err
		}
		isDir := info.IsDir()
		if isDir {
			if path != "." && path != ".." {
				fmt.Printf("%s\n", path)
				*count++
				*files = append(*files, path)
			}
		}
		return nil
	}
}

// Obtain a list of all files in the specified project tree starting
// at the root.
// Ignore directories starting with a .
// Ignore the assets directory
func getDirTree(path string, dirs *int) (tree []string, err error) {
	var files []string
	err = filepath.Walk(path, visit(&files, dirs))
	if err != nil {
    fmt.Printf("Error at %v: %v", path,  err.Error())
		return []string{}, err
	}
	return files, nil
}

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	dirCount := 0
	_, err := getDirTree(dir, &dirCount)
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
	} else {
		if dirCount != 1 {
      fmt.Printf("%v directories\n", dirCount)
    } else {
      fmt.Printf("%v directory\n", dirCount)
    }
	}
}

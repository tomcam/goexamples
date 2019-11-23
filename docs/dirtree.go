// Obtains the directory tree as a string slice.
// Allws you to to choose what files to exlcude.
package main

import (
	"strings"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

/* Searching a sorted slice is fast.
   This tracks whether the slice has been sorted
   and sorts it on first search.
*/
type searchInfo struct {
	list   []string
	sorted bool
}

func (s *searchInfo) Sort() {
	sort.Slice(s.list, func(i, j int) bool {
		s.sorted = true
		return s.list[i] <= s.list[j]
	})
}

func (s *searchInfo) Found(searchFor string) bool {
	if !s.sorted {
		s.Sort()
	}
	var pos int
	l := len(s.list)
	pos = sort.Search(l, func(i int) bool {
		return s.list[i] >= searchFor
	})
	return pos < l && s.list[pos] == searchFor

}



func visit(files *[]string) filepath.WalkFunc {
	var exclude searchInfo
	// Find out what directories to exclude
	exclude.list = []string{"node_modules", "main.bak", ".git", "pub", ".DS_Store", ".gitignore"}
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Quietly fail if unable to access path.
			return err
		}
		isDir := info.IsDir()

		name := filepath.Base(info.Name())

		// Skip any directory to be excluded, such as
		// the pub and .git directores 
		if exclude.Found(name) && isDir {
			return filepath.SkipDir

		}
		// It may be just a filename on the exclude list.
		if exclude.Found(name) {
			return  nil

		}
		*files = append(*files, path)
		return nil
	}
}

// Obtain a list of all files in the specified project tree starting
// at the root.
// Ignore directories starting with a .
// Ignore the assets directory
func getProjectTree(path string) (tree []string, err error) {
	var files []string
	err = filepath.Walk(path, visit(&files))
	if err != nil {
		return []string{}, err
	}
	return files, nil
}


func main() {
	files, err := getProjectTree(".")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%+v", strings.Join(files[:],"\n"))
	}
}

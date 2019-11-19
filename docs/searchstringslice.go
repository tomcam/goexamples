// Shows how to sort a string slice in place, then search it.
// My Go Playground for it is here: https://play.golang.org/p/yhmqfGgDvfE
package main

import (
	"fmt"
	"sort"
)

var excludedDirs = []string{".git", "pub", ".backup"}

// sortStringSlice does just that if you pass it the address of a string slice.
func sortStringSlice(s *[]string) {
	sort.Slice(*s, func(i, j int) bool {
		return (*s)[i] <= (*s)[j]
	})
}

// inSlice() returns true if the search term is found in a sorted slice.
func inSlice(search string, s []string) bool {
	pos := sort.Search(len(s), func(i int) bool {
		return string(s[i]) >= search
	})

	return s[pos] == search
}
func main() {
	searchFor := ".git"
	sortStringSlice(&excludedDirs)
	if inSlice(searchFor, excludedDirs) {
		fmt.Printf("Found %s\n", searchFor)
	} else {
		fmt.Printf("Couldn't find %s\n", searchFor)
	}
	searchFor = "node_modules"
	if inSlice(searchFor, excludedDirs) {
		fmt.Printf("Found %s\n", searchFor)
	} else {
		fmt.Printf("Couldn't find %s\n", searchFor)
	}

}

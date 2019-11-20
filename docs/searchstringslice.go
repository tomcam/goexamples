/* Sorts a slice in place the first time it's searched by creating
a compound data structure that holds the slice and a "sorted" flag. 
Useful when the list is created at runtime. Note that the sorted flag is 
false by default.
https://play.golang.org/p/a571mfn90E-
*/
package main

import (
	"fmt"
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

var exclude = searchInfo{
	list: []string{"node_modules", ".git", "test", "public", "backup"}, sorted: false}

func main() {
	searchFor := "backup"
	fmt.Printf("%s in list? %v\n", searchFor, exclude.Found(searchFor))
	searchFor = "keepme"
	fmt.Printf("%s in list? %v\n", searchFor, exclude.Found(searchFor))

}

// Reusable code to obtain the last modified date of a file
package main

import (
	"fmt"
	"os"
	"time"
)

// lastModified() returns the last modified date of the specified
// file.
func lastModified(filename string) (t time.Time, err error) {
	file, err := os.Stat(filename)
	var theTime time.Time
	if err != nil {
		return theTime, err
	}
	return file.ModTime(), nil
}

// Pass the name of the file of interest on the command line
// and its last modified date is displayed.
func main() {
	if t, err := lastModified(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stdout, "%v\n", err.Error())
		os.Exit(1)
	} else {
		fmt.Fprintf(os.Stdout, "%s was last modified at %v\n", os.Args[1], t)
	}
}

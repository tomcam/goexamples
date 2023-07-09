// Create a small tree of files for a minimal website.
// Run program then open the file WWW/index.html
package main

import (
	"os"
	"path/filepath"
)

var (
	// Location for tree of test files
	baseDir = "WWW"

	// Contents for a tree of test files, generated right here.
	// Array of elements that each consist of
	// [0] a filename
	// [1] the contents of that file.
	testFiles = [][]string{

		// Very simple HTML file
		{"index.html", "<!doctype html><title>.</title><head><link rel='stylesheet' href='css/root.css'><link rel='stylesheet' href='assets/img/background.css'></head><h1>hello, world.</h1>"},

		// Tiny stylesheet
		{"css/root.css", "html {max-width:70ch;font-family:sans-serif;padding:3em 1em;margin:auto;line-height:1.75;font-size:1.25em;}"},

		// Repeated background image of a light grey star
		{"assets/img/background.css", "body { background-image: url('data:image/svg+xml;charset=utf-8,%3Csvg%20xmlns%3D%22http%3A%2F%2Fwww.w3.org%2F2000%2Fsvg%22%20width%3D%2240%22%20height%3D%2240%22%20viewBox%3D%220%200%2024%2024%22%3E%3Cpath%20fill%3D%22%23000000%22%20d%3D%22M12.86%2010.44L11%206.06l-1.86%204.39l-4.75.41L8%2014l-1.08%204.63L11%2016.17l4.09%202.46L14%2014l3.61-3.14l-4.75-.42m3.73%2010.26L11%2017.34L5.42%2020.7l1.46-6.35l-4.92-4.28l6.49-.57l2.55-6l2.55%206l6.49.57l-4.92%204.27l1.47%206.36Z%22%20style%3D%22fill%3A%20rgb(214%2C%20214%2C%20214)%3B%22%3E%3C%2Fpath%3E%3C%2Fsvg%3E');"},
	}
)

// Generate a small tree of files for the
// file preview test. That tree starts
// in location baseDir, which can be anywhere
func createTestDir(baseDir string, testFiles [][]string) {
	// Array contains a filename followed by its contents,
	// like this:
	//{"index.html", "<h1>hello, world.</h1>"},
	for _, row := range testFiles {
		// First element of the row contains a file path
		// designation. Strip the directory.
		// Append that directory to the base directory.
		dir := filepath.Join(baseDir, filepath.Dir(row[0]))
		// Create the specified directory.
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			panic(err)
		}
		// Extract the filename from the first element of the row.
		filename := filepath.Base(row[0])
		// Append it to the directory path to create a
		// fully qualified filename.
		filePath := filepath.Join(dir, filename)
		// Obtain the contents of the file
		// from the second element of the row.
		// Write the file contents out.
		stringToTextFile(filePath, row[1])
	}
}

// stringToTextFile creates a file called filename without checking to see if it
// exists, then writes contents to it.
// filePath is a fully qualified pathname.
// contents is the string to write. Appends a newline to that string.
func stringToTextFile(filePath, contents string) {
	var out *os.File
	var err error
	defer out.Close()
	if out, err = os.Create(filePath); err != nil {
		panic(err)
	}
	if _, err = out.WriteString(contents + "\n"); err != nil {
		panic(err)
	}
}

func main() {
	createTestDir(baseDir, testFiles)
}

package main // MicroCMS: One-file Markdown file to HTML CMS
// git clone https://github.com/tomcam/microcms
// cd microcms
// go mod init github.com/tomcam/microcms
// go mod tidy

// Example invocations
// Include the 2 css files shown
// go run main.go -styles "theme.css light-mode.css" foo.md

// Get CSS file from CDN
// go run main.go -styles "https://unpkg.com/spectre.css/dist/spectre.min.css" foo.md > foo.html
// go run main.go -styles "//writ.cmcenroe.me/1.0.4/writ.min.css" foo.md > foo.html

// Notes:
// - www is a subdir of project
import (
	"bytes"
	"flag"
	"fmt"
	"github.com/yuin/goldmark"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

var defaultExample = `
# CMS example
hello, world.
`

var docType = `
<!DOCTYPE html>
<html lang=`

func assemble(article string, title string, language string, styles []string) string {
	var htmlFile string
	var stylesheets string
	for _, sheet := range styles {
		s := fmt.Sprintf("\t<link rel=\"stylesheet\" href=\"%s\"/>\n", sheet)
		stylesheets += s
	}
	htmlFile = docType + "\"" + language + "\">" + "\n" +
		"<head>\n" +
		"\t<meta charset=\"utf-8\">\n" +
		"\t<meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\">\n" +
		"\t<title>" + title + "</title>\n" +
		stylesheets +
		"</head>\n<body>" +
		article +
		"</body>\n</html>"
	fmt.Println(htmlFile)
	return htmlFile
}


func main() {
	var styles string
	flag.StringVar(&styles, "styles", "", "One or more stylesheets (use quotes if more than one)")

	var templs string
	flag.StringVar(&templs, "templates", "", "One or more templates (use quotes if more than one)")

	var title string
	flag.StringVar(&title, "title", "powered by microCMS", "Contents of the HTML title tag")

	var language string
	flag.StringVar(&language, "language", "en", "HTML language designation, such as en or fr")

	flag.Parse()
	filename := flag.Arg(0)

	stylesheets := strings.Split(styles, " ")
	//templates := strings.Split(templs, ", ")
	//fmt.Printf("Filename: %v\nStylesheets: %v\nTemplates: %v\nTitle: %v", filename, stylesheets, templates, title)

	var exclude searchInfo
	exclude.list = []string{"node_modules", "main.bak", ".git", "pub", ".DS_Store", ".gitignore"}

	var markdownExtensions searchInfo
	markdownExtensions.list = []string{".md", ".mkd", ".mdwn", ".mdown", ".mdtxt", ".mdtext", ".markdown"}

	mdDirectoryTreeToHTML(".", "WWW", exclude, markdownExtensions)
	quit("Complete. Check remaining code.", nil, 0)

	if HTML, err := mdFileToHTML(filename); err != nil {
		quit("Error creating Markdown file", err, 1)
	} else {
		assemble(HTML, title, language, stylesheets)
		//fmt.Println(HTML)
		quit("Complete", nil, 0)
	}

}

// mdToHTML takes Markdown source as a byte slice and converts it to HTML
// using Goldmark's default settings.
func mdToHTML(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	if err := goldmark.Convert(input, &buf); err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

// mdFileToHTML converts a source file to an HTML string
// using Goldmark's default settings.
func mdFileToHTML(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	if HTML, err := mdToHTML(bytes); err != nil {
		return "", err
	} else {
		return string(HTML), nil
	}
}

func quit(msg string, err error, exitCode int) {
	if err != nil {
		fmt.Printf("%s: %v\n", msg, err.Error())
	} else {
		fmt.Printf("%s\n", msg)
	}
	os.Exit(exitCode)
}

// mdDirectoryTreeToHTML takes startDir as the root directory,
// converts all files (except those in exclude.List) to HTML,
// and deposits them in www. Attempts to create www if it
// doesn't exist. www is expected to be a subdirectory of
// startDir.
func mdDirectoryTreeToHTML(startDir string, www string, exclude searchInfo, markdownExtensions searchInfo) {

	var err error

	// Change to requested directory
	if err = os.Chdir(startDir); err != nil {
		quit(fmt.Sprintf("Unable to change to directory %s", startDir), err, 1)
	}

  // Cache project's root directory
	var currDir string
	if currDir, err = os.Getwd(); err != nil {
		quit("Unable to get name of current directory", err, 1)
	}

  // Collect all the files required for this project.
  // exclude.List contains a list of files not to process.
	files, err := getProjectTree(".", exclude)
	if err != nil {
		quit("Unable to get directory tree", err, 1)
	}

	// Now have list of all files in directory tree.
	// If markdown, convert to HTML and copy that file to the HTML publication directory.
	// If not, copy to target publication directory unchanged.

	// Full pathname of file to copy to target directory
	var source string

  // Full pathname of output directory for copied files
  var target string

  // After Markdown file is converted to HTML, it ends up in this string.
  // and eventually 
	var HTML string

  // Relative directory of file. Required to determine where
  // to copy target file.
	var rel string


  // true if it was converted to HTML.
  // false if it's not a Markdown file, which means it will be copied 
  // unchanged to the output directory
  var converted bool


  // Main loop. Traverse the list of files to be copied. 
  // If a file is Markdown as determined by its file extension,
  // convert to HTML and copy to output directory.
  // If a file isn't Markdown, copy to output directory with
  // no processing.
	for _, filename := range files {

    // true if it's  Markdown file converted to HTML
    converted = false

    // Get the fully qualified pathname for this file.
		filename = filepath.Join(currDir, filename)

    // Separate out the file's origin directory
		sourceDir := filepath.Dir(filename)
		//fmt.Printf("%s\n", filename)

    // Get the relatve directory. For example, if your directory
    // is ~/raj/blog and you're in ~/raj/blog/2023/may, then
    // the relative directory is 2023/may.
		if rel, err = filepath.Rel(currDir, sourceDir); err != nil {
			quit(fmt.Sprintf("Unable to get relative paths of %s and %s\n", filename, www), err, 1)
		}

    // Determine the destination directory. If the base publish
    // directory is named WWW, then in the previous example
    // it would be ~/raj/blog/WWW, or ~/raj/blog/WWW/2023/may
		targetDir := filepath.Join(currDir, www, rel)
		// Obtain file extension.
		ext := path.Ext(filename)
		// Replace converted filename extension, from markdown to HTML.
		// Only convert to HTML if it has a Markdown extension.
		if markdownExtensions.Found(ext) {
			// Convert the Markdown file to an HTML string
			if HTML, err = mdFileToHTML(filename); err != nil {
				quit("Error converting Markdown file to HTML", err, 1)
			}
			// Strip origal file's Markdown extension and make 
      // the destination files' extension HTML
			source = filename[0:len(filename)-len(ext)] + ".html"
			target = filepath.Join(targetDir, filepath.Base(source))
      // Write the string to the new filename and location
			//writeStringToFile(target, HTML)
      converted = true
		} else {
			// Not a Markdown file. Copy unchanged.
			// Insert destination (WWW) directory
			//target = filepath.Join(filename, www)

			source = filename
			target = filepath.Join(targetDir, filepath.Base(source))
			fmt.Printf("NOT Markdown: Copy %s to %s\n", filename, target)
      converted = false
		}
		if sourceDir != currDir && !dirExists(targetDir) {
			//fmt.Printf("mkdir %s\n", targetDir)
			err := os.MkdirAll(targetDir, os.ModePerm)
			if err != nil && !os.IsExist(err) {
				quit(fmt.Sprintf("Unable to create directory %s", targetDir), err, 1)
			}
		}
		// xxx
    if converted { 
      writeStringToFile(target, HTML)
    } else {
		  //fmt.Printf("copy %s to %s\n", source, target)
		  copyFile(source, target)
    }
	}
}

// FILE UTILITIES
// Clear but
func copyFile(source string, target string) {
	if source == target {
		quit(fmt.Sprintf("copyFile: %s and %s are the same", source, target), nil, 1)
	}
	if source == "" {
		quit("copyFile: no source file specified", nil, 1)
	}
	if target == "" {
		quit(fmt.Sprintf("copyFile: no destination file specified for file %s", source), nil, 1)
	}
	if src, err := os.Open(source); err != nil {
		quit(fmt.Sprintf("copyFile: Unable to open file %s", source), err, 1)
	} else {
		if trgt, err := os.Create(target); err != nil {
			quit(fmt.Sprintf("copyFile: Unable to create file %s", target), err, 1)
		} else {
			if _, err := io.Copy(src, trgt); err != nil {
				quit(fmt.Sprintf("Error copying file %s to %s", source, target), err, 1)
			}
		}
	}
}

// dirExists() returns true if the name passed to it is a directory.
func dirExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	} else {
		return false
	}
}

// writeStringToFile creates a file called filename without checking to see if it
// exists, then writes contents to it.
func writeStringToFile(filename, contents string) {
	fmt.Printf("writeStringToFile filename: %s\n", filename)
	var out *os.File
	var err error
	if out, err = os.Create(filename); err != nil {
		quit(fmt.Sprintf("writeStringToFile: Unable to create file %s", filename), err, 1)
	}
	if _, err = out.WriteString(contents); err != nil {
		// TODO: Renumber error code?
		quit(fmt.Sprintf("Error writing to file %s", filename), err, 1)
	}
}

// SLICE UTILITIES
// Searching a sorted slice is fast.
// This tracks whether the slice has been sorted
// and sorts it on first search.

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

// DIRECTORY TREE

func visit(files *[]string, exclude searchInfo) filepath.WalkFunc {
	// var exclude searchInfo
	// Find out what directories to exclude
	//exclude.list = []string{"node_modules", "main.bak", ".git", "pub", ".DS_Store", ".gitignore"}
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			// Quietly fail if unable to access path.
			return err
		}
		isDir := info.IsDir()

		// Obtain just the filename.
		name := filepath.Base(info.Name())

		// Skip any directory to be excluded, such as
		// the pub and .git directores
		if exclude.Found(name) && isDir {
			return filepath.SkipDir
		}
		// It may be just a filename on the exclude list.
		if exclude.Found(name) {
			return nil
		}

		// Don't add directories to this list.
		if !isDir {
			*files = append(*files, path)
		}
		return nil
	}
}

// Obtain a list of all files in the specified project tree starting
// at the root.
// Ignore directories starting with a .
// Ignore the assets directory
func getProjectTree(path string, exclude searchInfo) (tree []string, err error) {
	var files []string
	err = filepath.Walk(path, visit(&files, exclude))
	if err != nil {
		return []string{}, err
	}
	return files, nil
}

package main

import (
	"flag"
	"fmt"
	"os"
)
var (
	// Initialize command-line flags.

	// After the command line has been parsed, if "build" ocurreed
	// on the command line then buildCmd.Parse() will be true, and
	// of course false if there was no init command specified. 
	// To see if the (optional) output dir has been specified,
	// look for the value of that string, which is obtained
	// using *buildOutputDir. It will be empty if not specified.
	buildCmd = flag.NewFlagSet("build", flag.ExitOnError)
	buildOutputDir = buildCmd.String("output-dir", "public", "Directory for generated HTML")
	buildBaseURL = buildCmd.String("base-url", "", "Override URL specified in config")

	// Init can take either of these forms:
	// init
	// init sitename=something
	initCmd = flag.NewFlagSet("init", flag.ExitOnError)
	initSiteName = initCmd.String("sitename", "", "Your site name")

)

// Get commands, including subcommands with optional flags, and filenames off command line.
// I do something gross here.
// Options can take these forms, where foo is the name of the executable:
// foo init filename1 filename2
// foo init -sitename="My site"
// foo init -sitename="My site"
// foo build -output-dir=public init -sitename=fred foo bar build -base-url=htttps:://mypp.test.com test.md
// foo filename -init
//
// Where:
//   init has the optional arument -sitename=value
//   init can be used by itself
//   build has multiple arguments that can be placed anywhere on the command line
//   filenames are anything leftover after the the rest 
//   has been processed, and can appear anywhere.
//
// The gross thing is that the default case falls through when an argument
// is being processed. Which means that while it works for filenames, in the
// case of the command line "foo init -sitename=test news.md" it correctly
// accepts news.cmd as a filename but also, incorrectly, -sitename=test. So  
// I had to add the parsingArgument flag. When an argument is consumed
// using Parse(os.Args[pos+1:]) then parsingArgument must be set manually to true.
//
func parseCmdLine() []string{
	// Anything that wasn't command, argument, or option will be returned
	// as a filename.
	// Upon return:
	//    if a command like init was encountered,
	var filenames []string
	var parsingArgument bool
	for  pos, nextArg := range os.Args {
		//fmt.Printf("\n\nTop of loop. os.Args[%v]: %v. parsingArgument: %v\n", pos, nextArg,parsingArgument)
		switch nextArg {
		case "init":
			// initCmd.Parsed() is now true
			initCmd.Parse(os.Args[pos+1:])
			if *initSiteName != "" {
				// *initSitename now points to the string value.
				// It's empty if -sitename  wasn't -sitename specified
				parsingArgument = true
			}
		case "build":
			buildCmd.Parse(os.Args[pos+1:])
			if *buildOutputDir != "" {
				parsingArgument = true
			}
			if *buildBaseURL != "" {
				parsingArgument = true
			}
		default:
			// If not in the middle of parsing a command-like subargument,
			// like the -sitename=test in this command line:
			//   foo init -sitename=test
			// Where foo is the name of the program, and -sitename is
			// an optional subcommand to init,
			//
			// os.Args[0] falls through so exclude it, since it's
			// the name of the invoking program.
			if !parsingArgument && pos > 0{
				filenames =  append(filenames, nextArg)
			} else {
				parsingArgument = false
			}

		}
	}
	return filenames
}

func main() {
	flag.Parse()
	files := parseCmdLine()
	if initCmd.Parsed() {
		fmt.Println("init")
		if *initSiteName != "" {
			fmt.Println("--sitename: " + *initSiteName)
		}
	}
	if buildCmd.Parsed() {
		fmt.Println("build")
		if *buildOutputDir != "" {
			fmt.Println("--output-dir: " + *buildOutputDir)
		}
		if *buildBaseURL != "" {
			fmt.Println("--base-url: " + *buildBaseURL)
		}
	}
	if len(files) > 0 {
		fmt.Println(files)
	} else {
		fmt.Println("No files specified")
	}
}

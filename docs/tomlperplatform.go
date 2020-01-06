/* Shows how to determine the OS at runtime and store OS-specific values
   so that they can be retrieved at runtime and, for example, be used
   to write platform-specific info in config files that can nevertheless
   be read and written to without having to set the OS manually.
   A config file is created if none exists. Any existing one is read in.
   If it's run on a previously unused operating system, adds a section
   in the config for that OS>
   Sniffs OS, then writes specific configuration file for that OS like this:

[darwin]
  [darwin.PlatformSpecific]
    home = "/Users/tom"

[windows]
  [windows.PlatformSpecific]
    home = "C:\Users\tom"

  Uses the go-homedir package as an example of platform-specific config.

  https://play.golang.org/p/FXrGXgrXABN
*/
package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"runtime"
)

// homeDir() returns the user's home directory, or just "." for
// the current directory if it can't be determined through system
// calls.
func homeDir() string {
	h, err := homedir.Dir()
	if err != nil {
		return "."
	}
	u, err := homedir.Expand(h)
	if err != nil {
		return "."
	}
	return u
}

// fileExists() returns true, well, if the named file exists
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// writeMapFile() creates a TOML file based on the filename and
// map passed in. This works unchanged with any TOML-compatible
// data structure.
func writeMapFile(filename string, target interface{}) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if err = toml.NewEncoder(f).Encode(target); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

// readMapFile() opens the TOML file, reads the contents
// of any TOML-compatible data structure, and marshals
// them into the data structure whose address has been
// passed to target.
func readMapFile(filename string, target interface{}) (err error) {
	var input []byte
	if input, err = ioutil.ReadFile(filename); err != nil {
		return err
	}
	if _, err = toml.Decode(string(input), target); err != nil {
		return err
	}
	return nil
}

func main() {
	// Find out what OS we're running
	OS := runtime.GOOS

	// Config defines the subsection of the config file that looks something like this,
	// where the "darwin" part is supplied at runtime
	/*
	  [darwin.PlatformSpecific]
	    home = "/Users/tom"
	*/
	type Config struct {
		PlatformSpecific map[string]string
	}
	const appCfgFilename = "app.cfg"

	// App defines the top-level section of the config file that looks
	// something like this, where the "darwin" part is determined by runtime.GOOS
	// at runtime:
	/*
	  [darwin]
	*/
	var App map[string]Config
	App = make(map[string]Config)

	// Is there already a config file?
	if fileExists(appCfgFilename) {
		fmt.Println(appCfgFilename, "exists. Now reading it in.")
		// Yes. Read it in. May be for a different OS.
		if err := readMapFile(appCfgFilename, &App); err != nil {
			panic(err.Error())
		}
	} else {
		fmt.Println(appCfgFilename, "doesn't exist yet. Now creating it.")
	}
	// See what existing config there is for this OS, if any.
	_, ok := App[OS]
	if !ok {
		// Not found, so create an entry for it.
		var c Config
		c.PlatformSpecific = make(map[string]string)
		c.PlatformSpecific["home"] = homeDir()
		App[OS] = c
	}

	// Create a TOML file with the current configuration.
	// If there was a previous configuration for a different
	// OS, it's preserved.
	if err := writeMapFile(appCfgFilename, &App); err != nil {
		panic(err.Error())
	}

	// Just to be sure, read back TOML file and display its contents.
	if err := readMapFile(appCfgFilename, &App); err != nil {
		panic(err.Error())
	}

	fmt.Printf("Home dir on %s: %s\n",
		OS, App[OS].PlatformSpecific["home"])
}

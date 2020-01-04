/* Sniffs OS, then writes specific configuration file for that OS like this:

[darwin]
  [darwin.PlatformSpecific]
    home = "/Users/tom"

[windows]
  [windows.PlatformSpecific]
    home = "C:"

  Uses the go-homedir package as an example of platform-specific config.
  
  https://play.golang.org/p/Au91bMKPPFh
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
	// TODO:Test on windows. Or without using ~
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

/// curDir() returns the current directory name. Doesn't deal with errors because
// it's for diagnostic purposes.
func currDir() string {
	if path, err := os.Getwd(); err != nil {
		return "."
	} else {
		return path
	}
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
// map passed in.
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

func writeTomlFile(filename string, target interface{}) error {
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
	OS := runtime.GOOS
	type Config struct {
		PlatformSpecific map[string]string
	}
	const appCfgFilename = "app.cfg"
	var App map[string]Config
	App = make(map[string]Config)
	// Is there a config file?
	if fileExists(appCfgFilename) {
		// Yes. Read it in. May be for a different OS.
		if err := readMapFile(appCfgFilename, &App); err != nil {
			panic(err.Error())
		}
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

	// Create a TOML file with all versions.
	if err := writeMapFile(appCfgFilename, &App); err != nil {
		panic(err.Error())
	}

	// Read back TOML file and display its contents.
	if err := readMapFile(appCfgFilename, &App); err != nil {
		panic(err.Error())
	}
	fmt.Printf("Contents of %s:\n%+v\n", appCfgFilename, App)
}

/* multiconfig.go creates a single config file with sections for each OS, like this:

[darwin]
  [darwin.Features]
    home = "/Users/tom/code/t"

[windows]
  [windows.Features]
    home = "C:\Users\Userdata\t\code"

*/
package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"runtime"
)

/// curDir() returns the current directory name. Doesn't deal with errors because
// it's for diagnostic purposes.
func currDir() string {
	if path, err := os.Getwd(); err != nil {
		return "unknown directory"
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

func readMapFile(filename string) (c map[string]Config, err error) {
	var input []byte
	var cfg map[string]Config
	if input, err = ioutil.ReadFile(filename); err != nil {
		return nil, err
	}
	if _, err = toml.Decode(string(input), &cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

const configFilename = "foo.cfg"

type Config struct {
	Features map[string]string
}

func main() {
	var cfg map[string]Config
	var c Config
	// Get operating system identifier as a string
	OS := runtime.GOOS
	// Can test creating new entries by uncommenting this
	// and assigning it different values
	//OS = "foo"

	// No config file at all, so create one.
	if !fileExists(configFilename) {
		fmt.Println("Creating", configFilename)
		// Create the global config.
		cfg = make(map[string]Config)
		// This is the "subfile", which doesn't know what OS it's for.
		c.Features = make(map[string]string)
		// Note current directory
		c.Features["home"] = currDir()
		// Add an entry in the file for the current OS.
		cfg[OS] = c
		// Create the file using current config info.
		if err := writeMapFile(configFilename, &cfg); err != nil {
			panic(err.Error())
		}
		fmt.Println("Created", configFilename)
		os.Exit(0)
	}
	// Config file exists. Don't know if it has an
	// entry for this OS.
	var err error
	// Read the existing file.
	if cfg, err = readMapFile(configFilename); err != nil {
		panic(err.Error())
	}
	_, ok := cfg[OS]
	// Is there an entry for this OS?
	if !ok {
		// No. Create one.
		var f Config
		// Have to allocate the submap.
		f.Features = make(map[string]string)
		f.Features["home"] = currDir()
		// Then allocate the new one by assigning the old one.
		cfg[OS] = f
		// And save as a file.
		if err := writeMapFile(configFilename, &cfg); err != nil {
			panic(err.Error())
		}
		os.Exit(0)
	} else {
		fmt.Println("Already have an entry for", OS)
	}
}

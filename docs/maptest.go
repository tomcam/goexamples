package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	//"time"
)

/// curDir() returns the name of the current working directory
func currDir() string {
	if path, err := os.Getwd(); err != nil {
		return ""
	} else {
		return path
	}
}


// suggestCfgDir() takes a stab at return the recommended location
// of the directory that stores user configuration data. It does
// so by sniffing the operating system. It returns the suggested
// directory in the dir variable and the OS it detected in the os variable.
func suggestCfgDir() (dir, system string) {
	system = runtime.GOOS
	// Try to determine where user application home dir would be
	switch system {
	case "windows":
		return os.Getenv("%APPDATA%"), system
	case "darwin":
		u, _ := user.Current()
		return filepath.Join(u.Username, u.HomeDir, "Library", "Preferences"), system
	case "linux":
		return os.Getenv("HOME"), system
	default:
		return os.Getenv("HOME"), system
	}
}

type Config struct {
	Features map[string]string
}

func main() {
	configDir, opsys := suggestCfgDir()
	fmt.Printf("Suggested config directory for %s: %s\n", opsys, configDir)
	var cfg map[string]Config
	blob := `
	[darwin]
		[darwin.Features]
			home = "/Users/tom"
			newline = "/n"
			editor = "textedit"
	[windows]
		[windows.Features]
			home = "c:"
			newline = "/r/n"
			editor = "notepad"
`

	filename := "foo.txt"
	if _, err := toml.Decode(blob, &cfg); err != nil {
		panic(err.Error())
	}
	cfg[opsys].Features["filename"]=currDir()

	if err := writeMapFile(filename, &cfg); err != nil {
		panic(err.Error())
	}
	if err := readMapFile(filename, &cfg); err != nil {
		panic(err.Error())
	}

	type C map[string]interface{}
	c := make(C)
	//c["foo"] = &interface{"ned":"fred"}
	//["foo"] = {"fred":"bed"}
	c["foo"] = "1"
	c["bar"] =  [1,2,3]

	fmt.Println(c)
/*
type Config struct {
	Features map[string]string
}

func main() {
	var cfg map[string]Config

*/



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

func readMapFile(filename string, m interface{}) (err error) {
	var input []byte
	if input, err = ioutil.ReadFile(filename); err != nil {
		return err
	}

	if _, err = toml.Decode(string(input), &m); err != nil {
		return err
	}

	return nil
}



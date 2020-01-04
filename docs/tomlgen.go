/* Shows how to create general purpose, freeform config file in TOML format  like this:
https://play.golang.org/p/LgZwOT363sZ

[darwin]
  home = "/users/tom"
  version = "0.5.1"
*/
package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

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
	const cfgFilename = "app.cfg"
	var Config map[string]map[string]string
	Config = make(map[string]map[string]string)
	kv := make(map[string]string)
	kv["home"] = "/users/tom"
	Config["darwin"] = kv
	kv["version"] = "0.5.1"
	Config["darwin"] = kv

	if err := writeMapFile(cfgFilename, &Config); err != nil {
		panic(err.Error())
	}
	if err := readMapFile(cfgFilename, &Config); err != nil {
		panic(err.Error())
	}
	fmt.Printf("Contents of %s:\n%+v\n", cfgFilename, Config)
}

// Write contents of struct to YAML file. Read YAML file back into a struct. Uses gopkg.in/yaml 
package main

import (
	"fmt"
	"io/ioutil"
	"gopkg.in/yaml.v3"
  "os"
)

type Theme struct {
	Name        string `yaml:"name"`
	Branding    string `yaml:"branding"`
	Description string `yaml:"description"`
}

func main() {
	filename := "./theme.yaml"
	theme := Theme{Name: "Debut",
			Branding:    "Debut by Metabuzz",
			Description: "Perfect theme to showcase a new product"}

	b, err := yaml.Marshal(&theme)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filename, b, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Created file %v\n", filename)

/* The generated file looks like this:

name: Debut
branding: Debut by Metabuzz
description: Perfect theme to showcase a new product

*/

	b, err = ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read file %v\n", filename)

  var t Theme
	err = yaml.Unmarshal(b, &t)
	if err != nil {
		panic(err)
	}

  fmt.Printf("File contents: %#v\n", t)

}

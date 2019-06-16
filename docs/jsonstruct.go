// Shows how to initialize nested struct in golang, write to JSON file,
// and read back from JSON file. Uses the real-life example
// of an OAuth2 Config struct.
package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"io/ioutil"
	"os"
)

func main() {
	// Initialize an OAuth2 Config struct, which contains
	// a nested struct of type Endpoint.
	c := oauth2.Config{
		ClientID:     "SampleClientID-SBX-ebf9-78dade18",
		ClientSecret: "SBX-ebf915a9-59ed-47e6-98f7-4201",
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://signin.sandbox.ebay.com/authorize",
			TokenURL:  "https://api.sandbox.ebay.com/identity/v1/oauth2/token",
			AuthStyle: oauth2.AuthStyleInParams},
		Scopes: []string{"https://api.ebay.com/oauth/api_scope"},
	}

	// Display its contents.
	fmt.Fprintf(os.Stdout, "Config object before marshaling:\n%+v\n\n", c)

	// Convert it to a JSON-encoded blob (byte array).
	b, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	// Write it to a file.
	err = ioutil.WriteFile("sample.json", b, 0644)
	if err != nil {
		panic(err)
	}

	// Read the file back into another byte array.
	b2, err := ioutil.ReadFile("sample.json")
	if err != nil {
		panic(err)
	}

	// Convert the JSON-encoded blob into an oauth2.Config struct
	var c2 oauth2.Config
	err = json.Unmarshal(b2, &c2)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(os.Stdout, "Config object after marshaling:\n%+v\n\n", c)
	fmt.Fprintf(os.Stdout, "Copy of config object read from file:\n%+v\n\n", c2)

}

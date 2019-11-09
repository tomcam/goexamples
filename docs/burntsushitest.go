package main
import (
	"fmt"
	"io/ioutil"
    "github.com/BurntSushi/toml"
	"os"
)
/* Contents of file site.toml:
   # Comments are ignored
   name = "foo.html
 */
type SiteConfig struct {
	Name string
}


func readSiteConfig(infile string) (err error) {
  var input[]byte
  if input, err = ioutil.ReadFile(infile); err != nil {
    return err
  }

  var s SiteConfig
  if _, err = toml.Decode(string(input),&s); err != nil {
    return err
  }

  fmt.Fprintf(os.Stdout, "Site config: %+v\n", s)
  return nil
}

func main() {
  if err := readSiteConfig("site.toml"); err != nil {
	  fmt.Fprintf(os.Stdout, "Error reading site config: %v\n", err.Error()) 
  }
}

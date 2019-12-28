package main
/* Golang command-line flag tutorial example for boolean flag
   Go Playground example at:
   https://play.golang.org/p/fEaMQCjxQUT
   Assuming the program is named ./t here's example usage and output:

	$ ./t 
	Verbose mode set? false
	$ ./t -v
	Verbose mode set? true
	$ ./t -v=t
	Verbose mode set? true
	$ ./t -v=f
	Verbose mode set? false
	$ ./t -v=true
	Verbose mode set? true
	$ ./t -v=false
	Verbose mode set? false
	$ ./t -v=fa   
	invalid boolean value "fa" for -v: parse error
	Usage of ./t:
	  -v	Explain what's happening while program runs
*/
import (
	"flag"
	"fmt"
	"os"
)

func main() {
	pVerbose := flag.Bool("v", false, "Explain what's happening while program runs")
	flag.Parse()
	fmt.Fprintf(os.Stdout, "Verbose mode set? %v\n", *pVerbose)
}

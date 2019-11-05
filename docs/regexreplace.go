package main


import (
    "fmt"
    "regexp"
)


const fakeMd = `

# Testing

What if we put template variables like {{.Name}} in markdown?
`


func main() {
    re := regexp.MustCompile(`\{\{\.[_a-zA-Z][_a-zA-Z0-9]*\}\}`)
    fmt.Println(re.ReplaceAllString(fakeMd, "your cat Fluffy Fusty Feather Fungus"))
}

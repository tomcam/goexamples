// structFieldByNameStrMust() takes any struct and field name (as a string) at runtime and
// returns the string value of that field. It returns an empty string if the
// object passed in isn't a struct, or if the named field isn't a struct.
package main

import (
	"fmt"
	"reflect"
)

type Site struct {
	Name string
	Path string
}

type Document struct {
	Published bool
	Title     string
}

func main() {
	s := Site{
		Name: "example",
		Path: "/Users/tom/metabuzz",
	}
	d := Document{
		Published: true,
		Title:     "yo mama",
	}
	fmt.Print(structFieldByNameStrMust(d, "Title"))
	fmt.Print(structFieldByNameStrMust(s, "Title"))
	fmt.Print(structFieldByNameStrMust(d, "title"))
}

// structFieldByNameStrMust() takes any struct and field name (as a string) 
// passed in at runtime and returns the string value of that field. 
// It returns an empty string if the
// object passed in isn't a struct, or if the named field isn't a struct.
func structFieldByNameStrMust(obj interface{}, field string) string {
	v := reflect.ValueOf(obj)
	if v.Kind() != reflect.Struct {
		return ""
	}
	kind := v.FieldByName(field).Kind()
	if kind != reflect.String {
		return ""
	}
	return (fmt.Sprint(v.FieldByName(field)))

}

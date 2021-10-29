// structInfo() takes any struct at runtime and displays its type name, field names and types, 
// and contents of each field. 
// structHasField() returns true if a struct passed to it at runtime contains a field name passed as a string 
// Playground version at https://play.golang.org/p/zeOTNfHEQlH
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
	d := Document{
		Published: true,
		Title:     "yo mama",
	}
	s := Site{
		Name: "foo",
		Path: "/Users/tom/go",
	}
	structInfo(d)
	structInfo(s)
	structInfo(9)	
	
	fmt.Println(structHasField(d, "dude"))
	fmt.Println(structHasField(d, "title"))
	fmt.Println(structHasField(d, "Title"))
}
// structHasField() returns true if a struct passed to it at runtime contains a field name passed as a string 
func structHasField(obj interface{}, field string) bool {
	v := reflect.ValueOf(obj)
	if reflect.TypeOf(obj).Kind() != reflect.Struct {
		return false
	}
	return reflect.Indirect(v).FieldByName(field).IsValid()
}

// structInfo() takes any struct at runtime and displays its type name, field names and types, 
// and contents of each field. 
func structInfo(obj interface{}) {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if reflect.TypeOf(obj).Kind() != reflect.Struct {
		fmt.Printf("\n%v is not a struct\n", obj)
		return
	}

	fmt.Printf("\nstruct type: %#v\n", t.Name())
	for i := 0; i < t.NumField(); i++ {
		fmt.Printf("\t%v %v\t%v\n", v.Type().Field(i).Name, v.Field(i).Type(), v.Field(i).Interface())

	}
}

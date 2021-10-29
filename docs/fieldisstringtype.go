

// fieldIsStringType() determines whether the struct passed in the
// argument has a field named in the second argument of type string.
// Playground version at https://play.golang.org/p/yAEXeeCvJMH

package main

import (
	"fmt"
	"reflect"
)

type Document struct {
	Published bool
	Title     string
}

func main() {
	d := Document{
		Published: true,
		Title:     "yo mama",
	}
	fmt.Println(fieldIsStringType(d, "Published"))
	fmt.Println(fieldIsStringType(d, "Title"))
	fmt.Println(fieldIsStringType(d, "title"))
}

// fieldIsStringType() determines whether the struct passed in the
// argument has a field named in the second argument of type string.
func fieldIsStringType(obj interface{}, key string) bool {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	if reflect.TypeOf(obj).Kind() != reflect.Struct {
		return false
	}
	for i := 0; i < t.NumField(); i++ {
		fieldType := fmt.Sprint(v.Field(i).Type())
		if fieldType == "string" && v.Type().Field(i).Name == key {
			return true
		}
	}
	return false
}


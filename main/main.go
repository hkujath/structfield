package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := MyStruct{"abc", "def"}
	stype := reflect.TypeOf(s)
	field := stype.Field(0)
	fmt.Println(field.Tag.Get("key1"), field.Tag.Get("key2"))
}

type MyStruct struct {
	Field1 string `key1:"value1" key2:"-"`
	Field2 string `key1:"value2"`
}

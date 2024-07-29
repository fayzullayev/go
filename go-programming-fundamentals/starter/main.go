package main

import (
	"fmt"
	"reflect"
	"strings"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Addr string `json:"addr,omitempty"`
}

func main() {

	t := reflect.TypeOf(Person{})

	field, _ := t.FieldByName("Addr")

	fmt.Println(strings.Split(field.Tag.Get("json"), ",")[1])

}

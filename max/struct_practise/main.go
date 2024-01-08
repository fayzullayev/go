package main

import (
	"fmt"
	"reflect"
)

func main() {

	fmt.Println(add(1, 2))
	fmt.Println(reflect.TypeOf(add(4.5, 5.5)))

}

func add[T int | float64](a, b T) T {
	return a + b
}

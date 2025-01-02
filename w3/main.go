//https://www.w3schools.com/go/go_arrays.php

package main

import (
	"fmt"
)

func main() {
	var txt1 string = "Hello"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %s\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %s\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %s\n", txt3, txt3[len(txt3)-1:])

}

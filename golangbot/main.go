package main

import "fmt"

// Define a new type based on `int`.
type myInt int

// Define a method on `myInt`.
func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	// Use the new type with the method.
	var num1 myInt = 10
	var num2 myInt = 20

	result := num1.add(num2)
	fmt.Println(result)
}

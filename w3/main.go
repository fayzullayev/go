package main

import "fmt"

func main() {
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var b = []string{"one", "two", "three", "four"} // defining the order

	for k, v := range a { // loop with no order
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()

	for _, element := range b { // loop with the defined order
		fmt.Printf("%v : %v, ", element, a[element])
	}

}

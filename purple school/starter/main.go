package main

import "fmt"

func main() {
	var name string
	fmt.Println("name before :", name)

	fmt.Scanf("Hel %v llo", &name)
	fmt.Println("My name is :", name)
}

package main

import "fmt"

func main() {
	var a [3]int

	fmt.Println("a", a)

	a[2] = 67
	a[0] = 23
	a[1] = 1

	fmt.Println("a", a)
	fmt.Println("len(a)", len(a))
	fmt.Println("cap(a)", cap(a))

}

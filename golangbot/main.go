package main

import "fmt"

func main() {
	a := [...]int{3: 12, 1: 33, 99: 1000}
	fmt.Println(a)
	fmt.Println(cap(a))
	fmt.Println(len(a))
}

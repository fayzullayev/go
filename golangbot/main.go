package main

import "fmt"

func main() {
	var a = [3]string{"A", "B", "C"}
	var b = [3]string{"A", "B", "C"}

	fmt.Println(a == b)
}

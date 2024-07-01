package main

import "fmt"

func main() {
	var num int

	if num < 5 {
		fmt.Println(num)
		num++
		goto LOOP
	}

LOOP:
	fmt.Println("Hello")

}

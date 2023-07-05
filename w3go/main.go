package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos := adder()

	fmt.Println(pos(1))
	fmt.Println(pos(1))
}

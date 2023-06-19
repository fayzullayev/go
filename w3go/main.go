package main

import (
	"fmt"
)

func factorialRecursion(x float64) (y float64) {
	if x > 0 {
		y = x * factorialRecursion(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	fmt.Println(factorialRecursion(40))

}

//https://www.w3schools.com/go/go_struct.php

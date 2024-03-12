package main

import "fmt"

func main() {
	for i := range 100 {
		isFizzBuzz(i)
	}
}

func isFizzBuzz(number int) {
	isDivisableBy3 := number%3 == 0
	isDivisableBy5 := number%5 == 0

	if isDivisableBy3 && isDivisableBy5 {
		fmt.Println(number, "FizzBuzz")
	} else if isDivisableBy3 {
		fmt.Println(number, "Fizz")
	} else if isDivisableBy5 {
		fmt.Println(number, "Buzz")
	} else {
		fmt.Println(number)
	}
}

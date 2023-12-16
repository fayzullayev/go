package main

import "fmt"

func main() {
	age := 32

	var agePointer *int

	//fmt.Println(*agePointer) ---> ERROR

	agePointer = &age

	fmt.Println(age)
	fmt.Println(agePointer)

	getAdultYears(agePointer)
	fmt.Println("-------------------")

	fmt.Println(age)
	fmt.Println(agePointer)
}

func getAdultYears(age *int) {
	*age -= 18
}

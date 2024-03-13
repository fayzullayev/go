package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy = iota
	Business
	FirstClass
)

func main() {

	switch p := price(); {
	case p < 2:
		fmt.Println("Cheap item")
	case p < 10:
		fmt.Println("Moderately priced item")
	default:
		fmt.Println("Expensive  item")
	}

	ticket := Economy

	switch ticket {
	case Economy:
		fmt.Println("Economy seating")
	case Business:
		fmt.Println("Business seating")
	case FirstClass:
		fmt.Println("FirstClass seating")

	default:
		fmt.Println("Other seating")
	}
}

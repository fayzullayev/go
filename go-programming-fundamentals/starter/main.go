package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota << 3
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Print()
}

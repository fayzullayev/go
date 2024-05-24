package main

import "fmt"

func main() {

	aSlice := make([]int, 4, 4)
	fmt.Println(aSlice)

	// The capacity is doubled
	fmt.Println("L:", len(aSlice), "C.:", cap(aSlice)) // Now add four elements
	aSlice = append(aSlice, []int{-1, -2, -3, -4, 6}...)

	fmt.Println(aSlice)
	// The capacity is doubled
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))
}

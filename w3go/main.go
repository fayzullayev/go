package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	arrNum := numbers[7:]

	fmt.Println(arrNum)
	fmt.Println("len()", len(arrNum))
	fmt.Println("cap()", cap(arrNum))

	arrNum2 := make([]int, len(arrNum))

	copy(arrNum2, arrNum)
	fmt.Println("-----------")
	fmt.Println(arrNum2)
	fmt.Println("len()", len(arrNum2))
	fmt.Println("cap()", cap(arrNum2))

}

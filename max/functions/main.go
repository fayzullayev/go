package main

import "fmt"

type numeric interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(add(66, nums...))

	fmt.Println(add(1, 2, 3, 4))
}

func add[T numeric](num int, nums ...T) T {

	fmt.Println("num", num)
	var sum T

	for _, v := range nums {
		sum += v
	}

	return sum
}

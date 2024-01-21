package main

import (
	"fmt"
)

type numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(add[float64](5.6, 6.6))
}

func add[T numbers](a, b T) T {
	return a * b
}

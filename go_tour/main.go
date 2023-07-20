package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {

	var v Vertex

	v = Vertex{
		X: 40,
		Y: 30,
	}

	fmt.Println(Abs(v))

}

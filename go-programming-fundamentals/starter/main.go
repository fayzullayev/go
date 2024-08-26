package main

import (
	"fmt"
	"math"
	"reflect"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Height, Width float64
}

func (r *Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

func main() {
	c := &Circle{10}
	r := &Rectangle{10, 20}

	fmt.Println(TotalArea(c, r))

}

func TotalArea(shapes ...Shape) float64 {
	var total float64

	fmt.Println(reflect.TypeOf(shapes))

	for _, shape := range shapes {
		total += shape.Area()
	}

	return total

}

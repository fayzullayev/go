package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) getArea() float64 {
	return (t.height * t.base) / 2
}

type square struct {
	sideLength float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func main() {
	t := triangle{
		base:   10,
		height: 10,
	}

	s := square{sideLength: 10}

	printArea(t)
	printArea(s)
}

func printArea(sh shape) {
	fmt.Println(sh.getArea())
}

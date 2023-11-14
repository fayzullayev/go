package main

import "time"

type Philosopher struct {
	name      string
	rightFork int
	leftFork  int
}

var philosopher = []Philosopher{
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}

const hunger = 3
const eatTime = 1 * time.Second
const thinkTime = 3 * time.Second
const sleepTime = 1 * time.Second

func main() {

}

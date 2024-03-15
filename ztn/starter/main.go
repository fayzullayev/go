package main

import "fmt"

type Counter struct {
	hint int
}

func increment(counter *Counter) {
	counter.hint += 1

	fmt.Printf("Counter %+v\n", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {
	counter := Counter{}
	fmt.Printf("Counter %+v\n", counter)
	hello := "Hello"
	world := "World!"

	fmt.Println(hello, world)

	replace(&hello, "Hi", &counter)

	fmt.Println(hello, world)
}

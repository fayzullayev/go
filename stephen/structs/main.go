package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func change(p Person) {
	p.LastName = "Hello"
}

func main() {

	p := Person{
		FirstName: "Mirodil",
		LastName:  "Fayzullyev",
	}

	fmt.Println(p)

	change(p)

	fmt.Println(p)

}

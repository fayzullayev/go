package main

import (
	"fmt"
)

type Greeter interface {
	Greet() string
}

type EnglishBot struct {
	Text string
}

func (eb EnglishBot) Greet() string {
	return eb.Text
}

type SpanishBot struct {
	Text string
}

func (eb SpanishBot) Greet() string {
	return eb.Text
}

type Name string

func (n Name) Greet() string {
	return string(n)
}

func main() {

	eb := EnglishBot{
		"Hello",
	}

	sb := SpanishBot{
		Text: "Hola!",
	}

	name := Name("Mirodil")

	Greeting(eb)
	Greeting(sb)
	Greeting(name)
}

func Greeting(gr Greeter) {
	fmt.Println(gr.Greet())
}

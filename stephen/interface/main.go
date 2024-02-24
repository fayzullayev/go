package main

import "fmt"

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

func main() {
	eb := EnglishBot{
		"Hello",
	}

	sb := SpanishBot{
		Text: "Hola!",
	}

	Greeting(eb)
	Greeting(sb)
}

func Greeting(gr Greeter) {
	fmt.Println(gr.Greet())
}

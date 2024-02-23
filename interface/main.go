package main

import "fmt"

// EnglishBot --------------englishBot--------------//
type EnglishBot struct {
	Text string
}

func (eb EnglishBot) Greet() string {
	return eb.Text
}

// SpanishBot --------------englishBot--------------//
// --------------spanishBot--------------//
type SpanishBot struct {
	Text string
}

func (sb SpanishBot) Greet() string {
	return sb.Text
}

//--------------spanishBot--------------//

type Greeter interface {
	Greet() string
}

func main() {

	eb := EnglishBot{
		"Hello",
	}

	sb := SpanishBot{
		Text: "Hola",
	}

	PrintGreeting(eb)
	PrintGreeting(sb)

}

func PrintGreeting(gr Greeter) {
	fmt.Println(gr.Greet())
}

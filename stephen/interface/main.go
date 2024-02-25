package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

	//eb := EnglishBot{
	//	"Hello",
	//}
	//
	//sb := SpanishBot{
	//	Text: "Hola!",
	//}
	//
	//name := Name("Mirodil")
	//
	//Greeting(eb)
	//Greeting(sb)
	//Greeting(name)

	res, err := http.Get("https://www.google.com")

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, res.Body)
}

func Greeting(gr Greeter) {
	fmt.Println(gr.Greet())
}

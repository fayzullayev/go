package main

import (
	"fmt"
	"time"
)

func greet(phrase string, d chan bool) {
	fmt.Println("Hello!", phrase)
	d <- true

}

func slowGreet(phrase string, d chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	d <- true
	close(d)
}

func main() {
	done := make(chan bool)

	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!", done)

	for range done {

	}
	//fmt.Println(<-done)
	//fmt.Println(<-done)
	//fmt.Println(<-done)
	//fmt.Println(<-done)
}

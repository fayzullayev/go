package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://amazon.com",
		"https://golang.org",
	}

	var ch = make(chan string)

	for _, link := range links {
		go checkLink(link, ch)
	}

	for {
		fmt.Println(<-ch)
	}

}

func checkLink(link string, ch chan string) {
	_, err := http.Get(link)

	if err != nil {
		ch <- link + " might be down!"
		return
	}

	ch <- link + " is up"
}

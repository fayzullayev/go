// package main
//
// import (
//
//	"fmt"
//	"net/http"
//	"time"
//
// )
//
//	func main() {
//		links := []string{
//			"https://google.com",
//			"https://facebook.com",
//			"https://stackoverflow.com",
//			"https://amazon.com",
//			"https://golang.org",
//		}
//
//		var ch = make(chan string)
//
//		for _, link := range links {
//			go checkLink(link, ch)
//		}
//
//		for l := range ch {
//			go func(link string) {
//				time.Sleep(time.Second * 3)
//				checkLink(link, ch)
//			}(l)
//		}
//
// }
//
// func checkLink(link string, ch chan string) {
//
//		_, err := http.Get(link)
//
//		if err != nil {
//			fmt.Println(link, " might be down")
//			ch <- link
//			return
//		}
//
//		fmt.Println(link, " is up")
//		ch <- link
//	}
package main

import "fmt"

func main() {
	c := make(chan string)

	for i := 0; i < 4; i++ {
		go printString("Hello there!", c)
	}

	for {
		fmt.Println(<-c)
	}
}

func printString(s string, c chan string) {
	fmt.Println(s)
	c <- "Done printing."
}

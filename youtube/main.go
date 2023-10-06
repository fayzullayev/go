//package main
//
//import (
//	"sync"
//)
//
//type httpPkg struct{}
//
//func (httpPkg) Get(url string) {}
//
//var http httpPkg
//
//func main() {
//	var wg sync.WaitGroup
//
//	var urls = []string{
//		"http://www.golang.org/",
//		"http://www.google.com/",
//		"http://www.example.com/",
//	}
//
//	for _, url := range urls {
//		// Increment the WaitGroup counter.
//		wg.Add(1)
//		// Launch a goroutine to fetch the URL.
//		go func(url string) {
//			// Decrement the counter when the goroutine completes.
//			defer wg.Done()
//			// Fetch the URL.
//			http.Get(url)
//		}(url)
//	}
//	// Wait for all HTTP fetches to complete.
//	wg.Wait()
//}

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Start")
	n := 500

	for i := 0; i < n; i++ {
		wg.Add(1)
		go myFunc(i)
	}

	wg.Wait()
	fmt.Println("End")
}

func myFunc(num int) {
	time.Sleep(3 * time.Second)
	fmt.Println("num", num)
	wg.Done()
}

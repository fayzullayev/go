package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	words := make(map[string]int)

	for _, v := range strings.Fields(s) {
		if _, ok := words[v]; ok {
			words[v]++
		} else {
			words[v] = 1
		}
	}
	fmt.Println("words", words)
	return words
}

func main() {
	wc.Test(WordCount)
}

package main

import (
	"fmt"
	"strconv"
)

func main() {
	i, err := strconv.Atoi("43333sdgvsdvds2")

	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}

	fmt.Println("Converted integer:", i)
}

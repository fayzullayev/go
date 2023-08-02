package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)

	fmt.Println(b)

	fmt.Println("-------------------")

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		fmt.Println("-------------------")
		if err == io.EOF {
			break
		}
	}
	fmt.Printf("b[:n] = %q\n", b)
}

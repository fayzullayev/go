package main

import (
	_ "embed"
	"fmt"
)

//go:embed test.txt
var myFile string

func main() {
	fmt.Println("content - ", myFile)
}

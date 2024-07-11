package main

import (
	"fmt"
	"os"
)

func main() {
	os.Kill
	handlePanic()

	fmt.Println("Execution resumes here, after recovery")
}

func handlePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	causePanic()

	fmt.Println("This statement will not be executed")
}

func causePanic() {
	defer fmt.Println("Deferred calls run even if a function panics.")
	panic("A runtime error has occurred.")
}

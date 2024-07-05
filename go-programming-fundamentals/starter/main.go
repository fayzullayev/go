package main

import "fmt"

func main() {
	causePanic()
}

func causePanic() {
	defer func() {
		fmt.Println("Deferred calls run even if a function panics.")
	}()

	panic("A runtime error has occurred")
}

func handlePanic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

}

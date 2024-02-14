package main

import "fmt"

type deck []string

func (d *deck) print() {
	for index, card := range *d {
		fmt.Println(index, card)
	}
}

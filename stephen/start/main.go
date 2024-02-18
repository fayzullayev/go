package main

import "fmt"

func main() {

	//numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for number := range 11 {
		if number%2 == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}

	//
	//cards := ReadFromFile("my_cards.txt")
	//
	////fmt.Println(cards.toString())
	//cards.Print()
	//cards.Shuffle()
	////fmt.Println()
	//cards.Print()

	//err := cards.SaveToFile("my_cards.txt")
	//if err != nil {
	//	log.Fatal("Something went wrong", err.Error())
	//}

}

package main

func main() {

	cards := ReadFromFile("my_cards.txt")

	//fmt.Println(cards.toString())
	cards.Print()
	cards.Shuffle()
	//fmt.Println()
	cards.Print()

	//err := cards.SaveToFile("my_cards.txt")
	//if err != nil {
	//	log.Fatal("Something went wrong", err.Error())
	//}

}

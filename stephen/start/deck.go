package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Deck []string

func newDeck() Deck {
	cards := Deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := fmt.Sprintf("%s of %s", value, suit)
			cards = append(cards, card)
		}
	}

	return cards
}

func (d *Deck) Print() {
	for index, card := range *d {
		fmt.Println(index, card)
	}
}

func deal(d Deck, handSize int) (Deck, Deck) {
	return d[:handSize], d[handSize:]

	//os.WriteFile()
	//os.ReadFile()

	//os.Open()
	//os.OpenFile()
}

func (d *Deck) SaveToFile(filename string) error {

	return os.WriteFile(filename, []byte(d.ToString()), 0666)
}

func ReadFromFile(filename string) Deck {

	bytes, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal("Error:", err)
	}

	s := strings.Split(string(bytes), ",")

	return s
}

func (d *Deck) ToString() string {
	fmt.Println(*d)
	return strings.Join(*d, ",")
}

//func (d *Deck) Shuffle() {
//	rand.Shuffle(len(*d), func(i, j int) {
//		fmt.Println(i, j)
//		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
//	})
//}

func (d *Deck) Shuffle() {

	t := time.Now().UnixNano()
	s := time.Now().Unix()

	newSource := rand.NewSource(t)
	fmt.Println("time.Now().UnixNano()", t)
	fmt.Println("time.Now()", s)

	r := rand.New(newSource)

	for index := range *d {
		newPosition := r.Intn(len(*d) - 1)

		(*d)[index], (*d)[newPosition] = (*d)[newPosition], (*d)[index]
	}
}

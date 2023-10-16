package main

import (
	"fmt"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

const numberOfPizzas = 10

var pizzaMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch

	return <-ch
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func main() {
	color.Cyan("The Pizzeria is open business!")
	color.Cyan("------------------------------")

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	go pizzeria(pizzaJob)
}

func makePizza(pizzaNumber int) *PizzaOrder {
	pizzaNumber++

	if pizzaNumber <= numberOfPizzas {
		delay := rand.Intn(5) + 1
		fmt.Printf("Recevied order #%d!\n", pizzaNumber)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzaMade++
		}

		total++

		fmt.Printf("Making pizza #%d. It will take %d seconds...\n", pizzasFailed, delay)
		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("*** We ran out of ingredients for pizza #%d!", pizzaNumber)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("*** The cook quit while making pizza #%d!", pizzaNumber)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready", pizzaNumber)
		}

		return &PizzaOrder{
			pizzaNumber: pizzaNumber,
			message:     msg,
			success:     success,
		}

	}

	return &PizzaOrder{
		pizzaNumber: pizzaNumber,
	}

}

func pizzeria(pizzaMaker *Producer) {

	var i = 0

	for {

		currentPizza := makePizza(i)

	}
}

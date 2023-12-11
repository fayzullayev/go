package main

import (
	"fmt"
	"log"
)

func main() {
	var accountBalance = 1000.0

	fmt.Println("-----------------------")
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")

BankOperations:
	for {
		fmt.Println("-----------------------")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		fmt.Println("-----------------------")

		var choice int
		fmt.Print("Your choice: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			log.Fatal(err)
		}

		switch choice {
		case 1:
			fmt.Println("Your balance:", accountBalance)
		case 2:
		case 3:
		case 4:
			break BankOperations
		default:

		}
	}

	fmt.Println("Thank you for your operation")
}

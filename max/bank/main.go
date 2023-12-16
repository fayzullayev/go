package main

import (
	"fmt"
	"github.com/bank/fileops"
	"log"
)

const FileName = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(FileName)

	if err != nil {
		fmt.Println("Error", err)
	}

	fmt.Println("-----------------------")
	fmt.Println("Welcome to Go Bank!")

mainLoop:
	for {
		presentOptions()

		var choice int
		fmt.Print("Your choice: ")
		_, err = fmt.Scan(&choice)
		if err != nil {
			log.Fatal(1, err)
		}

		switch choice {
		case 1:
			fmt.Println("-----------------------")
			fmt.Println("Your balance is", accountBalance)
		case 2:
			fmt.Print("Your deposit: ")
			var depositAmount float64
			_, err = fmt.Scan(&depositAmount)
			if err != nil {
				log.Fatal(2, err)
			}
			if depositAmount <= 0 {
				fmt.Println("-----------------------")
				fmt.Println("Invalid amount. Must be greater than 0.")
				continue
			}
			accountBalance += depositAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, FileName)
		case 3:
			fmt.Print("Withdrawal amount: ")
			var withdrawalAmount float64
			_, err = fmt.Scan(&withdrawalAmount)
			if err != nil {
				log.Fatal(3, err)
			}

			if withdrawalAmount <= 0 || withdrawalAmount > accountBalance {
				fmt.Println("-----------------------")
				fmt.Println("Invalid amount.")
				continue
			}
			accountBalance -= withdrawalAmount
			fmt.Println("Balance updated! New amount: ", accountBalance)
			fileops.WriteFloatToFile(accountBalance, FileName)
		default:
			break mainLoop
		}
	}

	fmt.Println("Thank you for your operation")
}

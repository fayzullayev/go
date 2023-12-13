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

mainLoop:
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
			log.Fatal(1, err)
		}

		if choice == 1 {
			fmt.Println("-----------------------")
			fmt.Println("Your balance is", accountBalance)
		} else if choice == 2 {
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
		} else if choice == 3 {
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
		} else {
			break mainLoop
		}

		//switch choice {
		//case 1:
		//	fmt.Println("-----------------------")
		//	fmt.Println("Your balance:", accountBalance)
		//case 2:
		//case 3:
		//case 4:
		//	break mainLoop
		//default:
		//	fmt.Println("-----------------------")
		//	fmt.Println("Wrong choice")
		//}
	}

	fmt.Println("Thank you for your operation")
}

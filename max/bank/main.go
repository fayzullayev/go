package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const FileName = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile("./" + FileName)
	if err != nil {
		return 1000.0, errors.New("failed to find balance file")
	}

	balanceText := string(data)

	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000.0, errors.New("failed to parse stored balance value")
	}

	return balance, nil

}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	err := os.WriteFile(FileName, []byte(balanceText), 0644)
	if err != nil {
		log.Fatal(4, err)
	}
}

func main() {
	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("Error", err)
		panic("Sorry bitch")
	}

	fmt.Println("-----------------------")
	fmt.Println("Welcome to Go Bank!")

mainLoop:
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("-----------------------")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
		fmt.Println("-----------------------")

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
		default:
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

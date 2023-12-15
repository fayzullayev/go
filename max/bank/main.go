package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const FileName = "balance.txt"

func getFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000.0, errors.New("failed to find file")
	}

	valueText := string(data)

	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000.0, errors.New("failed to parse stored value")
	}

	return value, nil

}

func writeBalanceToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	err := os.WriteFile(fileName, []byte(valueText), 0644)
	if err != nil {
		log.Fatal(4, err)
	}
}

func main() {
	var accountBalance, err = getFloatFromFile(FileName)

	if err != nil {
		fmt.Println("Error", err)
		panic("Sorry bitch")
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
			writeBalanceToFile(accountBalance, FileName)
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
			writeBalanceToFile(accountBalance, FileName)
		default:
			break mainLoop
		}
	}

	fmt.Println("Thank you for your operation")
}

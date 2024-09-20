package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var transactions []float64

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the transactions parser")

	for {
		fmt.Print("Enter amount of transactions: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			panic("Something went wrong!")
		}

		stringNumber := strings.ReplaceAll(strings.TrimSpace(line), "\n", "")

		floatNumber, err := strconv.ParseFloat(stringNumber, 64)
		if err != nil {
			panic("Error parsing!")
		}

		transactions = append(transactions, floatNumber)

		fmt.Print("Do you want to continue? (y/n): ")
		line, err = reader.ReadString('\n')
		if err != nil {
			panic("Something went wrong!")
		}

		line = strings.ReplaceAll(strings.TrimSpace(line), "\n", "")

		if !(line == "Y" || line == "y") {
			break
		}
	}

	fmt.Println("Transactions: ", transactions)
	fmt.Println("Total transactions: ", len(transactions))

}

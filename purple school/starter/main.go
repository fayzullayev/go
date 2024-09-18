package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a number: ")
		readString, err := reader.ReadString('\n')
		if err != nil {
			//fmt.Println("Error reading input-1:", err)
			panic("vsdvsdvsdv")
			break
		}

		errors.New("wefwefwefwefwef")

		num, err := strconv.Atoi(strings.ReplaceAll(readString, "\n", ""))
		if err != nil {
			fmt.Println("Error parsing input-2:", err)
			return
		}
		fmt.Printf("%d - %T \n --------------------- \n", num, num)

		fmt.Print("Do you want to continue: ")
		prompt, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		if strings.ToLower(strings.TrimSpace(strings.ReplaceAll(prompt, "\n", ""))) != "y" {
			break
		}

	}

	fmt.Println("Done")
}

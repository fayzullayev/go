package main

import (
	"fmt"
	"math"
)

const inflationRate = 6.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Investment Amount: ")
	_, err := fmt.Scan(&investmentAmount)
	if err != nil {
		fmt.Println("investmentAmount scan error", err)
	}

	fmt.Print("Expected Return Rate: ")
	_, err = fmt.Scan(&expectedReturnRate)
	if err != nil {
		fmt.Println("expectedReturnRate scan error", err)
	}

	fmt.Print("Years: ")
	_, err = fmt.Scan(&years)
	if err != nil {
		fmt.Println("years scan error", err)
	}

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}

package main

import (
	"fmt"
	"log"
	"math"
)

const inflationRate = 6.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	getUserInputs(&investmentAmount, &expectedReturnRate, &years)

	//var futureValue, futureRealValue float64 = calculateFutureValues(investmentAmount, expectedReturnRate, years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future value: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future real value: %.1f\n", futureRealValue)

	fmt.Print(formattedFV, formattedFRV)
	//
	//	fmt.Printf(`Future value: %.1f
	//Future real value: %.1f
	//`, futureValue, futureRealValue)

}

func getUserInputs(investmentAmount, expectedReturnRate, years *float64) {

	fmt.Print("Investment Amount: ")
	_, err := fmt.Scan(investmentAmount)
	if err != nil {
		log.Fatal("investmentAmount scan error", err)
	}

	fmt.Print("Expected Return Rate: ")
	_, err = fmt.Scan(expectedReturnRate)
	if err != nil {
		log.Fatal("expectedReturnRate scan error", err)
	}

	fmt.Print("Years: ")
	_, err = fmt.Scan(years)
	if err != nil {
		log.Fatal("years scan error", err)
	}
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (futureValue float64, futureRealValue float64) {
	futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue = futureValue / math.Pow(1+inflationRate/100, years)

	return
}

//max - 34,35,36,37,38 - done

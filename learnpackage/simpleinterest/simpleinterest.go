package simpleinterest

import "fmt"

func init() {
	fmt.Println("Simple interest package initialized")
}

// Calculate calculates and returns the simple interest for a principal p, rate of interest r for time duration t years
func Calculate(p, r, t float64) float64 {
	interest := p * (r / 100) * t
	return interest
}

package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	//fmt.Println(revenue, expenses, taxRate)
	//fmt.Println(reflect.TypeOf(revenue), reflect.TypeOf(expenses), reflect.TypeOf(taxRate))

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses

	profit := ebt * (1 - taxRate/100)

	ratio := ebt / profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

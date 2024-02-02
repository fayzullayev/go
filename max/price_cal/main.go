package main

import (
	"fmt"
	"price_cal/filemanager"
	"price_cal/prices"
)

func main() {

	var taxRates = []float64{0, 0.5, 0.6, 0.7}

	for _, taxRate := range taxRates {
		//cmd := cmdmanager.New()
		fm := filemanager.New("prices2.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Println("Something went wrong:", err.Error())
		}
	}
}

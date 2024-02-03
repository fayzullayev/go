package main

import (
	"fmt"
	"price_cal/filemanager"
	"price_cal/prices"
	"time"
)

func main() {
	start := time.Now()

	var taxRates = []float64{0, 0.5, 0.6, 0.7}

	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	fmt.Println("chanels", doneChans)

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)
		//cmd := cmdmanager.New()
		fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)
		go priceJob.Process(doneChans[i], errorChans[i])

		//if err != nil {
		//	fmt.Println("Something went wrong:", err.Error())
		//}
	}

	for i := range taxRates {
		select {
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Println(err)
			}
		case <-doneChans[i]:
			fmt.Println("Done")
		}
	}
	//
	//select {
	//case doneChans:
	//
	//}
	//
	//fmt.Println("chanels", doneChans)
	//
	//for _, doneChan := range doneChans {
	//	<-doneChan
	//}

	fmt.Println(time.Since(start))
}

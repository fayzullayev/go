package prices

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) Process() {

	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		tax := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", tax)
	}

	fmt.Printf("%+v", result)

}

func (job *TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("./prices.txt")
	if err != nil {
		log.Fatal("no such file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("file error 2", err)
		err = file.Close()
	}

	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {

		floatPrice, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Converting price to float failed")
			fmt.Println(err)
			file.Close()
			return
		}
		prices[lineIndex] = floatPrice
	}

	job.InputPrices = prices

}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

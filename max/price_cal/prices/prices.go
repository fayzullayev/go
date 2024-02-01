package prices

import (
	"fmt"
	"log"
	"price_cal/conversion"
	"price_cal/filemanager"
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
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Printf("%+v", result)

}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := filemanager.ReadLine("./prices.txt")

	if err != nil {
		log.Fatal(err)
	}

	//}

	prices, err := conversion.StringToFloats(lines)
	//for lineIndex, line := range lines {
	//
	//	floatPrice, err := strconv.ParseFloat(line, 64)
	//
	if err != nil {
		log.Fatal(err)
	}
	//	prices[lineIndex] = floatPrice
	//}

	job.InputPrices = prices
	file.Close()
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

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
	TaxIncludedPrices map[string]string
}

func (job *TaxIncludedPriceJob) Process() {

	job.LoadData()

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	err := filemanager.WriteJSON(fmt.Sprintf("result_%.0f.json", job.TaxRate*100), job)

	if err != nil {
		return
	}

	//fmt.Printf("%+v", result)

}

func (job *TaxIncludedPriceJob) LoadData() {

	lines, err := filemanager.ReadLine("./prices.txt")

	if err != nil {
		log.Fatal(err)
	}

	prices, err := conversion.StringToFloats(lines)

	if err != nil {
		log.Fatal(err)
	}

	job.InputPrices = prices
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {

	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

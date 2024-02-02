package prices

import (
	"fmt"
	"price_cal/conversion"
	"price_cal/iomanager"
)

type TaxIncludedPriceJob struct {
	iomanager.IOManager `json:"-"`
	TaxRate             float64           `json:"tax_rate"`
	InputPrices         []float64         `json:"input_prices"`
	TaxIncludedPrices   map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) Process() error {

	err := job.LoadData()

	if err != nil {
		return err
	}

	result := make(map[string]string)

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	return job.WriteResult(job)
}

func (job *TaxIncludedPriceJob) LoadData() error {

	lines, err := job.ReadLine()

	if err != nil {
		return err
	}

	prices, err := conversion.StringToFloats(lines)

	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil
}

func NewTaxIncludedPriceJob(fm iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate:   taxRate,
		IOManager: fm,
	}
}

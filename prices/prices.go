// Package prices provides functionality related to pricing calculations.
package prices

import (
	"fmt"

	"github.com/hisshihi/golang-lessons/conversion"
	"github.com/hisshihi/golang-lessons/iomanager"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
	IOManager         iomanager.IOManager `json:"-"`
}

func (job *TaxIncludedPriceJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return fmt.Errorf("ошибка при загрузке данных: %v", err)
	}

	if len(lines) == 0 {
		return fmt.Errorf("файл пустой или не содержит данных")
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	job.InputPrices = prices

	return nil
}

func (job *TaxIncludedPriceJob) Process() {
	job.LoadData()

	result := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	job.TaxIncludedPrices = result

	job.IOManager.WriteResult(job)
}

func NewTaxIncludedPriceJob(taxRate float64, iom iomanager.IOManager) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
		IOManager:   iom,
	}
}

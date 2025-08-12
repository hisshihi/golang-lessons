// Package prices provides functionality related to pricing calculations.
package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TaxIncludedPriceJob struct {
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func (job *TaxIncludedPriceJob) LoadData() error {
	// Загружаем файл и получаем данные о ценах
	file, err := os.Open("prices.txt")
	if err != nil {
		return fmt.Errorf("возникла ошибка с чтением файла: %v", err.Error())
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return fmt.Errorf("возникла ошибка с чтением данных из файла: %v", err.Error())
	}

	prices := make([]float64, len(lines))
	for lineIndex, line := range lines {
		price, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return fmt.Errorf("не удалось преобразовать строку в float: %v", err.Error())
		}

		prices[lineIndex] = price
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

	fmt.Println(result)
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		InputPrices:       []float64{10, 20, 30},
		TaxRate:           taxRate,
		TaxIncludedPrices: map[string]float64{},
	}
}

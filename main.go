package main

import (
	"fmt"

	"github.com/hisshihi/golang-lessons/filemanager"
	"github.com/hisshihi/golang-lessons/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	
	for _, taxRate := range taxRates {
		fm := filemanager.NewFileManager("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fm)
		priceJob.Process()
	}
	fmt.Println("Done.")
}

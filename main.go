package main

import (
	"fmt"

	"github.com/hisshihi/golang-lessons/cmdmanager"
	"github.com/hisshihi/golang-lessons/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		// fm := filemanager.NewFileManager("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.NewCMDManager()
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, cmdm)
		priceJob.Process()
	}
	fmt.Println("Done.")
}

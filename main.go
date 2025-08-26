package main

import (
	"fmt"
	"time"

	"github.com/hisshihi/golang-lessons/filemanager"
	"github.com/hisshihi/golang-lessons/prices"
)

func main() {
	now := time.Now()
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		fm := filemanager.NewFileManager("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.NewCMDManager()
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fm)
		go priceJob.Process(doneChans[i])
	}

	for _, doneChan := range doneChans {
		<- doneChan
	}

	fmt.Println("Done.")
	fmt.Println(time.Since(now))
}

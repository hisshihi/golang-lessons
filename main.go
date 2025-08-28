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
	errorChans := make([]chan error, len(taxRates))

	for i, taxRate := range taxRates {
		doneChans[i] = make(chan bool)
		errorChans[i] = make(chan error)
		fm := filemanager.NewFileManager("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		// cmdm := cmdmanager.NewCMDManager()
		priceJob := prices.NewTaxIncludedPriceJob(taxRate, fm)
		go priceJob.Process(doneChans[i], errorChans[i])
	}

	for i := range taxRates {
		select {
		case err := <-errorChans[i]:
			if err != nil {
				fmt.Printf("Возникла ошибка: %v\n", err)
			}
		case <-doneChans[i]:
			fmt.Println("Done.")
		}
	}

	fmt.Println(time.Since(now))
}

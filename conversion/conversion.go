// Package conversion пакет для конвертации данных
package conversion

import (
	"errors"
	"strconv"
)

func StringsToFloat(values []string) ([]float64, error) {
	var floats []float64

	for _, stringVal := range values {
		floatVal, err := strconv.ParseFloat(stringVal, 64)
		if err != nil {
			return []float64{}, errors.New("не удалось преобразовать строку в float64")
		}

		floats = append(floats, floatVal)
	}

	return floats, nil
}

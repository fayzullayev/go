package conversion

import (
	"errors"
	"strconv"
)

func StringToFloats(strs []string) ([]float64, error) {
	var floats []float64

	for _, str := range strs {
		floatPrice, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, errors.New("failed to convert string to float")
		}
		floats = append(floats, floatPrice)
	}
	return floats, nil
}

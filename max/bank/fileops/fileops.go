package fileops

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	valueText := string(data)

	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	err := os.WriteFile(fileName, []byte(valueText), 0644)
	if err != nil {
		log.Fatal(4, err)
	}
}

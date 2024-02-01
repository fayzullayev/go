package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLine(path string) ([]string, error) {

	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("failed to read lines")
	}

	file.Close()

	return lines, nil

}

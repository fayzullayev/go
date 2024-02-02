package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilePath  string `json:"-"`
	OutputFilePath string `json:"-"`
}

func (fm FileManager) ReadLine() ([]string, error) {

	file, err := os.Open(fm.InputFilePath)

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

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to covert data to json")
	}

	file.Close()

	return nil
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

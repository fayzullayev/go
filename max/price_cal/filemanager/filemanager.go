package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type FileManager struct {
	InputFilePath  string `json:"-"`
	OutputFilePath string `json:"-"`
}

func (fm FileManager) ReadLine() ([]string, error) {

	file, err := os.Open(fm.InputFilePath)

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {

		}
	}(file)

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
		return nil, errors.New("failed to read lines")
	}

	return lines, nil

}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			fmt.Println("File closing error")
		}
	}(file)

	if err != nil {
		return errors.New("failed to create file")
	}

	time.Sleep(2 * time.Second)

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)
	if err != nil {
		return errors.New("failed to covert data to json")
	}

	return nil
}

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

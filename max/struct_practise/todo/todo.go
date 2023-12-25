package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(content string) (*Todo, error) {
	if content == "" {
		return &Todo{}, errors.New("invalid input")
	}

	return &Todo{
		Text: content,
	}, nil
}

func (t *Todo) Display() {
	fmt.Println(t.Text)
}

func (t *Todo) Save() error {
	fileName := "todo.json"

	//fmt.Printf("%+v", n)

	jsonData, err := json.Marshal(t)

	//fmt.Println("jsonData", string(jsonData))
	//fmt.Println("jsonData", jsonData)

	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

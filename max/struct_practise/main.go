package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
}

type Saver interface {
	Save() error
}

func saveData(data Saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving the note failed!")
		return err
	}

	fmt.Println("The note was saved successfully")

	return nil
}

func main() {
	//var user []User
	//
	//text := `[
	//			{"name":"Mirodil","age":28,"job":"Developer"},
	//			{"name":"Iroda","age":20,"job":"Teacher"},
	//			{"name":"Zaynab","age":2,"job":"Child"},
	//			{"name":"Muhammad","age":0,"job":"Baby"}
	//		]`

	//data, err := json.Marshal(user)

	//err := json.Unmarshal([]byte(text), &user)

	//if err != nil {
	//	panic(err)
	//}

	//fmt.Println("User Name", user[2].Name)
	//fmt.Println(string(data))

	//var text string

	//text = []byte{11100110 10001111 10010011}
	//
	//fmt.Println("text string", text)
	//fmt.Println("text bytes:", []byte(text))
	//fmt.Printf("text bytes: %08b", []byte(text))
	fmt.Println([]byte("😍"))
	//fmt.Printf("%08b", []byte("Mirodil😍😍😍"))

	//title, content := getNoteData()
	//todoText := getUserInput("Todo text: ")
	//
	//myTodo, err := todo.New(todoText)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//userNote, err := note.New(title, content)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//myTodo.Display()
	//err = saveData(myTodo)
	//if err != nil {
	//	return
	//}
	//
	//userNote.Display()
	//err = saveData(userNote)
	//if err != nil {
	//	return
	//}
}

//func getTodoData() string {
//	text := getUserInput("Todo text: ")
//	return text
//}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt, " ")

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.ReplaceAll(text, "\n", "")
	strings.TrimSuffix(text, "\n")

	return text
}

//
//11100110 10001111 10010011
//
//
//111 1111
//
//
//
//11110000100111111001100010001101 = 😍
//
//11111011000001101 = 😍
//
//
//
//
//
//1 byte = 8 bit = 00000000 yoki 11111111

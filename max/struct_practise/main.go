package main

import (
	"bufio"
	"fmt"
	"os"
	"practise-structs/note"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Job  string `json:"job"`
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
	title, content := getNoteData()

	userNote, err := note.New(title, content)
	if err != nil {
		return
	}

	userNote.Display()

	err = userNote.Save()
	if err != nil {
		panic(err)
	}
}

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

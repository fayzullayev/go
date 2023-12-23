package main

import (
	"bufio"
	"fmt"
	"os"
	"practise-structs/note"
	"strings"
)

func main() {
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
	//var value string

	fmt.Print(prompt, " ")
	//_, err := fmt.Scanln(&value)
	//if err != nil {
	//	panic(err)
	//}

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.ReplaceAll(text, "\n", "")
	strings.TrimSuffix(text, "\n")

	return text
}

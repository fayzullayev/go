package main

import (
	"embed"
	_ "embed"
	"fmt"
	"io/fs"
	"log"
)

//go:embed just_files/* hhhhhh/*
var myFiles embed.FS

func main() {
	files, err := fs.ReadDir(myFiles, "just_files")

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println("Found file:", file.Name())
		content, _ := myFiles.ReadFile("just_files/" + file.Name())
		fmt.Println(string(content))
	}

	fmt.Println("----------------------")

	customFile, err := fs.ReadDir(myFiles, "hhhhhh")

	fmt.Println("customFile", customFile)

	for _, c := range customFile {
		fmt.Println("Found file:", c.Name())
		content, _ := myFiles.ReadFile("hhhhhh/" + c.Name())
		fmt.Println(string(content))
	}

}

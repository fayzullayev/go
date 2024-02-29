package main

import (
	"io"
	"log"
	"os"
)

func main() {
	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Could not find the file: ", filename)
	}

	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Fatal("Could not close the file: ", err.Error())
		}
	}(file)

	io.Copy(os.Stdout, file)

}

package main

import (
	"log"
	"tempaltes-tutorial/app"
)

func main() {
	myApp, err := app.NewApp()
	if err != nil {
		log.Fatal("Failed to create app: ", err)
	}

	err = myApp.Run()
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}

package main

import (
	"github.com/fayzullayev/go-project-example/internal/app"
	"github.com/fayzullayev/go-project-example/internal/database"
	"log"
)

func main() {
	db, err := database.Init()

	if err != nil {
		log.Fatal(err)
	}

	myApp := app.New("5544", db)

	err = myApp.Run()
	if err != nil {
		log.Fatal(err)
	}
}

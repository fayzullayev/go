package main

import "log"

func main() {
	app, err := NewApp()
	if err != nil {
		log.Fatal("Failed to create app:", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal("Error starting server", err)
	}

}

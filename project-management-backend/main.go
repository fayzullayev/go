package main

import (
	"net/http"
)

func main() {

	initDB()

	router := http.NewServeMux()

	router.HandleFunc("GET /projects", getProjects)
	router.HandleFunc("POST /projects", createProject)

	server := http.Server{
		Addr:    ":8088",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic("Something went wrong " + err.Error())
	}
}

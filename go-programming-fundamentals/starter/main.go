package main

import (
	"fmt"
	"net/http"
)

type Employee struct {
	Id   int
	Name string
	Role string
}

func main() {

	handler := http.NewServeMux()

	handler.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {

		fmt.Fprintf(w, "Hell   o %s", " Mirodil")
	})

	http.ListenAndServe(":8888", handler)

}

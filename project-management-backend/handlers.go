package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getProjects(w http.ResponseWriter, r *http.Request) {

	time.Sleep(time.Second * 3)

	enableCors(&w)

	projects, err := getAllProjects()

	if err != nil {
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	json, err := json.Marshal(projects)

	if err != nil {
		http.Error(w, "InternalServerError", http.StatusInternalServerError)
		return
	}

	w.Write(json)

}

func createProject(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()
}

package main

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request){
		tasks := []Task{
			{Id: 1, Name: "Learn Go"},
			{Id: 2, Name: "Build REST API"},
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	})

	http.ListenAndServe(":8080", mux)
}

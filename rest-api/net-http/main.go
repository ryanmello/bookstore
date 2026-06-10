package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Task struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func main(){
	mux := http.NewServeMux()

	tasks := []Task{
		{Id: 0, Name: "Learn Go"},
		{Id: 1, Name: "Build REST API"},
	}

	mux.HandleFunc("GET /tasks", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	})

	mux.HandleFunc("GET /tasks/{id}", func(w http.ResponseWriter, r *http.Request){
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, "Invalid task id", http.StatusBadRequest)
			return
		}
	
		log.Println("GET task with id:", id)

		for _, task := range tasks {
			if task.Id == id {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				json.NewEncoder(w).Encode(task)
				return
			}
		}

		http.Error(w, "Task not found", http.StatusNotFound)
	})

	log.Println("Starting server on port :8080")
	http.ListenAndServe(":8080", mux)
}

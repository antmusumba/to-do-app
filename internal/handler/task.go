package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"to-do/internal/storage"
	"to-do/internal/task"
)

var store = storage.NewFileStorage("tasks.json")

func TaskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGet(w)
	case "POST":
		handlePost(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGet(w http.ResponseWriter) {
	tasks, err := store.List()
	if err != nil {
		http.Error(w, "Could not read tasks", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	var t task.Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	t.CreatedAt = time.Now()
	t.Completed = false

	if strings.TrimSpace(t.Title) == "" {
		http.Error(w, "Title is required", http.StatusBadRequest)
		return
	}

	err = store.Add(t)
	if err != nil {
		http.Error(w, "Could not add task", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Task added")
}

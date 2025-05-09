package main

import (
	"fmt"
	"log"
	"net/http"

	"to-do/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/tasks", handler.TaskHandler)

	fmt.Println("🚀 Server running at http://localhost:8080/tasks")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

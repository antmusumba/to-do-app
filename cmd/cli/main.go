package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"to-do/internal/storage"
	"to-do/internal/task"
)

const dataFile = "tasks.json"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: gotasker [add|list] [args]")
		return
	}

	store := storage.NewFileStorage(dataFile)
	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gotasker add <task title>")
			return
		}
		title := strings.Join(os.Args[2:], " ")
		newTask := task.Task{
			Title:     title,
			CreatedAt: time.Now(),
			Completed: false,
		}
		if err := store.Add(newTask); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("✅ Task added!")
		}
	case "list":
		tasks, err := store.List()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		for _, t := range tasks {
			status := "❌"
			if t.Completed {
				status = "✅"
			}
			fmt.Printf("[%s] %d: %s\n", status, t.ID, t.Title)
		}
	default:
		fmt.Println("Unknown command:", cmd)
	}
}

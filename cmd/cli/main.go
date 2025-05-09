package main

import (
	"fmt"
	"os"
	"strconv"
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
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gotasker done <task ID>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		if err := store.MarkDone(id); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("✅ Task marked as done!")
		}

	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Usage: gotasker remove <task ID>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid task ID:", os.Args[2])
			return
		}
		if err := store.Remove(id); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("🗑️ Task removed!")
		}

	default:
		fmt.Println("Unknown command:", cmd)
	}
}

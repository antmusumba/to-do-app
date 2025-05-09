package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"

	"to-do/internal/task"
)

type FileStorage struct {
	Path  string
	mutex sync.Mutex
}

func NewFileStorage(path string) *FileStorage {
	return &FileStorage{Path: path}
}

func (fs *FileStorage) readTasks() ([]task.Task, error) {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	var tasks []task.Task
	if _, err := os.Stat(fs.Path); errors.Is(err, os.ErrNotExist) {
		return tasks, nil // no file yet
	}

	data, err := os.ReadFile(fs.Path)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func (fs *FileStorage) writeTasks(tasks []task.Task) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(fs.Path, data, 0o644)
}

func (fs *FileStorage) Add(t task.Task) error {
	tasks, err := fs.readTasks()
	if err != nil {
		return err
	}
	t.ID = len(tasks) + 1
	tasks = append(tasks, t)
	return fs.writeTasks(tasks)
}

func (fs *FileStorage) List() ([]task.Task, error) {
	return fs.readTasks()
}

func (fs *FileStorage) MarkDone(id int) error {
	tasks, err := fs.readTasks()
	if err != nil {
		return err
	}

	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = true
			return fs.writeTasks(tasks)
		}
	}

	return fmt.Errorf("task with ID %d not found", id)
}

func (fs *FileStorage) Remove(id int) error {
	tasks, err := fs.readTasks()
	if err != nil {
		return err
	}

	var newTasks []task.Task
	found := false
	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		} else {
			found = true
		}
	}

	if !found {
		return fmt.Errorf("task with ID %d not found", id)
	}

	// Reassign IDs to keep them consecutive
	for i := range newTasks {
		newTasks[i].ID = i + 1
	}

	return fs.writeTasks(newTasks)
}

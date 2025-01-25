package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

const taskFile = "tasks.json"

func loadTasks() []Task {
	file, err := os.ReadFile(taskFile)
	if err != nil {
		return []Task{}
	}
	var tasks []Task
	json.Unmarshal(file, &tasks)
	return tasks
}

func saveTasks(tasks []Task) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(taskFile, data, 0644)
}

func AddTask(description string) {
	tasks := loadTasks()
	tasks = append(tasks, Task{Description: description, Done: false})
	saveTasks(tasks)
}

func ListTasks() []Task {
	return loadTasks()
}

func CompleteTask(index int) {
	tasks := loadTasks()
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid task index")
		return
	}
	tasks[index].Done = true
	saveTasks(tasks)
}

func DeleteTask(index int) {
	tasks := loadTasks()
	if index < 0 || index >= len(tasks) {
		fmt.Println("Invalid task index")
		return
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	saveTasks(tasks)
}

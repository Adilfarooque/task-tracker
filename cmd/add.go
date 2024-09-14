package cmd

import (
	"fmt"
	"time"

	"github.com/Adilfarooque/task-tracker/task"
)

func AddTask(description string){
	tasks,_ := task.LoadTasks()
	id := len(tasks) + 1
	newTask := task.Task{
		ID: id,
		Description: description,
		Status: "todo",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	tasks = append(tasks, newTask)
	task.SaveTasks(tasks)
	fmt.Printf("Task added successfully (ID: %d)\n",id)
}
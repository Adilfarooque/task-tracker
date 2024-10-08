package cmd

import (
	"fmt"
	"time"

	"github.com/Adilfarooque/task-tracker/task"
)

func UpdateTask(id int, description string) {
	tasks, _ := task.LoadTasks()
	for i, tsk := range tasks {
		if tsk.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			task.SaveTasks(tasks)
			fmt.Printf("Task %d update successfully.\n", id)
			return
		}
	}
	fmt.Printf("Task with ID %d not found. \n", id)
}

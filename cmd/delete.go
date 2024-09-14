package cmd

import (
	"fmt"

	"github.com/Adilfarooque/task-tracker/task"
)

func DeleteTask(id int) {
	tasks, _ := task.LoadTasks()
	for i, tsk := range tasks {
		if tsk.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			task.SaveTasks(tasks)
			fmt.Printf("Task %d deleted successfully.\n", id)
			return
		}

	}
	fmt.Printf("Task with ID %d not found.\n", id)
}

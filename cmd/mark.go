package cmd

import (
	"fmt"
	"time"

	"github.com/Adilfarooque/task-tracker/task"
)

func MarkTask(id int, status string) {
	tasks, _ := task.LoadTasks()
	for i, tsk := range tasks {
		if tsk.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			task.SaveTasks(tasks)
			fmt.Printf("Task %d marked as %s. \n", id, status)
			return
		}
	}
	fmt.Printf("Task with ID %d not found. \n",id)
}

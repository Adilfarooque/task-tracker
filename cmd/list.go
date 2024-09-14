package cmd

import (
	"fmt"

	"github.com/Adilfarooque/task-tracker/task"
)

func ListTasks(statusFilter string) {
	tasks, _ := task.LoadTasks()
	for _, tsk := range tasks {
		if statusFilter == "" || tsk.Status == statusFilter {
			fmt.Printf("ID: %d, Description : %s, Status: %s, CreatedAt: %s\n", tsk.ID, tsk.Description, tsk.Status, tsk.CreatedAt)
		}
	}
}

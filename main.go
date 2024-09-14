package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Adilfarooque/task-tracker/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command.")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task description.")
		} else {
			cmd.AddTask(os.Args[2])
		}
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <id> <new description>")
		} else {
			id, _ := strconv.Atoi(os.Args[2])
			cmd.UpdateTask(id, os.Args[2])
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
		} else {
			id, _ := strconv.Atoi(os.Args[2])
			cmd.DeleteTask(id)
		}
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-in-progress <id>")
		}else{
			id,_ := strconv.Atoi(os.Args[2])
			cmd.MarkTask(id,"in-progress")
		}
	case "mark-done":
		if len(os.Args) < 3{
			fmt.Println("Usage: task-cli mark-done <id>")
		}else{
			id,_ := strconv.Atoi(os.Args[2])
			cmd.MarkTask(id,"done")
		}
	case "list":
		if len(os.Args) == 2{
			cmd.ListTasks("")
		}else{
			cmd.ListTasks(os.Args[2])
		}

	default:
		fmt.Println("Unknown command:",command)
	}

}

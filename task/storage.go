package task

import (
	"encoding/json"
	"os"
)

const filePath = "./tasks.json"

func LoadTasks() ([]Task, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return []Task{}, nil
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(file,&tasks)
	if err != nil{
		return nil,err
	}
	return tasks,nil
}

func SaveTasks(tasks []Task)error{
	data,err := json.MarshalIndent(tasks,"","  ")
	if err != nil{
		return err
	}

	return os.WriteFile(filePath,data,0644)
}
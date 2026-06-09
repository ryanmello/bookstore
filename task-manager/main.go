package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Task struct {
	Name string `json:"name"`
	IsCompleted bool `json:"isCompleted"`
}

func main() {
	// read the command line input from the user
	args := os.Args
	action := args[1]

	// actions --> add, list, complete, delete
	switch action {
	case "add":
		taskName := args[2]

		newTask := Task{
			Name: taskName,
			IsCompleted: false,
		}

		var tasks []Task

		data, err := os.ReadFile("data.json")
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		err = json.Unmarshal(data, &tasks)
		if err != nil {
			fmt.Println("Error at unmarshal:", err)
			return
		}

		tasks = append(tasks, newTask)

		updatedData, err := json.MarshalIndent(tasks, "", " ")
		if err != nil {
			fmt.Println("Error converting to JSON:", err)
			return
		}

		err = os.WriteFile("data.json", updatedData, 0644)
		if err != nil {
			fmt.Println("Error writing to file: ", err)
			return
		}

		fmt.Println("Task added", taskName)
	case "list":
		var tasks []Task

		data, err := os.ReadFile("data.json")
		if err != nil {
			return
		}
		err = json.Unmarshal(data, &tasks)

		for i, task := range tasks {
			fmt.Println(i, task.Name, task.IsCompleted)
		}
	case "complete":
		taskId := args[2]
		index, err := strconv.Atoi(taskId)

		var tasks []Task

		data, err := os.ReadFile("data.json")
		if err != nil {
			fmt.Println("Error reading file", err)
			return
		}

		err = json.Unmarshal(data, &tasks)
		if err != nil {
			fmt.Println("Error at unmarshal", err)
			return
		}

		tasks[index].IsCompleted = true

		updatedData, err := json.MarshalIndent(tasks, "", " ")
		if err != nil {
			fmt.Println("Error at unmarshal index", err)
			return
		}
		
		err = os.WriteFile("data.json", updatedData, 0644)
		if err != nil {
			fmt.Println("Error writing file", err)
			return
		}

		fmt.Println("Completed task", taskId)
	case "delete":
		taskId := args[2]
		fmt.Println("Deleting task: ", taskId)
	default:
		fmt.Println("I have no idea what ypu're doing: ", action)
	}
}

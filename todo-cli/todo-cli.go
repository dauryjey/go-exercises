package main

import (
	"fmt"
)

func main() {
	var tasks []string
	for {
		fmt.Println("Welcome to TODO CLI")
		fmt.Println("Please select an option")
		fmt.Println("1. Add new task")
		fmt.Println("2. List all tasks")
		fmt.Println("3. Delete task")
		fmt.Println("4. Update task")
		fmt.Println("5. Exit")

		var option int

		fmt.Scan(&option)

		switch option {
		case 1:
			var newTask string
			fmt.Println("Enter the task you want to add")
			fmt.Scan(&newTask)
			tasks = addNewTask(newTask, tasks)

		case 2:
			listAllTasks(tasks)

		case 3:
			deleteTask()

		case 4:
			updateTask()

		case 5:
			fmt.Println("Exiting the application")
			return
		}
	}
}

func addNewTask(task string, tasks []string) []string {
	return append(tasks, task)
}

func listAllTasks(tasks []string) {
	for i, task := range tasks {
		fmt.Println("TaskID:", i+1, "Task:", task)
	}
}

func deleteTask() {
	fmt.Println("Delete task")
}

func updateTask() {
	fmt.Println("Update task")
}

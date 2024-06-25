package main

import (
	"fmt"
)

func main() {
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
		addNewTask()

	case 2:
		listAllTasks()

	case 3:
		deleteTask()

	case 4:
		updateTask()

	case 5:
		fmt.Println("Exiting the application")
	}
}

func addNewTask() {
	fmt.Println("Add new task")
}

func listAllTasks() {
	fmt.Println("List all tasks")
}

func deleteTask() {
	fmt.Println("Delete task")
}

func updateTask() {
	fmt.Println("Update task")
}

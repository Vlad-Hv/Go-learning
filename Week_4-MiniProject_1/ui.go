package main

import (
	"fmt"
)

func getTaskName() string {
	var name string
	fmt.Println("Enter name for your task:")
	fmt.Scanln(&name)
	return name
}

func getTaskDescription() string {
	var description string
	fmt.Println("Enter task description:")
	fmt.Scanln(&description)
	return description
}

func getPriority() (int, error) {
	var priority int
	fmt.Println("Enter priority(1 - low, 2 - medium, 3 - high):")
	_, err := fmt.Scanln(&priority)
	return priority, err
}

func mainMenu() (int, error) {
	var option int
	fmt.Println("---Menu---\n1. Create task\n2. Show all tasks\n3. Start task\n4. Complete task\n5. Delete task\n6. Show task by ID\n7. Exit\nChoose the option:")
	_, err := fmt.Scanln(&option)
	return option, err
}

func chooseID(word string) (int, error) {
	var option int
	fmt.Println("Enter number of task that you want to", word, ":")
	_, err := fmt.Scanln(&option)
	option -= 1
	return option, err
}

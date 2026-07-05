package main

import (
	"errors"
	"fmt"
)

func main() {

	var tasks []string

	for {
		option, err := menuAndGetOption()

		if err != nil {
			fmt.Println(err)
			return
		}

		if option == 4 {
			break
		}
		tasks, err = optionsSwitch(tasks, option)

		if err != nil {
			fmt.Println(err)
		}
	}
	fmt.Println("ByeBye:)")
}

func menuAndGetOption() (int, error) {
	var option int
	fmt.Print("\n\n===== To-Do List =====\n1. Show tasks\n2. Add task\n3. Delete task\n4. Exit\nChoose option: ")
	_, err := fmt.Scanln(&option)

	if err != nil {
		return 0, errors.New("incorrect choice type")
	}
	return option, nil
}

func showTask(tasks []string) {

	if len(tasks) == 0 {
		fmt.Println("No tasks yet")
	} else {
		for index, numbers := range tasks {
			fmt.Println(index+1, numbers)
		}
	}
}

func addTask(tasks []string) ([]string, error) {
	var task string
	fmt.Println("Enter your task:")
	fmt.Scanln(&task)

	if task == "" {
		return tasks, errors.New("empty task")
	}

	return append(tasks, task), nil
}

func deleteTask(tasks []string) ([]string, error) {
	var index int
	fmt.Println("Enter index of task which you wanna delete from 0 to", len(tasks)-1)
	_, err := fmt.Scanln(&index)

	if err != nil {
		return tasks, errors.New("incorrect index type")
	}

	if index < 0 || index >= len(tasks) {
		return tasks, fmt.Errorf("incorrect index: %d", index)
	}

	tasks = append(tasks[:index], tasks[index+1:]...)

	return tasks, nil
}

func optionsSwitch(tasks []string, option int) ([]string, error) {
	switch option {
	case 1:
		showTask(tasks)
		return tasks, nil
	case 2:
		tasks, err := addTask(tasks)

		if err != nil {
			return tasks, fmt.Errorf("cannot continue: %w", err)
		}
		success()
		return tasks, nil

	case 3:
		tasks, err := deleteTask(tasks)

		if err != nil {
			return tasks, fmt.Errorf("cannot continue: %w", err)
		}
		success()
		return tasks, nil

	default:
		return tasks, fmt.Errorf("incorrect option: %d", option)
	}
}

func success() {
	fmt.Println("Tasks list changed successfully!")
}

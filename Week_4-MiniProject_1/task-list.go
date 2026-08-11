package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Task struct {
	ID          uuid.UUID
	Name        string
	Description string
	Priority    string
	Status      string
}

func createTasksList() []Task {
	var tasks []Task
	return tasks
}

func createTask(tasks *[]Task, name, description, priority, status string, id uuid.UUID) {
	task := Task{
		ID:          id,
		Name:        name,
		Description: description,
		Priority:    priority,
		Status:      status,
	}
	*tasks = append(*tasks, task)
}

func printAllTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("tasks is clear")
		return
	}
	for index, task := range tasks {
		fmt.Println("Task", index+1)
		fmt.Println("ID:", task.ID)
		fmt.Println("Name:", task.Name)
		fmt.Println("Description:", task.Description)
		fmt.Println("Priority:", task.Priority)
		fmt.Println("Status:", task.Status, "\n ")
	}
}

func (task *Task) startTask() {
	task.Status = "in_progress"
}

func (task *Task) completeTask() {
	task.Status = "done"
}

func deleteTask(tasks *[]Task, index int) {
	*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
}

func (task Task) printTaskInfo() {
	fmt.Print("Task #", task.ID, "\n ")
	fmt.Println("Name:", task.Name)
	fmt.Println("Description:", task.Description)
	fmt.Println("Priority:", task.Priority)
	fmt.Println("Status:", task.Status, "\n ")
}

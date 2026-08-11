package main

import (
	"fmt"
)

func main() {
	tasks := createTasksList()
	priorityMap := makePriorityMap()
	indexToUUID := modifyIndexToUUID(&tasks)
	UuidToTask := MapUuidToTask(&tasks)
	for {

		option, err := mainMenu()
		err = validateMainMenu(err)
		if err != nil {
			fmt.Println(err)
		}

		if option == 7 {
			break
		}

		switch option {
		case 1:
			name := getTaskName()
			description := getTaskDescription()
			priority, err := getPriority()
			err = validateTaskCreating(name, description, priority, err)
			id := generateID()

			if err != nil {
				fmt.Println(err)
				continue
			}

			createTask(&tasks, name, description, priorityMap[priority], "todo", id)
			syncTaskMaps(&tasks, &indexToUUID, &UuidToTask)
		case 2:
			printAllTasks(tasks)

		case 3:
			option, err := chooseID("start")
			err = validateID(tasks, option, err)

			if err != nil {
				fmt.Println(err)
				continue
			}
			task := UuidToTask[indexToUUID[option]]
			err = task.startValidate()

			if err != nil {
				fmt.Println(err)
				continue
			}
			err = task.pointerValidate()
			if err != nil {
				fmt.Println(err)
				return
			}

			task.startTask()

		case 4:
			option, err := chooseID("complete")
			err = validateID(tasks, option, err)

			if err != nil {
				fmt.Println(err)
				continue
			}
			task := UuidToTask[indexToUUID[option]]
			err = task.completeValidate()

			if err != nil {
				fmt.Println(err)
				continue
			}
			err = task.pointerValidate()
			if err != nil {
				fmt.Println(err)
				return
			}
			task.completeTask()

		case 5:
			option, err := chooseID("complete")
			err = validateID(tasks, option, err)

			if err != nil {
				fmt.Println(err)
				continue
			}
			deleteTask(&tasks, option)
			syncTaskMaps(&tasks, &indexToUUID, &UuidToTask)

		case 6:
			option, err := chooseID("check")
			err = validateID(tasks, option, err)

			if err != nil {
				fmt.Println(err)
				continue
			}
			task := UuidToTask[indexToUUID[option]]
			task.printTaskInfo()
		}
	}
	fmt.Println("TaskFlow turned off")
}

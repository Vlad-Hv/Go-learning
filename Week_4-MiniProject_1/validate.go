package main

import "errors"

func validateTaskCreating(name, description string, priority int, err error) error {
	if err != nil {
		return errors.New("incorrect input type")
	}

	if name == "" || description == "" {
		return errors.New("name and description mustnot be empty")
	}

	if priority < 1 || priority > 3 {
		return errors.New("incorrect priority")
	}

	return nil
}

func validateMainMenu(err error) error {
	if err != nil {
		return errors.New("incorrect input type")
	}
	return nil
}

func validateID(tasks []Task, option int, err error) error {
	var helper int
	if err != nil {
		return errors.New("incorrect input type")
	}

	for index := range tasks {
		if index == option {
			helper++
		}
	}

	if helper == 0 {
		return errors.New("invalid option")
	}

	return nil

}

func (task Task) startValidate() error {
	if task.Status == "in_progress" {
		return errors.New("task already in progress")
	}
	if task.Status == "done" {
		return errors.New("task already done")
	}

	return nil
}

func (task Task) completeValidate() error {
	if task.Status == "done" {
		return errors.New("task already done")
	}

	if task.Status == "todo" {
		return errors.New("the task must be in progresss, if you want to complete it")
	}

	return nil
}

func (task *Task) pointerValidate() error {
	if task == nil {
		return errors.New("task is nil")
	}

	return nil
}

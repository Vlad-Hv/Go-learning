package main

import (
	"errors"
	"fmt"
)

func validating(name, deviceName, problem string, priority int, err error) error {
	if err != nil {
		return errors.New("incorrect input type")
	}

	if name == "" || deviceName == "" || problem == "" {
		return errors.New("input must not be empty")
	}

	if priority < 1 || priority > 3 {
		return fmt.Errorf("priority must be 1-3, not %v", priority)
	}

	return nil
}

func validateUserOption(option int, err error) error {
	if err != nil {
		return errors.New("incorrect option type")
	}

	if option < 1 || option > 2 {
		return fmt.Errorf("there is no option %v", option)
	}
	return nil
}

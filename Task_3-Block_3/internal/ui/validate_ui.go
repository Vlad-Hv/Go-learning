package ui

import (
	"errors"
)

func ValidateMenuOption(option int, err error) error {
	if err != nil {
		return errors.New("incorrect input type")
	}

	if option < 1 || option > 5 {
		return errors.New("invalid option")
	}

	return nil
}

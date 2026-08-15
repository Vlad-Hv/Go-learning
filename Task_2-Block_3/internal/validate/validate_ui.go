package validate

import (
	"errors"
)

func ValidateExInfo(category string, amount int, err error) error {
	if err != nil {
		return errors.New("incorrect input type")
	}

	if category == "" {
		return errors.New("category mustnot be empty")
	}

	if amount <= 0 {
		return errors.New("amount cannot be empty")
	}

	return nil
}

func ValidateMenuOption(option int, err error) error {
	if err != nil {
		return errors.New("incorrect input type")
	}

	if option < 1 || option > 4 {
		return errors.New("invalid option")
	}

	return nil
}

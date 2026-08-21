package validation

import (
	"errors"
)

func ValidateMenuOption(option int, err error) error {
	if err != nil {
		return errors.New("incorrect option type")
	}

	if option < 1 || option > 5 {
		return errors.New("invalid option")
	}

	return nil
}

func ValidateCurrentFloor(err error) error {
	if err != nil {
		return errors.New("incorrect type")
	}

	return nil
}

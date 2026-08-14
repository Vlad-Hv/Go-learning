package ui

import (
	"errors"
)

func ValidateUi(option int, err error) error {
	if err != nil {
		return errors.New("incorrect option type")
	}

	if option < 1 || option > 5 {
		return errors.New("invalid option")
	}

	return nil
}

func ValidateTicketOption(option, lenthOptions int, err error) error {
	if err != nil {
		return errors.New("incorrect input type")
	}

	if option < 1 || option > lenthOptions {
		return errors.New("invalid option")
	}

	return nil
}

func ValidateName(ownerName string) error {
	if ownerName == "" {
		return errors.New("name mustnot be empty")
	}
	return nil
}

func ValidateUseOneTicket(option, lenthOptions int, err error) error {
	if err != nil {
		return errors.New("incorrect input type")
	}

	if option < 1 || option > lenthOptions {
		return errors.New("invalid option")
	}

	return nil
}

package ticket

import (
	"errors"
)

func ValidateUserVisit(user Ticket) error {
	if user.VisitsLeft <= 0 {
		return errors.New("user do not have enough visits")
	}
	return nil
}

func ValidateTicketStatus(user Ticket) error {
	if user.Status == "frozen" {
		return errors.New("your ticket is already frozen")
	}
	return nil
}

func ValidatePosibilityToUseTicket(timeTicket Ticket) error {
	if timeTicket.Status == "frozen" {
		return errors.New("you are not able to use visits, because your ticket is frozen")
	}
	return nil
}

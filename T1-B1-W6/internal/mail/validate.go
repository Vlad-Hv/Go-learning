package mail

import (
	"errors"
)

var (
	ErrMailDelivered = errors.New("mail is already delivered")
)

func alreadyDelivered(mail Mail) error {
	if mail.State == "delivered" {
		return ErrMailDelivered
	}

	return nil
}

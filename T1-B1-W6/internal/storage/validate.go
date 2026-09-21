package storage

import (
	"delivery/internal/mail"
	"errors"
)

var (
	errIDUsing              = errors.New("ID alreagy in using")
	errIDInvalid            = errors.New("invalid ID")
	errRecipientNameInvalid = errors.New("invelid recipient name")
	errWeightInvalid        = errors.New("weight must be more than zero")
)

func validateID(storage map[int]*mail.Mail, id int, mail mail.Mail) error {
	_, ok := storage[id]

	if ok {
		return errIDUsing
	}

	if id <= 0 {
		return errIDInvalid
	}

	if mail.Recipient == "" {
		return errRecipientNameInvalid
	}

	if mail.Weight <= 0 {
		return errWeightInvalid
	}

	return nil
}

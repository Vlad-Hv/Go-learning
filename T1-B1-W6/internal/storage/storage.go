package storage

import (
	"delivery/internal/mail"
)

func Create() map[int]*mail.Mail {
	return make(map[int]*mail.Mail)
}

func Add(storage map[int]*mail.Mail, id int, mail mail.Mail) error {
	err := validateID(storage, id, mail)
	if err != nil {
		return err
	}

	storage[id] = &mail
	return nil
}

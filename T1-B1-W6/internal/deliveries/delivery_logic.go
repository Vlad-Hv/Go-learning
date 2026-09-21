package deliveries

import (
	"delivery/internal/history"
	"delivery/internal/mail"
	"errors"
	"fmt"
)

var (
	errAddressIncorrect = errors.New("incorrect mail address")
)

type delivery interface {
	Delivery(mail *mail.Mail) (int, error)
}

func Delivery(deliverier delivery, mail *mail.Mail, history *history.History) error {
	if mail == nil {
		return errAddressIncorrect
	}
	price, err := deliverier.Delivery(mail)
	if err != nil {
		return err
	}

	err = mail.MarkAsDelivered()
	if err != nil {
		return err
	}
	mail.Price = price
	fmt.Println(price)
	err = history.Add("mail to", mail.Recipient, "Delivered successfully")
	if err != nil {
		return err
	}
	return nil
}

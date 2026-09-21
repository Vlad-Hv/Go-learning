package deliveries

import (
	"delivery/internal/mail"
	"errors"
)

var (
	errDeliveryBigWeight = errors.New("delivery person can not delivery mail more than 10 kg")
	errServiceBigWeight  = errors.New("over bif mail weight")
)

func validateDeliverier(mail mail.Mail) error {
	if mail.Weight > 10 {
		return errDeliveryBigWeight
	}

	return nil
}

func validateService(mail mail.Mail) error {
	if mail.Weight > 100 {
		return errServiceBigWeight
	}

	return nil
}

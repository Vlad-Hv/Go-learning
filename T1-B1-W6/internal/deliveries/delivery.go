package deliveries

import (
	"delivery/internal/mail"
)

type Person struct {
	Price int
}

type Service struct {
	Price int
}

func (d Person) Delivery(mail *mail.Mail) (int, error) {
	const mandatoryFee int = 100
	err := validateDeliverier(*mail)
	if err != nil {
		return 0, err
	}

	price := mandatoryFee + (d.Price * mail.Weight)
	return price, nil
}

func (d Service) Delivery(mail *mail.Mail) (int, error) {
	const mandatoryFee int = 300

	err := validateService(*mail)

	if err != nil {
		return 0, err
	}

	price := mandatoryFee + (d.Price * mail.Weight)
	return price, nil
}

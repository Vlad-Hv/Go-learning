package card

import (
	"errors"
)

func (c Card) payValidate(price int) error {

	if !c.IsAvailable {
		return errors.New("card is not available")
	}

	if price <= 0 {
		return errors.New("price must not be zero or less")
	}

	if (c.Balance - price) < 0 {
		return errors.New("too big price")
	}

	return nil

}

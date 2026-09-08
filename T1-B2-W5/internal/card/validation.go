package card

import (
	"errors"
)

var (
	errBlockedCard    = errors.New("card is not available")
	errInvalidPrice   = errors.New("price must not be zero or less")
	errNotEnoughMoney = errors.New("too big price")
)

func (c Card) payValidate(price int) error {

	if !c.IsAvailable {
		return errBlockedCard
	}

	if price <= 0 {
		return errInvalidPrice
	}

	if (c.Balance - price) < 0 {
		return errNotEnoughMoney
	}

	return nil

}

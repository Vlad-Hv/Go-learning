package card

import (
	"errors"
)

var (
	errCardBlocked    = errors.New("card is blocked")
	errInvalidPrice   = errors.New("invalid price")
	errNotEnoughMoney = errors.New("insufficient balance")
)

func (c *Card) validatePay(price int) error {

	if c.IsBlocked {
		return errCardBlocked
	}

	if price <= 0 {
		return errInvalidPrice
	}

	if price > c.Balance {
		return errNotEnoughMoney
	}

	return nil

}

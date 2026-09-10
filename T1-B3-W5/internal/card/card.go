package card

type Card struct {
	Balance   int
	IsBlocked bool
}

func (c *Card) Pay(price int) error {
	err := c.validatePay(price)
	if err != nil {
		return err
	}

	c.Balance -= price
	return nil
}

package card

type Card struct {
	Balance     int
	IsAvailable bool
}

func (c *Card) Pay(price int) error {
	err := c.payValidate(price)
	if err != nil {
		return err
	}

	c.Balance -= price
	return nil
}

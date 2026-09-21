package mail

type Mail struct {
	Recipient string
	Weight    int
	State     string
	Price     int
}

func (m *Mail) MarkAsDelivered() error {
	err := alreadyDelivered(*m)
	if err != nil {
		return err
	}
	m.State = "delivered"
	return nil
}

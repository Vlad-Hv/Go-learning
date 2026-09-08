package event

import "errors"

var (
	ErrInvalidAmount = errors.New("amount must be more than zero")
)

func (s EventReservationSystem) ValidateBooking(name string, amount int) error {

	if name == "" {
		return errors.New("name must not be empty")
	}

	if amount <= 0 {
		return ErrInvalidAmount
	}

	if (s.FreePlace - amount) < 0 {
		return errors.New("not enough free space")
	}

	return nil
}

func (s EventReservationSystem) ValidateCancelBooking(name string) error {
	if name == "" {
		return errors.New("name must not be empty")
	}

	_, ok := s.Reservations[name]

	if !ok {
		return errors.New("client must exist")
	}
	return nil

}

package eventreservation

import "errors"

var (
	errInvalidAmount      = errors.New("amount must be more than zero")
	errEmptyName          = errors.New("name must mot be empty")
	errNotEnoughFreeSpace = errors.New("not enough free space")
	errNonexistentClient  = errors.New("client must exist")
)

func (s EventReservationSystem) ValidateBooking(name string, amount int) error {

	if name == "" {
		return errEmptyName
	}

	if amount <= 0 {
		return errInvalidAmount
	}

	if (s.FreePlace - amount) < 0 {
		return errNotEnoughFreeSpace
	}

	return nil
}

func (s EventReservationSystem) ValidateCancelBooking(name string) error {
	if name == "" {
		return errEmptyName
	}

	_, ok := s.Reservations[name]

	if !ok {
		return errNonexistentClient
	}
	return nil

}

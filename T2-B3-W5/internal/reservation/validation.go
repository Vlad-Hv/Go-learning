package reservation

import "errors"

var (
	errInvalidAmount      = errors.New("invalid amount")
	errNameEmpty          = errors.New("name must not be empty")
	errNotEnoughFreeSpace = errors.New("not enough free space")
	errNonexistentClient  = errors.New("client must exist")
)

func (s System) validateBooking(name string, amount int) error {

	if name == "" {
		return errNameEmpty
	}

	if amount <= 0 {
		return errInvalidAmount
	}

	if s.FreePlace < amount {
		return errNotEnoughFreeSpace
	}

	return nil
}

func (s System) validateCancel(name string) error {
	if name == "" {
		return errNameEmpty
	}

	_, ok := s.Reservations[name]

	if !ok {
		return errNonexistentClient
	}
	return nil

}

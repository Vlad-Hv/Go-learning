package event

import (
	"fmt"
)

type Reservation struct {
	Name        string
	PlaceAmount int
}

func (s *EventReservationSystem) Booking(name string, amount int) error {
	const historyMessageChoice int = 1
	err := s.ValidateBooking(name, amount)

	if err != nil {
		return err
	}
	reserv, ok := s.Reservations[name]

	if !ok {
		reservation := Reservation{Name: name, PlaceAmount: amount}
		s.Reservations[name] = &reservation
		s.FreePlace -= amount
	} else {
		reserv.PlaceAmount += amount
		s.FreePlace -= amount
	}

	historyMessage := createHistoryMessage(name, historyMessageChoice, amount)
	s.History = append(s.History, historyMessage)
	return nil
}

func (s *EventReservationSystem) CancellationOfBooking(name string) error {
	const historyMessageChoice int = 3
	err := s.ValidateCancelBooking(name)

	if err != nil {
		return err
	}
	amount := s.Reservations[name].PlaceAmount
	s.FreePlace += s.Reservations[name].PlaceAmount
	delete(s.Reservations, name)
	historyMessage := createHistoryMessage(name, historyMessageChoice, amount)
	s.History = append(s.History, historyMessage)

	return nil
}

func createHistoryMessage(name string, choice, amount int) string {
	if choice == 1 {
		return fmt.Sprintf("%v booked %d places", name, amount)
	} else {
		return fmt.Sprintf("%v unbooked %d places", name, amount)
	}

}

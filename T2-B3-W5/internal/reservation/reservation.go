package reservation

import (
	"fmt"
)

type Reservation struct {
	Name        string
	PlaceAmount int
}

func (s *System) Book(name string, amount int) error {
	const action string = "book" //using to make correct history message down
	err := s.validateBooking(name, amount)

	if err != nil {
		return fmt.Errorf("book reservation: %w", err)
	}

	reserv, ok := s.Reservations[name]

	if !ok {
		reservation := Reservation{Name: name, PlaceAmount: amount}
		s.Reservations[name] = &reservation
	} else {
		reserv.PlaceAmount += amount
	}

	s.FreePlace -= amount

	historyMessage := createHistoryMessage(name, action, amount)
	s.History = append(s.History, historyMessage)
	return nil
}

func (s *System) Cancel(name string) error {
	const action string = "unbook" //as well using to create correct history message
	err := s.validateCancel(name)

	if err != nil {
		return err
	}
	amount := s.Reservations[name].PlaceAmount
	s.FreePlace += amount
	delete(s.Reservations, name)
	historyMessage := createHistoryMessage(name, action, amount)
	s.History = append(s.History, historyMessage)

	return nil
}

func createHistoryMessage(name, choice string, amount int) string {
	if choice == "book" {
		return fmt.Sprintf("%v booked %d places", name, amount)
	} else {
		return fmt.Sprintf("%v unbooked %d places", name, amount)
	}

}

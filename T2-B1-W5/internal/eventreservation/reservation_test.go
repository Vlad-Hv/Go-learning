package event

import "testing"

func TestBookingSuccess(t *testing.T) {
	system := EventReservationSystem{Capacity: 20, FreePlace: 20, Reservations: make(map[string]*Reservation)}

	err := system.Booking("vlad", 5)
	if err != nil {
		t.Fatalf("unexpected error, got %v", err)
	}

	reservation, ok := system.Reservations["vlad"]

	if !ok {
		t.Fatal("expected reservation for vlad")
	}

	if reservation.PlaceAmount != 5 {
		t.Errorf("expected 5, got %d", reservation.PlaceAmount)
	}

	if system.FreePlace != 15 {
		t.Errorf("expected 15 free places, actual %d", system.FreePlace)
	}

	if len(system.Reservations) == 0 {
		t.Errorf("expected 1 reservation, got 0") // кай, это вопрос к тебе: как в таких случаях проверять, через длину или через типо if *system.Reservations[vlad] != {Name: "vlad", PlaceAmount: 5}
	}

	if len(system.History) == 0 {
		t.Errorf("expected 1 history message, got 0")
	}
}

func TestBookingFullFreePlaces(t *testing.T) {
	system := EventReservationSystem{Capacity: 20, FreePlace: 20, Reservations: make(map[string]*Reservation)}

	err := system.Booking("vlad", 20)
	if err != nil {
		t.Fatalf("unexpected error, got %v", err)
	}

	reserv, ok := system.Reservations["vlad"]
	if !ok {
		t.Fatal("expected reservation for vlad")
	}

	if reserv.PlaceAmount != 20 {
		t.Errorf("expected 20, actual %d", reserv.PlaceAmount)
	}

	if system.FreePlace != 0 {
		t.Errorf("expected 0 free places, actual %d", system.FreePlace)
	}

	if len(system.Reservations) == 0 {
		t.Errorf("expected 1 reservation, got 0") // кай, это вопрос к тебе: как в таких случаях проверять, через длину или через типо if *system.Reservations[vlad] != {Name: "vlad", PlaceAmount: 5}
	}

	if len(system.History) == 0 {
		t.Errorf("expected 1 history message, got 0")
	}
}

func TestBookingWithZeroBookingPlaces(t *testing.T) {
	system := EventReservationSystem{Capacity: 10, FreePlace: 10, Reservations: make(map[string]*Reservation)}

	err := system.Booking("vlad", 0)
	if err == nil {
		t.Fatalf("expected an a error")
	}

	_, ok := system.Reservations["vlad"]

	if ok {
		t.Fatal("unexpected reservation for vlad")
	}

	if system.FreePlace != 10 {
		t.Errorf("expected 10, got %d", system.FreePlace)
	}

	if len(system.Reservations) != 0 {
		t.Errorf("expected lenth reservations 0, got %d", len(system.Reservations))
	}

	if len(system.History) != 0 {
		t.Errorf("expected empty history")
	}
}

func TestBookingWithLessThanZeroBookingPlaces(t *testing.T) {
	system := EventReservationSystem{Capacity: 10, FreePlace: 10, Reservations: make(map[string]*Reservation)}

	err := system.Booking("vlad", -3)
	if err == nil {
		t.Fatalf("expected an a error, got %v", err)
	}

	_, ok := system.Reservations["vlad"]

	if ok {
		t.Fatal("unexpected reservation for vlad")
	}

	if system.FreePlace != 10 {
		t.Errorf("expected 10, got %d", system.FreePlace)
	}

	if len(system.Reservations) != 0 {
		t.Errorf("expected lenth reservations 0, got %d", len(system.Reservations))
	}

	if len(system.History) != 0 {
		t.Errorf("expected empty history")
	}
}

func TestBookingWithNotEnoughFreePlaces(t *testing.T) {
	system := EventReservationSystem{
		Capacity:     6,
		FreePlace:    6,
		Reservations: make(map[string]*Reservation),
	}

	err := system.Booking("vlad", 10)

	if err == nil {
		t.Fatalf("expected an error")
	}

	_, ok := system.Reservations["vlad"]

	if ok {
		t.Fatal("unexpected order for vlad")
	}

	if system.FreePlace != 6 {
		t.Errorf("expected 6, got %d", system.FreePlace)
	}

	if len(system.Reservations) != 0 {
		t.Errorf("expected 0, got %d", len(system.Reservations))
	}

	if len(system.History) != 0 {
		t.Errorf("expected 0, got %d", len(system.History))
	}
}

func TestUnbookingSuccess(t *testing.T) {
	system := EventReservationSystem{
		Capacity:  10,
		FreePlace: 4,
		Reservations: map[string]*Reservation{
			"vlad": {Name: "vlad", PlaceAmount: 6},
		},
	}

	err := system.CancellationOfBooking("vlad")

	if err != nil {
		t.Fatalf("unexpected error, got %v", err)
	}

	_, ok := system.Reservations["vlad"]

	if ok {
		t.Fatal("unexpected order for vlad")
	}

	if system.FreePlace != 10 {
		t.Errorf("expected 10 free places, actual %d", system.FreePlace)
	}
	if len(system.Reservations) != 0 {
		t.Errorf("expected 0, got %d", len(system.Reservations))
	}

	if len(system.History) != 1 {
		t.Errorf("expected 1, got %d", len(system.History))
	}
}

func TestUnbookingNonExistentBooking(t *testing.T) {
	system := EventReservationSystem{
		Capacity:  10,
		FreePlace: 5,
		Reservations: map[string]*Reservation{
			"vlad": {
				Name:        "vlad",
				PlaceAmount: 5,
			},
		},
	}

	err := system.CancellationOfBooking("shenia")
	if err == nil {
		t.Fatalf("expected an error")
	}

	_, ok := system.Reservations["vlad"]

	if !ok {
		t.Fatal("expected reservation for vlad")
	}

	if system.FreePlace != 5 {
		t.Errorf("expected 5, actual 10")
	}

	if len(system.Reservations) == 0 {
		t.Errorf("expected 1, got 0")
	}

	if len(system.History) == 1 {
		t.Errorf("expected 0, got 1")
	}
}

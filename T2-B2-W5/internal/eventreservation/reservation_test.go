package eventreservation

import (
	"errors"
	"testing"
)

func TestBooking(t *testing.T) {
	tests := []struct {
		name                   string
		capacity               int
		freePlace              int
		reservations           map[string]*Reservation
		bookerName             string
		placeAmount            int
		expectedFreePlace      int
		expectedReservedPlaces int
		expectedOk             bool
		expectedHistoryLen     int
		wantErr                error
	}{ //успешное, на пустое имя, на количество 0, на количество менньше нуля, на больше мест чем есть, на заниманеи всех мест,
		{name: "successful booking", capacity: 20, freePlace: 20, reservations: make(map[string]*Reservation), bookerName: "vlad", placeAmount: 10, expectedFreePlace: 10, expectedReservedPlaces: 10, expectedOk: true, expectedHistoryLen: 1, wantErr: nil},
		{name: "validate empty reserver name", capacity: 20, freePlace: 20, reservations: make(map[string]*Reservation), bookerName: "", placeAmount: 10, expectedFreePlace: 20, expectedReservedPlaces: 0, expectedOk: false, expectedHistoryLen: 0, wantErr: errEmptyName},
		{name: "reserve zero places", capacity: 20, freePlace: 20, reservations: make(map[string]*Reservation), bookerName: "vlad", placeAmount: 0, expectedFreePlace: 20, expectedReservedPlaces: 0, expectedOk: false, expectedHistoryLen: 0, wantErr: errInvalidAmount},
		{name: "reserve less than zero places", capacity: 20, freePlace: 20, reservations: make(map[string]*Reservation), bookerName: "vlad", placeAmount: -1, expectedFreePlace: 20, expectedReservedPlaces: 0, expectedOk: false, expectedHistoryLen: 0, wantErr: errInvalidAmount},
		{name: "reserve more than available", capacity: 20, freePlace: 20, reservations: make(map[string]*Reservation), bookerName: "vlad", placeAmount: 30, expectedFreePlace: 20, expectedReservedPlaces: 0, expectedOk: false, expectedHistoryLen: 0, wantErr: errNotEnoughFreeSpace},
		{name: "reserve full place", capacity: 20, freePlace: 20, reservations: make(map[string]*Reservation), bookerName: "vlad", placeAmount: 20, expectedFreePlace: 0, expectedReservedPlaces: 20, expectedOk: true, expectedHistoryLen: 1, wantErr: nil},
		{name: "add places to existent booker", capacity: 20, freePlace: 10, reservations: map[string]*Reservation{"vlad": {Name: "vlad", PlaceAmount: 10}}, bookerName: "vlad", placeAmount: 5, expectedFreePlace: 5, expectedReservedPlaces: 15, expectedOk: true, expectedHistoryLen: 1, wantErr: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			system := EventReservationSystem{Capacity: tt.capacity, FreePlace: tt.freePlace, Reservations: tt.reservations}

			err := system.Booking(tt.bookerName, tt.placeAmount)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if system.FreePlace != tt.expectedFreePlace {
				t.Errorf("expected %d free places, actual %d", tt.expectedFreePlace, system.FreePlace)
			}

			reservation, ok := system.Reservations[tt.bookerName]

			if ok != tt.expectedOk {
				t.Fatalf("expected availability of reservation as %v, got %v", tt.expectedOk, ok)
			} //подумать за проверки на наличие резервации, и за в целом цикл проверки

			if tt.expectedOk {
				if reservation.PlaceAmount != tt.expectedReservedPlaces {
					t.Errorf("expected reserved place %d, actual %d", tt.expectedReservedPlaces, reservation.PlaceAmount)
				}

			}

			if len(system.History) != tt.expectedHistoryLen {
				t.Errorf("expected history length %d, actual %d", tt.expectedHistoryLen, len(system.History))
			}

		})
	}
}

func TestCancellationOfBooking(t *testing.T) {
	tests := []struct {
		name                       string
		capacity                   int
		freePlaces                 int
		expectedFreePlaces         int
		unbookerName               string
		reservations               map[string]*Reservation
		expectedBookerOk           bool
		expectedHistoryLen         int
		expectedBookerName         string
		expectedBookerPlacesAmount int
		wantErr                    error
	}{
		{name: "successful unbooking", capacity: 20, freePlaces: 10, reservations: map[string]*Reservation{"vlad": {Name: "vlad", PlaceAmount: 10}}, expectedFreePlaces: 20, expectedBookerOk: false, unbookerName: "vlad", expectedHistoryLen: 1, expectedBookerName: "", expectedBookerPlacesAmount: 0, wantErr: nil},
		{name: "validate empty name", capacity: 20, freePlaces: 10, reservations: map[string]*Reservation{"vlad": {Name: "vlad", PlaceAmount: 10}}, expectedFreePlaces: 10, expectedBookerOk: true, unbookerName: "", expectedHistoryLen: 0, expectedBookerName: "vlad", expectedBookerPlacesAmount: 10, wantErr: errEmptyName},
		{name: "validate incorrect name", capacity: 20, freePlaces: 10, reservations: map[string]*Reservation{"vlad": {Name: "vlad", PlaceAmount: 10}}, expectedFreePlaces: 10, expectedBookerOk: true, unbookerName: "John", expectedHistoryLen: 0, expectedBookerName: "vlad", expectedBookerPlacesAmount: 10, wantErr: errNonexistentClient},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			system := EventReservationSystem{
				Capacity:     tt.capacity,
				FreePlace:    tt.freePlaces,
				Reservations: tt.reservations,
			}

			err := system.CancellationOfBooking(tt.unbookerName)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if system.FreePlace != tt.expectedFreePlaces {
				t.Errorf("expected %d free places, actual %d", tt.expectedFreePlaces, system.FreePlace)
			}
			if tt.expectedBookerName == "" {
				_, ok := system.Reservations[tt.unbookerName]

				if ok != tt.expectedBookerOk {
					t.Errorf("expected availability %v, got %v", tt.expectedBookerOk, ok)
				}
			} else {
				reservation, ok := system.Reservations[tt.expectedBookerName]

				if reservation == nil {
					t.Fatalf("expected reservation for %v", tt.expectedBookerName)
				}

				if ok != tt.expectedBookerOk {
					t.Fatalf("expected availability %v, got %v", tt.expectedBookerOk, ok)
				}

				if reservation.PlaceAmount != tt.expectedBookerPlacesAmount {
					t.Errorf("expected booker place amount %d, actual %d", tt.expectedBookerPlacesAmount, reservation.PlaceAmount)
				}
			}

			if len(system.History) != tt.expectedHistoryLen {
				t.Errorf("expected history length %d, actual %d", tt.expectedHistoryLen, len(system.History))
			}

		})
	}
}

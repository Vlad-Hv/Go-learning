package reservation

import (
	"errors"
	"testing"
)

func TestBooking(t *testing.T) {
	tests := []struct {
		name               string
		capacity           int
		freePlace          int
		reservations       map[string]*Reservation
		bookerName         string
		placeAmount        int
		wantFreePlace      int
		wantReservedPlaces int
		wantOk             bool
		wantHistoryLen     int
		wantErr            error
	}{
		{
			name:               "successful",
			capacity:           20,
			freePlace:          20,
			reservations:       make(map[string]*Reservation),
			bookerName:         "vlad",
			placeAmount:        10,
			wantFreePlace:      10,
			wantReservedPlaces: 10,
			wantOk:             true,
			wantHistoryLen:     1,
			wantErr:            nil,
		},
		{
			name:               "empty reserver name",
			capacity:           20,
			freePlace:          20,
			reservations:       make(map[string]*Reservation),
			bookerName:         "",
			placeAmount:        10,
			wantFreePlace:      20,
			wantReservedPlaces: 0,
			wantOk:             false,
			wantHistoryLen:     0,
			wantErr:            errNameEmpty},
		{
			name:     "reserve zero places",
			capacity: 20, freePlace: 20,
			reservations:       make(map[string]*Reservation),
			bookerName:         "vlad",
			placeAmount:        0,
			wantFreePlace:      20,
			wantReservedPlaces: 0,
			wantOk:             false,
			wantHistoryLen:     0,
			wantErr:            errInvalidAmount},
		{

			name:               "reserve less than zero places",
			capacity:           20,
			freePlace:          20,
			reservations:       make(map[string]*Reservation),
			bookerName:         "vlad",
			placeAmount:        -1,
			wantFreePlace:      20,
			wantReservedPlaces: 0,
			wantOk:             false,
			wantHistoryLen:     0,
			wantErr:            errInvalidAmount},
		{
			name:               "reserve more than available",
			capacity:           20,
			freePlace:          20,
			reservations:       make(map[string]*Reservation),
			bookerName:         "vlad",
			placeAmount:        30,
			wantFreePlace:      20,
			wantReservedPlaces: 0,
			wantOk:             false,
			wantHistoryLen:     0,
			wantErr:            errNotEnoughFreeSpace},
		{
			name:     "reserve full place",
			capacity: 20, freePlace: 20,
			reservations:       make(map[string]*Reservation),
			bookerName:         "vlad",
			placeAmount:        20,
			wantFreePlace:      0,
			wantReservedPlaces: 20,
			wantOk:             true,
			wantHistoryLen:     1,
			wantErr:            nil},
		{
			name:               "add places to existent booker",
			capacity:           20,
			freePlace:          10,
			reservations:       map[string]*Reservation{"vlad": {Name: "vlad", PlaceAmount: 10}},
			bookerName:         "vlad",
			placeAmount:        5,
			wantFreePlace:      5,
			wantReservedPlaces: 15,
			wantOk:             true,
			wantHistoryLen:     1,
			wantErr:            nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			system := System{Capacity: tt.capacity, FreePlace: tt.freePlace, Reservations: tt.reservations}

			err := system.Book(tt.bookerName, tt.placeAmount)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if system.FreePlace != tt.wantFreePlace {
				t.Errorf("expected %d free places, actual %d", tt.wantFreePlace, system.FreePlace)
			}

			reservation, ok := system.Reservations[tt.bookerName]

			if ok != tt.wantOk {
				t.Fatalf("expected availability of reservation as %v, got %v", tt.wantOk, ok)
			}

			if tt.wantOk {
				if reservation.PlaceAmount != tt.wantReservedPlaces {
					t.Errorf("expected reserved place %d, actual %d", tt.wantReservedPlaces, reservation.PlaceAmount)
				}

			}

			if len(system.History) != tt.wantHistoryLen {
				t.Errorf("expected history length %d, actual %d", tt.wantHistoryLen, len(system.History))
			}

		})
	}
}

func TestCancel(t *testing.T) {
	tests := []struct {
		name                   string
		capacity               int
		freePlaces             int
		wantFreePlaces         int
		unbookerName           string
		reservations           map[string]*Reservation
		wantOk                 bool
		wantHistoryLen         int
		wantName               string
		wantBookerPlacesAmount int
		wantErr                error
	}{
		{
			name:                   "successful",
			capacity:               20,
			freePlaces:             10,
			reservations:           map[string]*Reservation{"vlad": {Name: "vlad", PlaceAmount: 10}},
			wantFreePlaces:         20,
			wantOk:                 false,
			unbookerName:           "vlad",
			wantHistoryLen:         1,
			wantName:               "",
			wantBookerPlacesAmount: 0,
			wantErr:                nil},
		{
			name:                   "empty name",
			capacity:               20,
			freePlaces:             10,
			reservations:           map[string]*Reservation{"vlad": {Name: "vlad", PlaceAmount: 10}},
			wantFreePlaces:         10,
			wantOk:                 true,
			unbookerName:           "",
			wantHistoryLen:         0,
			wantName:               "vlad",
			wantBookerPlacesAmount: 10,
			wantErr:                errNameEmpty},
		{
			name:                   "incorrect name",
			capacity:               20,
			freePlaces:             10,
			reservations:           map[string]*Reservation{"vlad": {Name: "vlad", PlaceAmount: 10}},
			wantFreePlaces:         10,
			wantOk:                 true,
			unbookerName:           "John",
			wantHistoryLen:         0,
			wantName:               "vlad",
			wantBookerPlacesAmount: 10,
			wantErr:                errNonexistentClient},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			system := System{
				Capacity:     tt.capacity,
				FreePlace:    tt.freePlaces,
				Reservations: tt.reservations,
			}

			err := system.Cancel(tt.unbookerName)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if system.FreePlace != tt.wantFreePlaces {
				t.Errorf("expected %d free places, actual %d", tt.wantFreePlaces, system.FreePlace)
			}
			if tt.wantName == "" {
				_, ok := system.Reservations[tt.unbookerName]

				if ok != tt.wantOk {
					t.Errorf("expected availability %v, got %v", tt.wantOk, ok)
				}
			} else {
				reservation, ok := system.Reservations[tt.wantName]

				if reservation == nil {
					t.Fatalf("expected reservation for %v", tt.wantName)
				}

				if ok != tt.wantOk {
					t.Fatalf("expected availability %v, got %v", tt.wantOk, ok)
				}

				if reservation.PlaceAmount != tt.wantBookerPlacesAmount {
					t.Errorf("expected booker place amount %d, actual %d", tt.wantBookerPlacesAmount, reservation.PlaceAmount)
				}
			}

			if len(system.History) != tt.wantHistoryLen {
				t.Errorf("expected history length %d, actual %d", tt.wantHistoryLen, len(system.History))
			}

		})
	}
}

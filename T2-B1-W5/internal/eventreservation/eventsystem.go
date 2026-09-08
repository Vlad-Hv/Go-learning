package event

type EventReservationSystem struct {
	Capacity     int
	FreePlace    int
	Reservations map[string]*Reservation
	History      []string
}

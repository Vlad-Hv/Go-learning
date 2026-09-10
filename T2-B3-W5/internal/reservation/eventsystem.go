package reservation

type System struct {
	Capacity     int
	FreePlace    int
	Reservations map[string]*Reservation
	History      []string
}

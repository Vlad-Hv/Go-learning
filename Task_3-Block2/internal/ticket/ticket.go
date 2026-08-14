package ticket

import (
	"github.com/google/uuid"
)

type Ticket struct {
	ID         uuid.UUID
	OwnerName  string
	VisitsLeft int
	Status     string
}

func CreateTicket(ownerName string, option int, id uuid.UUID) Ticket {
	optionsMap := CreateTicketOptions()
	return Ticket{
		ID:         id,
		OwnerName:  ownerName,
		VisitsLeft: optionsMap[option],
		Status:     "active",
	}
}

func (timeTicket *Ticket) UseOneVisit() {
	timeTicket.VisitsLeft--
}

func (timeTicket *Ticket) FreezeTicket() {
	timeTicket.Status = "frozen"
}

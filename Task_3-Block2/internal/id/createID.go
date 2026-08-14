package id

import (
	"task3/internal/ticket"

	"github.com/google/uuid"
)

func createID() uuid.UUID {
	ID := uuid.New()
	return ID
}

func CreateOriginalID(ticketsStorage []ticket.Ticket) uuid.UUID {
	var ID uuid.UUID
	const restartIndex int = -1
	ID = createID()
	for i := 0; i < len(ticketsStorage); i++ {
		if (ticketsStorage)[i].ID == ID {
			ID = createID()
			i = restartIndex
		}
	}

	return ID
}

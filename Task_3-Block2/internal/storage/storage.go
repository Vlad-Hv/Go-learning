package storage

import (
	"task3/internal/ticket"
)

func CreateStorage() []ticket.Ticket {
	var ticketStorage []ticket.Ticket
	return ticketStorage
}

func AddTicket(timeTicket ticket.Ticket, storage *[]ticket.Ticket) {

	*storage = append(*storage, timeTicket)
}

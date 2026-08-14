package ui

import (
	"fmt"
	"task3/internal/ticket"
)

func PrintAllTickets(tickets []ticket.Ticket) {
	for i := 0; i < len(tickets); i++ {
		fmt.Println("\nMembership", i+1)
		fmt.Println("ID:", tickets[i].ID)
		fmt.Println("Name:", tickets[i].OwnerName)
		fmt.Println("Visits Left:", tickets[i].VisitsLeft)
		fmt.Println("Status:", tickets[i].Status)
	}
}

package ui

import (
	"fmt"
	"task3/internal/ticket"
)

func GetName() string {
	var ownerName string
	askName()
	fmt.Scanln(&ownerName)
	return ownerName
}

func askName() {
	fmt.Println("Enter your name:")
}

func askTicketOption() int {
	ticketOption := ticket.CreateTicketOptions()
	fmt.Println("The options to buy the ticket:")
	for option, offer := range ticketOption {
		fmt.Println(option, ":", offer, "visits")
		fmt.Println("Choose the order:")
	}
	return len(ticketOption)
}

func GetTicketOption() (int, int, error) {
	var option int
	lenthOptions := askTicketOption()
	_, err := fmt.Scanln(&option)
	return option, lenthOptions, err
}

func GetChosenTicket() (int, error) {
	var chosenTicketNumber int
	fmt.Println("Enter number of your ticket:")
	_, err := fmt.Scanln(&chosenTicketNumber)
	return chosenTicketNumber, err
}

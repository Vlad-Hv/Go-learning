package main

import (
	// "github.com/google/uuid"
	"fmt"
	"task3/internal/id"
	"task3/internal/storage"
	"task3/internal/ticket"
	"task3/internal/ui"
)

func main() {
	ticketStorage := storage.CreateStorage()
	for {
		option, err := ui.GetMenuOption()
		err = ui.ValidateUi(option, err)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if option == 5 {
			fmt.Println("You left the menu")
			break
		}

		switch option {
		case 1:
			ownerName := ui.GetName()
			err = ui.ValidateName(ownerName)

			if err != nil {
				fmt.Println(err)
				continue
			}
			option, lenthOptions, err := ui.GetTicketOption()

			err = ui.ValidateTicketOption(option, lenthOptions, err)

			if err != nil {
				fmt.Println(err)
				continue
			}

			ID := id.CreateOriginalID(ticketStorage)

			timeTicket := ticket.CreateTicket(ownerName, option, ID)
			storage.AddTicket(timeTicket, &ticketStorage)

		case 2:
			ui.PrintAllTickets(ticketStorage)

		case 3:
			indexToUserMap := ticket.CreateIndexToUser(&ticketStorage)
			//find by option, validate, make map, make the method to use one visit
			chosenOption, err := ui.GetChosenTicket()
			err = ui.ValidateUseOneTicket(chosenOption, len(ticketStorage), err)
			if err != nil {
				fmt.Println(err)
				continue
			}
			user := indexToUserMap[chosenOption]

			err = ticket.ValidateUserVisit(*user)
			if err != nil {
				fmt.Println(err)
				continue
			}

			err = ticket.ValidatePosibilityToUseTicket(*user)
			if err != nil {
				fmt.Println(err)
				continue
			}

			user.UseOneVisit()

		case 4:
			indexToUserMap := ticket.CreateIndexToUser(&ticketStorage)
			chosenOption, err := ui.GetChosenTicket()
			err = ui.ValidateTicketOption(chosenOption, len(ticketStorage), err)
			if err != nil {
				fmt.Println(err)
				continue
			}
			user := indexToUserMap[chosenOption]
			err = ticket.ValidateTicketStatus(*user)
			if err != nil {
				fmt.Println(err)
				continue
			}

			user.FreezeTicket()
		}

		//switch case
	}
}

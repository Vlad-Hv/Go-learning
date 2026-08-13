package main

import (
	"fmt"
	"task3/internal/delivery"
	"task3/internal/id"
	"task3/internal/storage"
	"task3/internal/ui"
	"task3/internal/validation"
	//"github.com/google/uuid"
)

func main() {
	deliveryStorage := storage.CreateStorage()
	numToID := delivery.MakeNumToID(deliveryStorage)
	idToDelivery := delivery.MakeIdToDelivery(&deliveryStorage)

	for {
		option, err := ui.GetMenuOption()
		err = validation.ValidateMenu(option, err)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if option == 5 {
			fmt.Println("you left delivery menu")
			break
		}

		switch option {
		case 1:
			deliveryName, deliveryWeight, err := ui.GetDeliveryInfo()
			err = validation.ValidateCreatingDelivery(deliveryName, deliveryWeight, err)
			if err != nil {
				fmt.Println(err)
				continue
			}

			ID := id.GenerateUniqueID(deliveryStorage)
			order := delivery.CreateDelivery(deliveryName, deliveryWeight, ID)
			storage.StorageDeliveries(&deliveryStorage, order)

			numToID = delivery.MakeNumToID(deliveryStorage)
			idToDelivery = delivery.MakeIdToDelivery(&deliveryStorage)
			fmt.Println("Delivery done successfully")

		case 2:
			err := validation.ValidateDeliveries(deliveryStorage)
			if err != nil {
				fmt.Println(err)
				continue
			}

			ui.ShowAllDeliveries(deliveryStorage)

		case 3:
			option, err := ui.GetID()
			err = validation.ValidateID(option, err, deliveryStorage)

			if err != nil {
				fmt.Println(err)
				continue
			}

			delivery := idToDelivery[numToID[option]]
			ui.PrintDeliveryByID(delivery)

		case 4:
			option, err := ui.GetID()
			err = validation.ValidateID(option, err, deliveryStorage)

			if err != nil {
				fmt.Println(err)
				continue
			}

			order := idToDelivery[numToID[option]]
			err = validation.ValidatePossibilityMarkShipped(order)

			if err != nil {
				fmt.Println(err)
				continue
			}

			order.MarkAsShipped()

		}

	}
}

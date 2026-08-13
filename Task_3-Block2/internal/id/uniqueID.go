package id

import (
	"task3/internal/delivery"
	"task3/internal/validation"

	"github.com/google/uuid"
)

func GenerateUniqueID(deliveryStorage []delivery.Delivery) uuid.UUID {
	var ID uuid.UUID
	for {
		ID = generateID()
		err := validation.ValidateUniqueID(ID, deliveryStorage)

		if err != nil {
			continue
		} else {
			break
		}
	}
	return ID
}

package delivery

import (
	"github.com/google/uuid"
)

func MakeNumToID(deliveries []Delivery) map[int]uuid.UUID {
	numToID := make(map[int]uuid.UUID)

	for i := 0; i < len(deliveries); i++ {
		numToID[i] = deliveries[i].ID
	}
	return numToID
}

func MakeIdToDelivery(deliveries *[]Delivery) map[uuid.UUID]*Delivery {
	idToDelivery := make(map[uuid.UUID]*Delivery)

	for index := range *deliveries {
		idToDelivery[(*deliveries)[index].ID] = &(*deliveries)[index]
	}

	return idToDelivery
}

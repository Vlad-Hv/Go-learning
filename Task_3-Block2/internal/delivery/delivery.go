package delivery

import (
	"github.com/google/uuid"
)

type Delivery struct {
	ID        uuid.UUID
	Recipient string
	Weight    float64
	Status    string
}

func CreateDelivery(recipient string, weight float64, id uuid.UUID) Delivery {
	return Delivery{
		ID:        id,
		Recipient: recipient,
		Weight:    weight,
		Status:    "pending",
	}
}

func (delivery *Delivery) MarkAsShipped() {
	delivery.Status = "shipped"
}

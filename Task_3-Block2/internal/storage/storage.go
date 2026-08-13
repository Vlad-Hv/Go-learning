package storage

import (
	"task3/internal/delivery"
)

func CreateStorage() []delivery.Delivery {
	var deliveryStorage []delivery.Delivery
	return deliveryStorage
}

func StorageDeliveries(deliveryStorage *[]delivery.Delivery, order delivery.Delivery) {
	*deliveryStorage = append(*deliveryStorage, order)
}

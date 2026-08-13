package validation

import (
	"errors"
	"task3/internal/delivery"

	"github.com/google/uuid"
)

func ValidateMenu(option int, err error) error {
	if err != nil {
		return errors.New("incorrect option type")
	}

	if option < 1 || option > 5 {
		return errors.New("incorrect option")
	}
	return nil
}

func ValidateCreatingDelivery(name string, weight float64, err error) error {
	if err != nil {
		return errors.New("incorrect weight type")
	}

	if name == "" {
		return errors.New("delivery name mustnot be empty")
	}

	if weight <= 0 {
		return errors.New("weight must be more than zero")
	}
	return nil
}

func ValidateDeliveries(deliveries []delivery.Delivery) error {
	if len(deliveries) == 0 {
		return errors.New("deliveries list is empty")
	}
	return nil
}

func ValidateUniqueID(id uuid.UUID, orders []delivery.Delivery) error {
	for _, order := range orders {
		if id == order.ID {
			return errors.New("id already using")
		}
	}

	return nil
}

func ValidateID(option int, err error, deliveries []delivery.Delivery) error {
	if err != nil {
		return errors.New("incorrect number type")
	}

	if option < 0 || option >= len(deliveries) {
		return errors.New("invalid delivery number")
	}

	return nil
}

func ValidatePossibilityMarkShipped(order *delivery.Delivery) error {
	if order.Status == "shipped" {
		return errors.New("delivery is already shipped")
	}
	return nil
}

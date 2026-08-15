package expence

import "github.com/google/uuid"

type Expence struct {
	ID       uuid.UUID
	Category string
	Amount   int
}

func CreateExpence(category string, amount int) Expence {
	id := uuid.New()

	return Expence{
		ID:       id,
		Category: category,
		Amount:   amount,
	}
}

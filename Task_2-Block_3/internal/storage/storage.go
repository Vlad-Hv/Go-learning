package storage

import (
	"task2/internal/expence"
)

func CreateStorage() []expence.Expence {
	var storage []expence.Expence
	return storage
}

func PrintTotal(storage []expence.Expence) int {
	var total int
	for _, expence := range storage {
		total += expence.Amount
	}
	return total
}

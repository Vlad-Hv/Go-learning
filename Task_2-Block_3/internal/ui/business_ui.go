package ui

import (
	"fmt"
	"task2/internal/expence"
)

func ShowExpences(storage []expence.Expence) {
	for index, expence := range storage {
		fmt.Println("\nExpence", index+1)
		fmt.Println(expence)
	}
}

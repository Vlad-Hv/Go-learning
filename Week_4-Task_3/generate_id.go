package main

import "github.com/google/uuid"

func generateOrderID() uuid.UUID {
	id := uuid.New()
	return id
}

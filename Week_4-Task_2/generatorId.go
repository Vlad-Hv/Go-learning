package main

import "github.com/google/uuid"

func generatorID() uuid.UUID {
	id := uuid.New()
	return id
}

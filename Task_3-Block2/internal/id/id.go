package id

import (
	"github.com/google/uuid"
)

func generateID() uuid.UUID {
	id := uuid.New()
	return id
}

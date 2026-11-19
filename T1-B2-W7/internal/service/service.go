package service

import (
	order "orderHandle/internal/service/orders"
	"orderHandle/internal/storage"
)

type Storage interface {
	Save(storage.Result) error
	Statistick(string, error) storage.Statistick
}

type Handler interface {
	Check(*order.Order, <-chan struct{}) string
}
type Service struct {
	Storage Storage
	Logic   Handler
}

func NewService(storage Storage, logic Handler) *Service {
	return &Service{
		Storage: storage,
		Logic:   logic,
	}
}

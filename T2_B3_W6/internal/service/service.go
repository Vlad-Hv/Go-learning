package service

import (
	"errors"
	"product/internal/storage"
)

var (
	errAmountNotEnought = errors.New("actual amount not enough")
	errAmountInvalid    = errors.New("invalid amount")
)

type Service struct {
	storage storage.Storage
}

func New(storage storage.Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s Service) ReserveProducts(reserveAmount, ID int) (int, error) {
	amount, err := s.storage.CheckProduct(ID)
	if err != nil {
		return 0, err
	}

	if amount < reserveAmount {
		return 0, errAmountNotEnought
	}

	if reserveAmount <= 0 {
		return 0, errAmountInvalid
	}

	actual := amount - reserveAmount
	if err = s.storage.Reserve(ID, actual); err != nil {
		return 0, err
	}

	return actual, nil
}

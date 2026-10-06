package internal

import (
	"errors"
)

type User struct {
	IsSubscribeActive bool
}

type Storage map[int]User

func (s Storage) Find(ID int) (bool, error) {
	if _, ok := s[ID]; !ok {
		return false, errors.New("undefined user")
	}
	return s[ID].IsSubscribeActive, nil
}

func CreateStorage() Storage {
	return Storage{
		10: User{true},
		23: User{false},
		11: User{true},
	}
}

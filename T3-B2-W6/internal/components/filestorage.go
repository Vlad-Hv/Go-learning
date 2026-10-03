package components

import (
	"errors"
	"fmt"
)

type Storage struct {
	BaseInfo
	Min    int
	Actual int
}

func (s *Storage) Send() {
	var err error

	if s.IsCheking {
		if s.Actual < s.Min {
			s.IsWorking = "malfunctioning"
			err = errors.New("storage is not working, actual places less than minimum")
		} else {
			s.IsWorking = "working well"
		}
	} else {
		s.IsWorking = "not checking"
	}

	fmt.Println("Name:", s.Name)
	fmt.Println("Condition", s.IsWorking)
	if err != nil {
		fmt.Println(err)
	}
}

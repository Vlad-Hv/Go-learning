package components

import (
	"errors"
	"fmt"
)

type Queue struct {
	BaseInfo
	Limit        int
	WaitMessages int
}

func (q *Queue) Send() {
	var err error

	if q.IsCheking {
		if q.WaitMessages > q.Limit {
			q.IsWorking = "malfunctioning"
			err = errors.New("wait messages more than queue limit")
		} else {
			q.IsWorking = "working well"
		}
	} else {
		q.IsWorking = "not cheking"
	}

	fmt.Println("Name:", q.Name)
	fmt.Println("Condition", q.IsWorking)
	if err != nil {
		fmt.Println(err)
	}
}

package components

import (
	"errors"
	"fmt"
)

type DataBase struct {
	BaseInfo
	IsConect bool
}

func (d *DataBase) Send() {
	var err error

	if d.IsCheking {
		if !d.IsConect {
			d.IsWorking = "malfunctioning"
			err = errors.New("there is no conection")
		} else {
			d.IsWorking = "working well"
		}
	} else {
		d.IsWorking = "not cheking"
	}
	fmt.Println("Name:", d.Name)
	fmt.Println("Condition", d.IsWorking)
	if err != nil {
		fmt.Println(err)
	}
}

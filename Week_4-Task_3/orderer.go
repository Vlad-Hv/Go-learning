package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Order struct {
	Username    string
	DeviceName  string
	ProblemInfo string
	Priority    string
	IsClosed    bool
	ID          uuid.UUID
}

func createOrder(name, deviceName, problem string, priority string, id uuid.UUID) Order {

	return Order{
		Username:    name,
		DeviceName:  deviceName,
		ProblemInfo: problem,
		Priority:    priority,
		IsClosed:    false,
		ID:          id,
	}
}

func (order *Order) closeOrder() {
	order.IsClosed = true
}

func (order Order) checkInfo() {
	var isClosed string
	if !order.IsClosed {
		isClosed = "Open"
	} else {
		isClosed = "Closed"
	}
	fmt.Println("ID: ", order.ID)
	fmt.Println("Name", order.Username)
	fmt.Println("Device name", order.DeviceName)
	fmt.Println("Problem", order.ProblemInfo)
	fmt.Println("Priority", order.Priority)
	fmt.Println("Status", isClosed)
}

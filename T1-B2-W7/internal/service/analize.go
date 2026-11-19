package service

import (
	order "orderHandle/internal/service/orders"
	"time"
)

type Analizier struct{}

func (a Analizier) Check(order *order.Order, done <-chan struct{}) string {
	select {
	case <-time.After(time.Millisecond * time.Duration(order.HandlerTime)):
		order.WasHandled = true
		if order.Sum > 0 {
			return "approved"
		} else {
			return "cancled"
		}

	case <-done:
		return "timeout"
	}
}

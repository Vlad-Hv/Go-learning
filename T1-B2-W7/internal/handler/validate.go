package handler

//провалидировать инварианты

import (
	"errors"
	order "orderHandle/internal/service/orders"
	"strings"
)

var (
	errPlentyOfWorkers = errors.New("Too much workers amount")
	errOrderHandled    = errors.New("order already handled")
	errOrdersEmpty     = errors.New("orders are empty")
	errWorkerAmount    = errors.New("invalid workers amount")
	errTimeInvalid     = errors.New("invalid work time amount")
	errStatusTimeOut   = errors.New("timeout")
)

func validate(workerAmount int, orders []order.Order, workTime int) error {
	if workerAmount > len(orders) {
		return errPlentyOfWorkers
	}

	if workerAmount <= 0 {
		return errWorkerAmount
	}

	if len(orders) == 0 {
		return errOrdersEmpty
	}

	if workTime <= 0 {
		return errTimeInvalid
	}

	return nil
}

func validateOrder(order order.Order) error {
	if order.WasHandled {
		return errOrderHandled
	}

	return nil
}

func validateStatus(status string) error {
	if strings.EqualFold(status, "timeout") {
		return errStatusTimeOut
	}
	return nil
}

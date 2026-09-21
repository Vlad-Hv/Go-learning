package main

import (
	"delivery/internal/deliveries"
	"delivery/internal/history"
	"delivery/internal/storage"
	"fmt"
)

func main() {
	mailStorage := storage.Create()
	history := history.Create()
	service := deliveries.Service{Price: 10}
	person := deliveries.Person{Price: 20}

	SuccessPersonDelivery(mailStorage, person, &history)
	SuccessServiceDelivery(mailStorage, service, &history)
	PersonBigDelivery(mailStorage, person, service, &history)
	ServiceDeliveryAgain(mailStorage, service, &history)
	IDincorrect(mailStorage, service, &history)

	for i := 10; i < 13; i++ {
		fmt.Println(*mailStorage[i])
	}

	for _, history := range history {
		fmt.Println(history)
	}
}

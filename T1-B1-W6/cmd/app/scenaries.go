package main

import (
	"delivery/internal/deliveries"
	"delivery/internal/history"
	"delivery/internal/mail"
	"delivery/internal/storage"
	"fmt"
)

func SuccessPersonDelivery(mailStorage map[int]*mail.Mail, person deliveries.Person, history *history.History) {
	mail := mail.Mail{Recipient: "vlad", Weight: 9, State: "undelivered"}
	err := storage.Add(mailStorage, 10, mail)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = deliveries.Delivery(person, mailStorage[10], history)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func SuccessServiceDelivery(mailStorage map[int]*mail.Mail, service deliveries.Service, history *history.History) {
	mail := mail.Mail{Recipient: "vlad", Weight: 90, State: "undelivered"}
	err := storage.Add(mailStorage, 11, mail)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = deliveries.Delivery(service, mailStorage[11], history)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func PersonBigDelivery(mailStorage map[int]*mail.Mail, person deliveries.Person, service deliveries.Service, history *history.History) {
	mail := mail.Mail{Recipient: "vlad", Weight: 50, State: "undelivered"}
	err := storage.Add(mailStorage, 12, mail)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = deliveries.Delivery(person, mailStorage[12], history)
	if err != nil {
		fmt.Println(err)
		err = deliveries.Delivery(service, mailStorage[12], history)
		if err != nil {
			fmt.Println(err)
			return
		}

		return
	}

}

func ServiceDeliveryAgain(mailStorage map[int]*mail.Mail, service deliveries.Service, history *history.History) {
	mail := mail.Mail{Recipient: "vlad", Weight: 90, State: "undelivered"}
	err := storage.Add(mailStorage, 11, mail)
	if err != nil {
		fmt.Println(err)
	}

	err = deliveries.Delivery(service, mailStorage[11], history)
	if err != nil {
		fmt.Println(err)
		return
	}

}

func IDincorrect(mailStorage map[int]*mail.Mail, service deliveries.Service, history *history.History) {
	mail := mail.Mail{Recipient: "vlad", Weight: 90, State: "undelivered"}
	err := storage.Add(mailStorage, 13, mail)
	if err != nil {
		fmt.Println(err)
	}

	err = deliveries.Delivery(service, mailStorage[14], history)
	if err != nil {
		fmt.Println(err)
		return
	}

}

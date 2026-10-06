package main

import (
	"fmt"
	"service/internal"
)

func main() {
	storage := internal.CreateStorage()
	service := internal.NewService(storage)
	decision, err := service.Deciding(10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decision)

	decision, err = service.Deciding(23)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decision)

	decision, err = service.Deciding(11)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decision)

	decision, err = service.Deciding(0)
	if err != nil {
		fmt.Println(err)
	}
}

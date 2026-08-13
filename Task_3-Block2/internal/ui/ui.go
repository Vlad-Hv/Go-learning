package ui

import (
	"fmt"
	"task3/internal/delivery"
)

func GetMenuOption() (int, error) {
	var option int
	printMenu()
	_, err := fmt.Scanln(&option)
	return option, err
}

func printMenu() {
	fmt.Println("---Menu---")
	fmt.Println("1. Create delivery\n2. Show all deliveries\n3. Find delivery by ID\n4. Mark delivery as shipped\n5. Exit")
}

func ShowAllDeliveries(deliveries []delivery.Delivery) {
	for index, delivery := range deliveries {
		fmt.Println("\nDelivery", index+1)
		fmt.Println("ID:", delivery.ID)
		fmt.Println("Recipient:", delivery.Recipient)
		fmt.Println("Weight:", delivery.Weight)
		fmt.Println("Status:", delivery.Status)
	}
}

func GetID() (int, error) {
	printGetID()
	var option int
	_, err := fmt.Scanln(&option)
	return option - 1, err
}

func printGetID() {
	fmt.Println("Enter please the delivery number:")
}

func PrintDeliveryByID(order *delivery.Delivery) {
	fmt.Println("ID:", order.ID)
	fmt.Println("Recipient:", order.Recipient)
	fmt.Println("Weight:", order.Weight)
	fmt.Println("Status:", order.Status)
}

package ui

import (
	"fmt"
)

func GetDeliveryInfo() (string, float64, error) {
	var deliveryName string
	var deliveryWeight float64

	printName()
	fmt.Scanln(&deliveryName)

	printWeight()
	_, err := fmt.Scanln(&deliveryWeight)
	return deliveryName, deliveryWeight, err
}

func printName() {
	fmt.Println("Enter recipient name:")
}

func printWeight() {
	fmt.Println("Enter delivery weight:")
}

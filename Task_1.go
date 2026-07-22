package main

import (
	"errors"
	"fmt"
)

type Device struct {
	Name    string
	Price   float64
	InStock bool
}

func main() {
	device := getMap()

	deviceId, price, err := askUserInfo(device)

	if err != nil {
		fmt.Println(err)
		return
	}

	deviceF := device[deviceId]
	deviceF.Price = price
	deviceF.InStock = false
	device[deviceId] = deviceF

	fmt.Println(device[deviceId])

}

func getMap() map[string]Device {
	device := map[string]Device{
		"BNA21": {
			Name:    "Lenovo",
			Price:   199.99,
			InStock: true,
		},

		"LAR93": {
			Name:    "MacBoook",
			Price:   1982.19,
			InStock: true,
		},

		"NLE68": {
			Name:    "Phone",
			Price:   2984.74,
			InStock: true,
		},
	}
	return device
}

func askUserInfo(devices map[string]Device) (string, float64, error) {
	var deviceId string
	var price float64

	fmt.Print("Enter device key: ")
	fmt.Scanln(&deviceId)

	_, ok := devices[deviceId]
	if !ok {
		return "", 0, errors.New("incorrect key")
	}

	fmt.Println("Enter the price: ")
	_, err := fmt.Scanln(&price)
	if err != nil {
		return "", 0, errors.New("invalid price")
	}

	return deviceId, price, nil
	/*this message to you kai. I thought, that I'm able no change map herre just like devices[deviseId].Price = price and later just return changed map, but how I understand, I have to change this data in the main func*/
}

package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Adress struct {
	City       string
	Street     string
	HomeNumber float64
}

type Recipient struct {
	Name        string
	PhoneNumber string
	Adress      Adress
}

type Parcel struct {
	ID        int
	Weight    float64
	Cost      float64
	Recipient Recipient
	IsPost    bool
}

func main() {
	var cost float64 = 1500.50
	name, phoneNumber, adress, err := getRecipientInfo()
	id := getId()

	if err != nil {
		fmt.Println(err)
		return
	}

	changedPhoneNumber := fmt.Sprintf("+%d", phoneNumber)
	weight, err := getParcelInfo()

	if err != nil {
		fmt.Println(err)
		return
	}

	if weight > 10 {
		cost += cost / 2
	}

	recipient := Recipient{
		Name:        name,
		PhoneNumber: changedPhoneNumber,
		Adress:      adress,
	}

	parcel := Parcel{
		ID:        id,
		Weight:    weight,
		Cost:      cost,
		Recipient: recipient,
		IsPost:    false,
	}

	printParcelInfo(parcel)
	parcel = sentParcel(parcel)
	printParcelInfo(parcel)
}

func getRecipientInfo() (string, int, Adress, error) {
	var city string
	var street string
	var homeNumber float64
	var name string
	var phoneNumber int
	fmt.Println("Hello, nice to meet you in our interet market!\nCould you enter your name ant phone number please: ")
	fmt.Scanln(&name)
	_, err := fmt.Scanln(&phoneNumber)

	if err != nil {
		return "", 0, Adress{}, errors.New("incorrect input")
	}

	fmt.Println("\nWell enter your adress(city, street, nome number) to get your recepie: ")
	fmt.Scanln(&city)
	fmt.Scanln(&street)
	_, err = fmt.Scanln(&homeNumber)

	if err != nil {
		return "", 0, Adress{}, errors.New("invalid input")
	}

	if name == "" || city == "" || street == "" {
		return "", 0, Adress{}, errors.New("your info mustnot be empty")
	}

	adress := Adress{
		City:       city,
		Street:     street,
		HomeNumber: homeNumber,
	}

	return name, phoneNumber, adress, nil
}

func getParcelInfo() (float64, error) {
	var weight float64
	fmt.Println("\nEnter parcel weight:")
	_, err := fmt.Scanln(&weight)

	if err != nil || weight <= 0 {
		return 0, errors.New("invalid weight")
	}

	return weight, nil
}

func getId() int {
	id := rand.Intn(1000000) + 1000000
	return id
}

func printParcelInfo(parcel Parcel) {
	var isPost string
	if !parcel.IsPost {
		isPost = "Not Delivered"
	} else {
		isPost = "Delivered"
	}
	fmt.Println("\nParcel ID: ", parcel.ID)
	fmt.Println("Weight:", parcel.Weight)
	fmt.Println("Cost:", parcel.Cost)
	fmt.Println("Name:", parcel.Recipient.Name)
	fmt.Println("Number:", parcel.Recipient.PhoneNumber)
	fmt.Println("City:", parcel.Recipient.Adress.City)
	fmt.Println("Street:", parcel.Recipient.Adress.Street)
	fmt.Println("Home number:", parcel.Recipient.Adress.HomeNumber)
	fmt.Println("Status:", isPost)
}

func sentParcel(parcel Parcel) Parcel {
	parcel.IsPost = true

	return parcel
}

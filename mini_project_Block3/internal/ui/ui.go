package ui

import (
	"fmt"
	"mini/internal/elevator"
)

func GetMenuOption() (int, error) {
	var option int
	printMenuOption()
	_, err := fmt.Scanln(&option)
	return option, err
}

func GetCurrentFloor() (int, error) {
	var floor int
	printCurrentFloor()
	_, err := fmt.Scanln(&floor)

	return floor, err
}

func ReportStatus(elevator elevator.Elevator) {
	fmt.Println("Current floor:", elevator.CurrentFloor)
	fmt.Println("Is door open:", elevator.DoorOpen)
	fmt.Println("Is elevator moving:", elevator.IsMoving)
}

func printCurrentFloor() {
	fmt.Println("Choose floor you want to go(from 1 to 5):")
}

func printMenuOption() {
	fmt.Println("---Menu---")
	fmt.Println("1. Choose floor\n2. Open door\n3. Close door\n4. Show status\n5. Exit")
	fmt.Print("Choose the option: ")
}

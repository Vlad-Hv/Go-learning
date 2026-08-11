package main

import (
	"fmt"
)

func nameRequest() string {
	var name string
	fmt.Println("Enter your name:")
	fmt.Scanln(&name)
	return name
}

func getDeviceName() string {
	var device string
	fmt.Println("Enter device name:")
	fmt.Scanln(&device)
	return device
}

func getProblemInfo() string {
	var problemInfo string
	fmt.Println("Enter problem info:")
	fmt.Scanln(&problemInfo)
	return problemInfo
}

func getPriorety() (int, error) {
	var priority int
	fmt.Println("Choose priority(1 - low, 2 - medium, 3 - high):")
	_, err := fmt.Scanln(&priority)
	return priority, err
}

func checkInfoAsk() (int, error) {
	var option int
	fmt.Println("Do you want to check info?(1 - Yes, 2 - no)")
	_, err := fmt.Scanln(&option)
	return option, err
}

func IsOrderClose() (int, error) {
	var option int
	fmt.Println("Do you want to close the order?(1 - Yes, 2 - no):")
	_, err := fmt.Scanln(&option)
	return option, err
}

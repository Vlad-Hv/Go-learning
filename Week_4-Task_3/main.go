package main

import "fmt"

func main() {
	var option int
	priorityMap := createPriorityMap()

	id := generateOrderID()
	name := nameRequest()
	deviceName := getDeviceName()
	problemInfo := getProblemInfo()
	priority, err := getPriorety()
	err = validating(name, deviceName, problemInfo, priority, err)

	if err != nil {
		fmt.Println(err)
		return
	}

	order := createOrder(name, deviceName, problemInfo, priorityMap[priority], id)

	option, err = IsOrderClose()
	err = validateUserOption(option, err)
	if err != nil {
		fmt.Println(err)
		return
	}
	if option == 1 {
		order.closeOrder()
		fmt.Println("order closed\n ")
	}

	option, err = checkInfoAsk()
	err = validateUserOption(option, err)
	if err != nil {
		fmt.Println(err)
		return
	}

	if option != 2 {
		order.checkInfo()
	}

	fmt.Println("Bye Bye")

}

func createPriorityMap() map[int]string {
	return map[int]string{
		1: "low",
		2: "medium",
		3: "high",
	}
}

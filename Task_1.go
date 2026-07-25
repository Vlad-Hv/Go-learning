package main

import (
	"errors"
	"fmt"
	"math/rand"
)

type Device struct {
	Name  string
	Model string
}

type Client struct {
	Name    string
	Age     int
	Device  Device
	Balance float64
}

type RepairOrder struct {
	ID         int
	ClientInfo Client
	RepairCost float64
	IsDone     bool
	Status     string
}

func main() {
	var completedOrders []RepairOrder
	devices := getDevices()
	var orders []RepairOrder
	var id int
	//var pointerToOrder *RepairOrder
	clientsInfo, err := getClientInfo(devices)

	if err != nil {
		fmt.Println(err)
		return
	}

	orders, err = createRepairOrders(devices, clientsInfo)

	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < len(orders); i++ {
		id, err = getID(orders)

		if err != nil {
			fmt.Println(err)
			return
		}

		//orderMap := createMap(orders)

		//order := orderMap[id] //вот тут ордер уже имеет скопированую структуру, а можно пройтись ФОРом и найти нужную нам структуру в ориге и даь уже на нее указатеь

		for i := 0; i < len(orders); i++ {
			if id == orders[i].ID {
				changeIsDone(&orders[i])
				err = changeUserBalance(&clientsInfo[i], orders[i])
				changeStatus(&orders[i])
				if err != nil {
					fmt.Println(err)
					return
				}

				changeCost(&orders[i])
				history(&completedOrders, orders[i])
			}
		}
		fmt.Println("Completed Succesfully")
	}
	fmt.Println(orders)
	fmt.Println(completedOrders)

}

func getDevices() []Device {
	devices := []Device{
		{
			Name:  "MacBook",
			Model: "M3 pro",
		},

		{
			Name:  "iPhone",
			Model: "17 pro",
		},

		{
			Name:  "Huawei",
			Model: "A13",
		},

		{
			Name:  "Poco",
			Model: "X3 pro",
		},
	}

	return devices
	/*for i := 0; i < len(devices); i++{
		var name string
		fmt.Println("Enter your name: ")
		fmt.Scanln(&name)
		client.Name = name
		client.RepairDevice = devices[i]
		clientList = append(clientList, client)
	}

	return clientList*/
}

func getClientInfo(devices []Device) ([]Client, error) {
	var clientInfo []Client

	for i := 0; i < len(devices); i++ {
		var name string
		var age int
		var balance float64

		fmt.Println("\nEnter your name, age and balance: ")
		_, err := fmt.Scanln(&name, &age, &balance)

		if err != nil {
			return nil, errors.New("incorrect input")
		}

		if balance <= 0 {
			return nil, errors.New("not enough money")
		}

		client := Client{
			Name:    name,
			Age:     age,
			Device:  devices[i],
			Balance: balance,
		}

		clientInfo = append(clientInfo, client)
	}
	return clientInfo, nil
}

func createRepairOrders(devices []Device, clientInfo []Client) ([]RepairOrder, error) {
	var orders []RepairOrder
	prices, err := getPrices(devices)

	if err != nil {
		return nil, fmt.Errorf("cannot create order, reason:%w", err)
	}

	for i := 0; i < len(clientInfo); i++ {
		ID := rand.Intn(500) + 100

		order := RepairOrder{
			ID:         ID,
			ClientInfo: clientInfo[i],
			RepairCost: prices[i],
			IsDone:     false,
			Status:     "Pending",
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func getPrices(devices []Device) ([]float64, error) {
	var prices []float64
	var price float64
	for i := 0; i < len(devices); i++ {
		fmt.Print("\nEnter repair price to this device ", devices[i], ": ")
		_, err := fmt.Scanln(&price)

		if err != nil {
			return nil, errors.New("incorrect price type")
		}

		prices = append(prices, price)
	}
	return prices, nil
}

func getID(orders []RepairOrder) (int, error) {
	var id int
	var checker int

	for i := 0; i < len(orders); i++ {
		if !orders[i].IsDone {
			fmt.Println("\n", orders[i])
		}
	}
	fmt.Println("\nChoose order ID:")
	_, err := fmt.Scanln(&id)

	if err != nil {
		return 0, errors.New("invalid type")
	}

	for i := 0; i < len(orders); i++ {
		if orders[i].ID == id {
			checker++
		}
	}

	if checker == 0 {
		return 0, errors.New("incorrect id")
	}

	return id, nil
}

/*func createMap(orders []RepairOrder) map[int]RepairOrder {
	orderMap := make(map[int]RepairOrder)

	for i, order := range orders {
		orderMap[order.ID] = orders[i]
	}

	return orderMap
}*/

//закончил на том, что хочу создать мапу где ключ - айди, значение - структура из слайса заказов, чтобы можно было обращаться к структуре через айди
//затем пункт 5 и тд
//твою налево, трабл с айди и указателями, (пункты 4,5)
//короче, как я понял все пошло по пизде именно в моменте создания мапы айди: нужная структура
// есть вариант поебаться в этой функции попытавшись передать тот же адресс в структуру, что и в ориг стрктуре или типо такого
//или просто после выбор айди в мейн передать поинтер именно на совпадающую структуру с той что копия

func changeIsDone(order *RepairOrder) {
	if order == nil {
		return
	}

	order.IsDone = true
}

func changeUserBalance(client *Client, order RepairOrder) error {
	client.Balance -= order.RepairCost

	if client.Balance < 0 {
		return errors.New("not enough money")
	}

	return nil
}
func changeStatus(order *RepairOrder) {
	if order.IsDone == true {
		order.Status = "Completed"
	}
}
func changeCost(order *RepairOrder) {
	if order == nil {
		return
	}
	order.RepairCost = 2000
	order.ClientInfo.Balance += 2000
}

func history(history *[]RepairOrder, order RepairOrder) {
	*history = append(*history, order)
}

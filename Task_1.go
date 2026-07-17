package main

import "fmt"

type Customer struct {
	Name  string
	Email string
}

type Order struct {
	ProductName string
	Price       float64
	Customer    Customer
}

func main() {
	name, email, prodName, price := getData()
	order := Order{
		ProductName: prodName,
		Price:       price,
		Customer: Customer{
			Name:  name,
			Email: email,
		},
	}
	printOrder(order)
}

func getData() (string, string, string, float64) {
	var name string
	var Email string
	var ProdName string
	var Price float64

	fmt.Println("Enter name and email:")
	fmt.Scanln(&name, &Email)
	fmt.Println("Enter product name and price:")

	fmt.Scanln(&ProdName, &Price)

	return name, Email, ProdName, Price
}

func printOrder(order Order) {
	fmt.Println(order.ProductName, order.Price, order.Customer.Name, order.Customer.Email)
}

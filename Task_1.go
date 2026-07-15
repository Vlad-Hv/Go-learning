package main

import "fmt"

type Product struct {
	Name        string
	Price       float64
	Amount      int
	IsAvailable bool
}

func main() {
	name, price, option := getInfo()
	product := Product{
		Name:        name,
		Price:       price,
		Amount:      1,
		IsAvailable: true,
	}

	fmt.Println(product)

	if option == 1 {
		product.Amount = 0
		product.IsAvailable = false
		fmt.Println("Product is not available")
	} else {
		fmt.Println("Product is available")
	}

}

func getInfo() (string, float64, int) {
	var name string
	var price float64
	var option int

	fmt.Println("Enter product name and price:")
	fmt.Scanln(&name, &price)
	fmt.Print("Do you want to buy smth y-1, n-2: ")
	fmt.Scanln(&option)

	return name, price, option
}

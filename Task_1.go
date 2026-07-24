package main

import (
	"fmt"
	"math/rand"
)

type Product struct {
	Name   string
	Price  float64
	Amount int
}

func main() {
	product := []Product{
		{
			Name:   "Carrot",
			Price:  679.12,
			Amount: 10,
		},
		{
			Name:   "Potato",
			Price:  1299.10,
			Amount: 3,
		},
		{
			Name:   "Milk",
			Price:  1300,
			Amount: 1,
		},
	}
	fmt.Println("Before:", product)
	for i := 0; i < len(product); i++ {
		changeProduct(&product[i])
	}

	fmt.Println("After:", product)

}

func changeProduct(product *Product) {
	price := rand.Intn(500) + 100
	product.Price = float64(price)
	product.Amount--
}

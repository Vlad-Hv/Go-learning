package main

import "fmt"

type Product struct {
	Name        string
	Price       float64
	Amount      int
	IsAvailable bool
}

func main() {
	product := createProduct()
	product.PrintInfo()
	fmt.Println("Tatal price:", product.TotalPrice())
	fmt.Println("Is it in stock:", product.IsInStock()) //Kai, is it good to write like this? or it''s better to initilize variable first and print ..., variableName
}

func createProduct() Product {
	return Product{
		Name:        "IPhone",
		Price:       1999.99,
		Amount:      10,
		IsAvailable: true,
	}
}

func (product Product) PrintInfo() {
	fmt.Println("====INFO====")
	fmt.Println("Name:", product.Name)
	fmt.Println("Price:", product.Price)
	fmt.Println("Amount:", product.Amount)
}

func (product Product) TotalPrice() float64 {
	return product.Price * float64(product.Amount)
}

func (product Product) IsInStock() bool {
	return product.Amount > 0 && product.IsAvailable == true
}

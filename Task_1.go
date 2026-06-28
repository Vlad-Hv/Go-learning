package main

import "fmt"

func main() {
	var price int = 199
	var discount int = 15

	floatPrice := float64(price)
	floatDiscount := float64(discount) //I created new variables just to use another type of type conversation

	result := floatPrice - ((floatPrice * floatDiscount) / 100)

	fmt.Println("Final price:", result)

}

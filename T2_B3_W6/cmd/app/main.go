package main

import (
	"fmt"
	"product/internal/service"
	"product/internal/storage"
)

func main() {
	productByID := storage.CreateMap()
	Service := service.New(&productByID)
	if actual, err := Service.ReserveProducts(10, 10); err != nil {
		fmt.Println("1:", err)
	} else {
		fmt.Println(productByID[10].Name, actual)
	}

	if actual, err := Service.ReserveProducts(20, 11); err != nil {
		fmt.Println("2:", err)
	} else {
		fmt.Println(productByID[11].Name, actual)
	}

	if actual, err := Service.ReserveProducts(15, 12); err != nil {
		fmt.Println("3:", err)
	} else {
		fmt.Println(productByID[12].Name, actual)
	}

}

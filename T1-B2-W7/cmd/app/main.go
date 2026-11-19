package main

import (
	"orderHandle/internal/handler"
	"orderHandle/internal/service"
	order "orderHandle/internal/service/orders"
	"orderHandle/internal/storage"
	"orderHandle/internal/ui"
)

func main() {
	resultStorage := storage.Create()
	analizier := service.Analizier{}
	service := service.NewService(&resultStorage, analizier)

	workerAmount := 3
	orders := order.Create()

	statistick := handler.Handle(orders, workerAmount, 100, service)

	ui.PrintStatistick(statistick, orders)
}

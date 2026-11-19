package ui

import (
	"fmt"
	order "orderHandle/internal/service/orders"
	"orderHandle/internal/storage"
)

func PrintStatistick(statistick storage.Statistick, orders []order.Order) {
	results := calculateResults(statistick)

	unHandled := len(orders) - (results["approved"] + results["cancled"])
	approved := results["approved"]
	cancled := results["cancled"]

	printInfo(approved, cancled, unHandled, statistick, statistick.StopReason)
}

func calculateResults(statistick storage.Statistick) map[string]int {
	final := make(map[string]int)

	for _, result := range statistick.Results {
		final[result.Status]++
	}

	return final
}

func printInfo(approved, canceld, unHandled int, statistick storage.Statistick, reason string) {
	fmt.Println("finish reason", reason)
	fmt.Println("Approved:", approved)
	fmt.Println("Cancled:", canceld)
	fmt.Println("Did not handled:", unHandled)

	for id, result := range statistick.Results {
		fmt.Println(id, ":", result.Status)
	}
}

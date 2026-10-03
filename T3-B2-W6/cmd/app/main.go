package main

import (
	"collector/internal/collector"
	"fmt"
)

func main() {
	conditions := collector.Create()

	for _, condition := range conditions {
		condition.Send()
		fmt.Println()
	}
}

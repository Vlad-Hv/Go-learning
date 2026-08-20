package ui

import (
	"fmt"
	"task3/internal/log"
)

func PrintAllLogs(logs []log.Log) {
	for mode, info := range logs {
		fmt.Println(mode+1, info)

	}
}

func PrintStatistic(stat map[string]int) {
	var total int
	for mode, Type := range stat {
		fmt.Println(mode, ":", Type)
		total += Type
	}

	fmt.Println("Total:", total)
}

func PrintOnlyErrors(logs []log.Log) {
	for _, log := range logs {
		fmt.Println(log.Mode, ":", log.Type)
	}
}

func PrintFrequentEvent(Type string, amount int) {
	fmt.Println(Type, ":", amount)
}

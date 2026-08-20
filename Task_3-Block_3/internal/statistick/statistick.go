package statistick

import (
	"task3/internal/log"
)

func CreateStatistic(logs []log.Log) map[string]int {
	statistick := make(map[string]int)

	for _, log := range logs {
		statistick[log.Mode] += 1
	}
	return statistick
}

func createFrequentEventStat(logs []log.Log) map[string]int {
	statistick := make(map[string]int)
	for _, log := range logs {
		statistick[log.Type] += 1
	}

	return statistick
}

func createMirrorFrequentEventStat(logs []log.Log) map[int]string {
	mainStat := make(map[int]string)
	stat := createFrequentEventStat(logs)

	for Type, amount := range stat {
		mainStat[amount] = Type
	}

	return mainStat
}

func CalculateFrequentEvent(logs []log.Log) (string, int) {
	var amountHelper int

	stat := createMirrorFrequentEventStat(logs)

	for amount := range stat {
		if amountHelper < amount {
			amountHelper = amount
		}
	}

	return stat[amountHelper], amountHelper
}

func GetErrors(logs []log.Log) []log.Log {
	var errorLog []log.Log
	for _, log := range logs {
		if log.Mode == "ERROR" {
			errorLog = append(errorLog, log)
		}
	}

	return errorLog
}

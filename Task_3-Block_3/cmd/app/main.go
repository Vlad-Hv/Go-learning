package main

import (
	"fmt"
	"task3/internal/log"
	"task3/internal/statistick"
	"task3/internal/ui"
)

func main() {
	logs := log.CreateLogs()
	statistic := statistick.CreateStatistic(logs)
	messege, amount := statistick.CalculateFrequentEvent(logs)
	for {
		option, err := ui.GetOption()
		err = ui.ValidateMenuOption(option, err)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if option == 5 {
			fmt.Println("you left the menu")
			break
		}

		switch option {
		case 1:
			ui.PrintAllLogs(logs)

		case 2:
			ui.PrintStatistic(statistic)

		case 3:
			errorLog := statistick.GetErrors(logs)
			ui.PrintOnlyErrors(errorLog)

		case 4:
			ui.PrintFrequentEvent(messege, amount)
		}

	}
}

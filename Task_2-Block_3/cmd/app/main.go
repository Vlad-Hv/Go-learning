package main

import (
	"fmt"
	"task2/internal/expence"
	"task2/internal/storage"
	"task2/internal/ui"
	"task2/internal/validate"
)

func main() {
	expenceStorage := storage.CreateStorage()

	for {
		option, err := ui.GetMenuOption()
		err = validate.ValidateMenuOption(option, err)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if option == 4 {
			break
		}

		switch option {

		case 1:
			category, amount, err := ui.GetExpenceInfo()
			err = validate.ValidateExInfo(category, amount, err)

			if err != nil {
				fmt.Println(err)
				continue
			}

			expenc := expence.CreateExpence(category, amount)
			expenceStorage = append(expenceStorage, expenc)

		case 2:
			ui.ShowExpences(expenceStorage)

		case 3:
			total := storage.PrintTotal(expenceStorage)
			fmt.Println("total:", total)
		}

	}
}

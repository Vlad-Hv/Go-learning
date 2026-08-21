package main

import (
	"fmt"
	"mini/internal/elevator"
	"mini/internal/ui"
	"mini/internal/validation"
)

func main() {
	elevator := elevator.CreateElevator()

	for {
		option, err := ui.GetMenuOption()

		err = validation.ValidateMenuOption(option, err)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if option == 5 {
			fmt.Println("You left the menu")
			break
		}
		//создать валмдвцию инвариантов и начать делать методы под функции меню

		switch option {
		case 1:
			floor, err := ui.GetCurrentFloor()

			err = validation.ValidateCurrentFloor(err)
			if err != nil {
				fmt.Println(err)
				continue
			}

			err = validation.CurrentFloor(floor, elevator)
			if err != nil {
				fmt.Println(err)
				continue
			}

			elevator.GoFloor(floor)
			elevator.ArriveFloor()
			fmt.Println("Elevator arrived a current floor succesfully")

		case 2:
			err = validation.OpenDoor(elevator)

			if err != nil {
				fmt.Println(err)
				continue
			}

			elevator.OpenDoor()

		case 3:
			err = validation.CloseDoor(elevator)

			if err != nil {
				fmt.Println(err)
				continue
			}

			elevator.CloseDoor()

		case 4:
			ui.ReportStatus(elevator)

		}
	}
}

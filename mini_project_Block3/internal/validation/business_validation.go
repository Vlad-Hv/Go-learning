package validation

import (
	"errors"
	"mini/internal/elevator"
)

func CurrentFloor(floor int, elevator elevator.Elevator) error {
	if floor < 1 || floor > 5 {
		return errors.New("invalid floor")
	}

	if elevator.CurrentFloor == floor {
		return errors.New("you are already on this floor")
	}

	if elevator.DoorOpen {
		return errors.New("you are not able to use elevator, while the doors are open")
	}

	return nil
}

func OpenDoor(elevator elevator.Elevator) error {
	if elevator.DoorOpen {
		return errors.New("door is already open")
	}

	if elevator.IsMoving {
		return errors.New("cannot open the door, while elevator is moving")
	}

	return nil
}

func CloseDoor(elevator elevator.Elevator) error {
	if !elevator.DoorOpen {
		return errors.New("door is already close")
	}

	return nil
}

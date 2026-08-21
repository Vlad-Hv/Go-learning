package elevator

type Elevator struct {
	CurrentFloor int
	TargetFloor  int
	DoorOpen     bool
	IsMoving     bool
}

func CreateElevator() Elevator {
	return Elevator{
		CurrentFloor: 1,
		DoorOpen:     false,
		IsMoving:     false,
	}
}

func (elevator *Elevator) GoFloor(floor int) {
	elevator.TargetFloor = floor
	elevator.IsMoving = true
}

func (elevator *Elevator) ArriveFloor() {
	elevator.CurrentFloor = elevator.TargetFloor
	elevator.IsMoving = false
}

func (elevator *Elevator) OpenDoor() {
	elevator.DoorOpen = true
}

func (elevator *Elevator) CloseDoor() {
	elevator.DoorOpen = false
}

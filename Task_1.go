package main

import "fmt"

type Character struct {
	Name       string
	Health     int
	Level      int
	GoldAmount int
	IsAlive    bool
}

func main() {
	var name string
	fmt.Println("Enter hero name: ")
	fmt.Scanln(&name)

	startCharacter := Character{
		Name:       name,
		Health:     100,
		GoldAmount: 0,
		Level:      1,
		IsAlive:    true,
	}

	actions := actionSlice()

	fmt.Println(startCharacter)

	for i := 0; (len(actions) - 1) > i; i++ {

		switch i {
		case 0:
			fmt.Println(actions[i])
			startCharacter.Health -= 35

		case 1:
			fmt.Println(actions[i])
			startCharacter.GoldAmount += 50

		case 2:
			fmt.Println(actions[i])
			startCharacter.Level += 1

		case 3:
			fmt.Println(actions[i])
			startCharacter.Health -= 70
		}

		if startCharacter.Health <= 0 {
			startCharacter.Health = 0
			startCharacter.IsAlive = false
			fmt.Print(actions[4], "\n ")
		}

	}

	fmt.Println(startCharacter)

}

func actionSlice() []string {
	actions := []string{"\nYour hero got 35 damage", "\nYour hero found 50 golds", "\nYour hero level up for 1 level", "\nYour hero got 70 damage", "\nYour hero died"}
	return actions
}

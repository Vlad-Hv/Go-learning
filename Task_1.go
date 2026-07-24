package main

import (
	"fmt"
)

type Player struct {
	Name      string
	Inventory []string
}

func main() {
	item := newItem()
	var pointer *Player
	player := Player{
		Name:      "Vlad",
		Inventory: []string{"Axe", "Sword"},
	}

	addItem(pointer, item)
	pointer = &player
	addItem(pointer, item)

	fmt.Println(*pointer) //Kai, I did it specially, for fun more, or to fix in my mind, this topick

}

func newItem() string {
	var item string
	fmt.Println("Enter item name: ")
	fmt.Scanln(&item)

	return item
}

func addItem(player *Player, item string) {
	if player == nil {
		fmt.Println("pointer mustnot be empty")
		return
	}

	player.Inventory = append(player.Inventory, item)
}

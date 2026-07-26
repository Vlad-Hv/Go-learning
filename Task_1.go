package main

import (
	"errors"
	"fmt"
)

type Player struct {
	ID        int
	Username  string
	Health    int
	Gold      int
	Inventory []string
}

func main() {
	var history []Player
	players := createPlayers()
	fmt.Println("Before:", players)
	playerMap := createPlayerMap(players)
	id := findID()
	player := findPlayer(playerMap, id)
	damage := findDamage()
	err := takeDamage(player, damage)

	if err != nil {
		fmt.Println(err)
		return
	}

	gold, item := findGoldAndItem()

	addReward(player, gold, item)
	addHistory(&history, *player)

	id = 000
	player = findPlayer(playerMap, id)
	addReward(player, gold, item)

	fmt.Println("After:", players, "\nHistory:", history)
}

func createPlayers() []Player {
	players := []Player{
		{
			ID:        101,
			Username:  "Vlad",
			Health:    100,
			Gold:      120,
			Inventory: []string{"Axe", "Hammer"},
		},

		{
			ID:        164,
			Username:  "Vadim",
			Health:    100,
			Gold:      115,
			Inventory: []string{"Gun", "Bullet"},
		},

		{
			ID:        189,
			Username:  "Dima",
			Health:    100,
			Gold:      1000,
			Inventory: []string{"Cabel"},
		},
	}

	return players
}

func createPlayerMap(players []Player) map[int]*Player {
	playerMap := make(map[int]*Player)
	for i := range players {
		playerMap[players[i].ID] = &players[i]
	}
	return playerMap
}

func findPlayer(playerMap map[int]*Player, ID int) *Player {
	_, ok := playerMap[ID]

	if !ok {
		return nil
	}

	return playerMap[ID]
}

func findID() int {
	var id int
	fmt.Println("Enter ID:")
	fmt.Scanln(&id)
	return id
}

func findDamage() int {
	var damage int
	fmt.Println("Enter damage:")
	fmt.Scanln(&damage)
	return damage
}

func findGoldAndItem() (int, string) {
	var gold int
	var item string

	fmt.Println("Enter award gold and item:")
	fmt.Scanln(&gold, &item)

	return gold, item
}

func takeDamage(player *Player, damage int) error {
	if player == nil {
		return errors.New("icorrect id")
	}

	if damage < 0 {
		return errors.New("damage mustnot be less than zero")
	}

	healthDEMO := player.Health - damage

	if healthDEMO < 0 {
		player.Health = 0
	} else {
		player.Health -= damage
	}

	return nil
}

func addReward(player *Player, gold int, item string) {
	if player == nil {
		return
	}
	player.Gold += gold
	player.Inventory = append(player.Inventory, item)
}

func addHistory(players *[]Player, player Player) {
	*players = append(*players, player)
}

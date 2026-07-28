package main

import "fmt"

type Character struct {
	Name    string
	Health  int
	Level   int
	Gold    int
	IsAlive bool
}

func main() {
	var damage int = 51
	var amount int = 29
	character := Character{
		Name:    "Vlad",
		Health:  100,
		Level:   5,
		Gold:    15,
		IsAlive: true,
	}

	character.TakeDamage(damage)
	character.AddGold(amount)
	character.LevelUp()
	IsNotAlive := character.IsDead()
	fmt.Println(character)
	fmt.Println("Is it thue, that character died:", IsNotAlive)

}

func (hero *Character) TakeDamage(damage int) {
	if damage < 0 {
		return
	}
	hero.Health -= damage
	if hero.Health <= 0 {
		hero.IsAlive = false
		hero.Health = 0
	}
	fmt.Println(*hero)
}

func (hero *Character) AddGold(amount int) {
	if amount < 0 {
		return
	}
	hero.Gold += amount
	fmt.Println(*hero)
}

func (hero *Character) LevelUp() {
	hero.Level++
	fmt.Println(*hero)
}

func (hero *Character) IsDead() bool {
	return hero.Health <= 0 && hero.IsAlive == false
}

package main

import (
	"fmt"
	"task2/internal/profile"
	"task2/internal/ui"
)

func main() {
	name := ui.GetName()
	level := ui.GetLevel()
	profile := profile.CreateProfile(name, level)
	fmt.Println("Player:")
	fmt.Println("Name:", profile.Name)
	fmt.Println("Level:", profile.Level)
}

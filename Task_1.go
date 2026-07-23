package main

import (
	"errors"
	"fmt"
)

func main() {
	var health int = 100
	var healthPointer *int = &health

	for i := 0; i < 2; i++ {

		err := changeHealth(healthPointer)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Health:", health)
	}

}

func changeHealth(health *int) error {
	var damage int

	fmt.Println("Enter damage count")
	_, err := fmt.Scanln(&damage)

	if err != nil {
		return errors.New("invlid type")
	}

	healthCheck := *health

	if (healthCheck - damage) < 0 {
		return errors.New("too much damage")
	}

	*health -= damage

	return nil
}

package main

import "fmt"

func main() {
	var username string = "Vlad"
	fmt.Println(username)

	changeName(&username)
	fmt.Println(username)
}

func changeName(username *string) {
	*username = "Dima"
}

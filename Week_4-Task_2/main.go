package main

func main() {
	username := getUsername()
	userAge := getUserAge()
	id := generatorID()

	user := getUser(username, userAge, id)
	user.PrintInfo()
}

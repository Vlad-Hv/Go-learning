package main

import (
	"errors"
	"fmt"
)

func main() {
	commandList := commandList()
	users := activeUsers()
	eventHandler := infoMap()
	userCounter := make(map[string]int)
	var command string
	var err error
	var history []string
	var username string
	for {
		users, history, command, username, err = checkerCommand(commandList, users, eventHandler, history, userCounter)

		if err != nil {
			fmt.Println(err)
			eventHandler["Rejected actions"] += 1
			continue
		}

		if len(history) > 5 {
			history = history[1:]
		}

		userCounter[username] += 1

		if command == "report" {
			break
		}

	}
}

func activeUsers() []string {
	var users []string
	return users
}

func getInfo(commands []string) (string, string, error) {
	var user string
	var command string
	var checker int
	fmt.Print("\nEnter username and command: ")
	fmt.Scanln(&user, &command)

	for i := 0; i < len(commands); i++ {
		if command == commands[i] {
			checker += 1
			break
		}
	}

	if checker == 0 {
		return "", "", errors.New("invalid command")
	}

	return user, command, nil
}

func checkerCommand(comList []string, users []string, eventCheck map[string]int, history []string, uC map[string]int) ([]string, []string, string, string, error) {
	username, command, err := getInfo(comList)
	var message string

	if username == "report" {
		command = username
	}

	if err != nil {
		return users, history, "", "", fmt.Errorf("error: %w", err)
	}

	switch command {
	case "login":
		users, history, err = userLogin(users, username, eventCheck, history)

		if err != nil {
			return users, history, "", "", err
		}

		return users, history, command, username, nil
	case "logout":
		users, history, err = userLogout(users, username, eventCheck, history)

		if err != nil {
			return users, history, "", "", err
		}

		return users, history, command, username, nil
	case "buy":
		message, history, err = userBuy(users, username, eventCheck, history)

		if err != nil {
			return users, history, "", "", err
		}
		fmt.Println(message)

	case "report":
		report(users, eventCheck, history, uC)
		return users, history, command, username, nil
	}
	return users, history, command, username, nil

}

func commandList() []string {
	commands := []string{"login", "logout", "buy", "report"}
	return commands
}

func userLogin(userList []string, userName string, eventCheck map[string]int, history []string) ([]string, []string, error) {
	if len(userList) != 0 {
		for _, user := range userList {
			if user == userName {
				eventCheck["Rejected actions"] += 1
				return userList, history, errors.New("duplicate login")
			}
		}
	}

	userList = append(userList, userName)
	eventCheck["Successful actions"] += 1

	message := fmt.Sprintf("%q logget in succesfully", userName)
	history = append(history, message)
	fmt.Println("user loged successfully!")
	return userList, history, nil
}

func userLogout(userList []string, username string, event map[string]int, history []string) ([]string, []string, error) {
	if len(userList) == 0 {
		event["Rejected actions"] += 1
		return userList, history, errors.New("userlist is clear")
	}

	/*for _, user := range userList{
		if user == username
	}*/

	for i := 0; i < len(userList); i++ {
		if userList[i] == username {
			userList = append(userList[:i], userList[i+1:]...)
			fmt.Println("user logged out successfully!")
			message := fmt.Sprintf("%q logget out succesfully", username)
			history = append(history, message)
			event["Successful actions"] += 1
			return userList, history, nil
		}
	}
	event["Rejected actions"] += 1
	return userList, history, errors.New("invalid logout")
}

func userBuy(userList []string, username string, event map[string]int, history []string) (string, []string, error) {
	var buyMessage string
	for _, user := range userList {
		if user == username {
			buyMessage = "purchase accepted"
			event["Successful actions"] += 1
			message := fmt.Sprintf("%q bought successfully!", username)
			history = append(history, message)
			return buyMessage, history, nil
		}
	}
	event["Rejected actions"] += 1
	return "", history, fmt.Errorf("user(%q) is not logged in", username)
}

func report(users []string, info map[string]int, history []string, uC map[string]int) {
	fmt.Println("\n\n", users, "\n", info, "\n\n", history, "\n\n", uC)
}

func infoMap() map[string]int {
	info := map[string]int{
		"Successful actions": 0,
		"Rejected actions":   0,
	}
	return info
}

package main

import (
	"errors"
	"fmt"
)

func main() {

	words := make(map[string]string)
	var option int

	for {
		fmt.Print("===== Vocabulary Dictionary =====\n1. Show all words\n2. Add word\n3. Find translation\n4. Delete word\n5. Exit\nChoose option: ")
		_, err := fmt.Scanln(&option)

		if err != nil {
			fmt.Println("incorrect input type")
			return
		}

		if option == 5 {
			break
		}

		switch option {
		case 1:
			showAllWords(words)
		case 2:
			err = addWord(words)

			if err != nil {
				fmt.Println(err)
				return
			}

		case 3:
			err = printTranslation(words)

			if err != nil {
				fmt.Println(err)
				return
			}

		case 4:
			err = deleteWord(words)

			if err != nil {
				fmt.Println(err)
				return
			}

		default:
			fmt.Println("incorrect option")

		}

	}

	fmt.Println("ByeBye!")
}

func showAllWords(words map[string]string) {

	if len(words) == 0 {
		fmt.Println("Dictionary is empty")
		return
	}
	for englishWords, russianWords := range words {
		fmt.Println(englishWords, "-", russianWords)
	}
}

func addWord(words map[string]string) error {
	var engWordKey string
	var rusWordValue string
	fmt.Println("Enter new english word with translation down please:")
	fmt.Scanln(&engWordKey, &rusWordValue)

	_, ok := words[engWordKey]

	switch {
	case engWordKey == "" || rusWordValue == "":
		return errors.New("invalid key or value")

	case ok:
		return fmt.Errorf("there is translation under this word")

	default:
		words[engWordKey] = rusWordValue
		fmt.Println("word added successfully")
	}
	return nil
}

func printTranslation(words map[string]string) error {
	var engWord string
	fmt.Println("Enter english word you wanna know the the translation")
	fmt.Scanln(&engWord)

	if engWord == "" {
		return errors.New("empty key")
	}

	value, ok := words[engWord]

	if !ok {
		return errors.New("word not found, check did you write it correctly?")
	}

	fmt.Println(engWord, "means", value)
	return nil
}

func deleteWord(words map[string]string) error {
	var userWord string
	fmt.Println("Enter word which you wanna delete:")
	fmt.Scanln(&userWord)

	_, ok := words[userWord]

	if !ok {
		return errors.New("word not found")
	}

	delete(words, userWord)
	fmt.Println("\n           word deleted successfully\n ")
	return nil
}

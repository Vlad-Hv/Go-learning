package main

import (
	"fmt"
	"math/rand"
)

type Book struct {
	Title  string
	Author string
	Pages  int
	IsRead bool
}

func main() {
	var readBooks int
	books := []Book{
		{
			Title:  "Brick",
			Author: "Daniil",
			Pages:  200,
			IsRead: false,
		},

		{
			Title:  "Death Planet",
			Author: "NN",
			Pages:  3234,
			IsRead: false,
		},

		{
			Title:  "45 mind picture",
			Author: "Daba setre",
			Pages:  34,
			IsRead: false,
		},

		{
			Title:  "Garry Potter",
			Author: "Ronald hfhfh",
			Pages:  316,
			IsRead: false,
		},
	}

	for _, book := range books {
		fmt.Println(book.Title)
	}

	books[rand.Intn(len(books))].IsRead = true

	for _, book := range books {
		if book.IsRead == true {
			readBooks += 1
		}
	}
	fmt.Println(readBooks)
}

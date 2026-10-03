package ui

import (
	"errors"
	"fmt"
	"publish/internal/materials"
	"publish/internal/publicate"
)

func Report(materialByID map[int]publicate.Material, ID int) error {
	if materialByID == nil {
		return errors.New("incorect map")
	}

	material, ok := materialByID[ID].(*materials.Article)
	if !ok {
		newMaterial, ok := materialByID[ID].(*materials.Announcement)
		if !ok {
			return errors.New("incorrect ID")
		}
		announcement(*newMaterial)
		return nil
	}
	article(*material)
	return nil
}

func announcement(a materials.Announcement) {
	fmt.Println("Type: Announcement")
	fmt.Println("Name: ", a.Name)
	fmt.Println("Header: ", a.Header)
	fmt.Println("Status: ", a.Status)
	history(a.ActionLog)
}

func article(a materials.Article) {
	fmt.Println("Type: Article")
	fmt.Println("Name: ", a.Name)
	fmt.Println("Header: ", a.Header)
	fmt.Println("Status: ", a.Status)
	history(a.ActionLog)
}

func history(actionLog []string) {
	fmt.Println("Action Log: ")
	for _, action := range actionLog {
		fmt.Println(action)
	}
	fmt.Println("")
}

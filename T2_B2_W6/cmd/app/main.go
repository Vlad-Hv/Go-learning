package main

import (
	"fmt"
	"publish/internal/materials"
	"publish/internal/publicate"
	"publish/internal/storage"
	"publish/internal/ui"
)

func main() {
	materialsByID := storage.CreateMaterialByID()
	if err := materials.CreateArticle("vlad", "found gold", "GOLD", "hello I'm vlad, I like money, I like Golang, I really like to write on this language", 20, materialsByID); err != nil {
		fmt.Println(err)
		return
	}

	if err := materials.CreateAnnouncement("dima", "lost cat", "POOR CAT", "Guis, I lost my cat please, helm me find them", 30, 19, materialsByID); err != nil {
		fmt.Println(err)
		return
	}

	if err := materialsByID[20].Edit("", "NO GOLD ANYMORE", "", 0); err != nil {
		fmt.Println(err)
		return
	}

	if err := materialsByID[19].Edit("vadim", "", "", -1); err != nil {
		fmt.Println(err)
		return
	}

	err := publicate.Publicate(materialsByID[20])
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := materialsByID[20].Edit("dime", "", "", 0); err != nil {
		fmt.Println(err)
	}
	ui.Report(materialsByID, 19)
	ui.Report(materialsByID, 20)
	if err := ui.Report(materialsByID, 21); err != nil {
		fmt.Println(err)
		return
	}
}

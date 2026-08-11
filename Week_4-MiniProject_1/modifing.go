package main

import (
	"github.com/google/uuid"
)

func makePriorityMap() map[int]string {
	return map[int]string{
		1: "low",
		2: "medium",
		3: "high",
	}
}

func modifyIndexToUUID(tasks *[]Task) map[int]uuid.UUID {
	//add validation
	indexToUUID := make(map[int]uuid.UUID)

	for index := range *tasks {
		indexToUUID[index] = (*tasks)[index].ID
	}

	return indexToUUID
}

func MapUuidToTask(tasks *[]Task) map[uuid.UUID]*Task {
	// add validation
	mapa := make(map[uuid.UUID]*Task)

	for index := range *tasks {
		mapa[(*tasks)[index].ID] = &(*tasks)[index]
	}

	return mapa

}

package main

import "github.com/google/uuid"

func syncTaskMaps(tasks *[]Task, indexToUUID *map[int]uuid.UUID, UuidToTask *map[uuid.UUID]*Task) {
	*indexToUUID = modifyIndexToUUID(tasks)
	*UuidToTask = MapUuidToTask(tasks)
}

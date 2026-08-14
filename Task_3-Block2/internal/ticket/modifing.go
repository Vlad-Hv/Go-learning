package ticket

func CreateIndexToUser(tickets *[]Ticket) map[int]*Ticket {
	indexToUserMap := make(map[int]*Ticket)
	for index := range *tickets {
		indexToUserMap[index+1] = &(*tickets)[index]
	}
	return indexToUserMap
}

func CreateTicketOptions() map[int]int {
	return map[int]int{
		1: 5,
		2: 10,
		3: 30,
	}
}

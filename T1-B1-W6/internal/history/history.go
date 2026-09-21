package history

import "fmt"

type History []string

func Create() History {
	var history History
	return history
}

func (h *History) Add(firstMessage string, whoIs string, secondMessage string) error {
	message := fmt.Sprint(firstMessage, " ", whoIs, " ", secondMessage)
	err := validateHistory(firstMessage)
	if err != nil {
		return err
	}

	*h = append(*h, message)
	return nil
}

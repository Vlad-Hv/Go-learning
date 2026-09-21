package history

import (
	"errors"
)

var (
	errMessageEmpty = errors.New("empty storage")
)

func validateHistory(message string) error {
	if message == "" {
		return errMessageEmpty
	}
	return nil
}

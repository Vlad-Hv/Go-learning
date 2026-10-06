package internal

import (
	"errors"
	"testing"
)

type Storages map[int]string

var (
	errTestDecide = errors.New("testing")
)

func (s Storages) Find(ID int) (bool, error) {
	/*result := s[ID]
	if result != ""{
		return true, nil
	}
	return false, nil*/
	return true, errTestDecide
}

func TestDeciding(t *testing.T) {
	tests := []struct {
		name    string
		storage Storages
		wantErr error
	}{
		{
			name:    "check storage error",
			wantErr: errTestDecide,
			storage: Storages{10: "test"},
		},
	}

	for _, tt := range tests {
		Service := NewService(tt.storage)
		result, err := Service.Deciding(10)
		if !errors.Is(err, tt.wantErr) {
			t.Fatalf("expected %v error, got %v error", tt.wantErr, err)
		}

		if result != "" {
			t.Error("unexpected result")
		}
	}
}

package history

import (
	"errors"
	"slices"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name          string
		history       History
		wantHistory   History
		firstMessage  string
		whoIs         string
		secondMessage string
		wantErr       error
	}{
		{
			name:          "successful case",
			wantHistory:   History{"mail for vlad delivered successful"},
			firstMessage:  "mail for",
			whoIs:         "vlad",
			secondMessage: "delivered successful",
			wantErr:       nil,
		},

		{
			name:    "empty message",
			wantErr: errMessageEmpty,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.history.Add(tt.firstMessage, tt.whoIs, tt.secondMessage)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if !slices.Equal(tt.history, tt.wantHistory) {
				t.Errorf("expected history %v, got %v", tt.wantHistory, tt.history)
			}
		})
	}
}

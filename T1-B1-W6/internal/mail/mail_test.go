package mail

import (
	"errors"
	"testing"
)

func TestMarkAsDelivered(t *testing.T) {
	tests := []struct {
		name     string
		mail     Mail
		wantMail Mail
		wantErr  error
	}{
		{
			name: "success case",
			mail: Mail{
				State: "undelivered",
			},
			wantMail: Mail{
				State: "delivered",
			},
			wantErr: nil,
		},

		{
			name: "already delivered mail case",
			mail: Mail{
				State: "delivered",
			},
			wantMail: Mail{
				State: "delivered",
			},
			wantErr: ErrMailDelivered,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.mail.MarkAsDelivered()

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if tt.mail != tt.wantMail {
				t.Errorf("expected mail state %v, actual %v", tt.wantMail.State, tt.mail.State)
			}
		})
	}
}

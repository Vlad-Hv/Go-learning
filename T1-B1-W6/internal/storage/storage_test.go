package storage

import (
	"delivery/internal/mail"
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name    string
		storage map[int]*mail.Mail
		id      int
		mail    mail.Mail
		wantOk  bool
		wantErr error
	}{
		{
			name:    "succesfull case",
			storage: make(map[int]*mail.Mail),
			id:      10,
			mail: mail.Mail{
				Recipient: "vlad",
				Weight:    9,
				State:     "undelivered",
			},
			wantOk:  true,
			wantErr: nil,
		},

		{
			name: "repeating ID",
			storage: map[int]*mail.Mail{
				10: {
					Recipient: "vlad",
					Weight:    8,
				},
			},
			id: 10,
			mail: mail.Mail{
				Recipient: "dima",
				Weight:    5,
				State:     "Delivered",
				Price:     42,
			},
			wantOk:  true,
			wantErr: errIDUsing,
		},

		{
			name:    "invalid ID",
			storage: make(map[int]*mail.Mail),
			id:      -10,
			mail: mail.Mail{
				Recipient: "vlad",
				Weight:    20,
				State:     "undelivered",
			},
			wantOk:  false,
			wantErr: errIDInvalid,
		},

		{
			name:    "empty recipient",
			storage: make(map[int]*mail.Mail),
			id:      10,
			mail: mail.Mail{
				Recipient: "",
				Weight:    20,
				State:     "undelivered",
			},
			wantOk:  false,
			wantErr: errRecipientNameInvalid,
		},

		{
			name:    "invalid weight",
			storage: make(map[int]*mail.Mail),
			id:      10,
			mail: mail.Mail{
				Recipient: "vlad",
				Weight:    -4,
				State:     "undelivered",
			},
			wantOk:  false,
			wantErr: errWeightInvalid,
		},

		{
			name:    "invalid zero weight",
			storage: make(map[int]*mail.Mail),
			id:      10,
			mail: mail.Mail{
				Recipient: "vlad",
				Weight:    0,
				State:     "undelivered",
			},
			wantOk:  false,
			wantErr: errWeightInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var oldMail mail.Mail
			if tt.wantErr == errIDUsing {
				oldMail = *tt.storage[tt.id]
			}
			err := Add(tt.storage, tt.id, tt.mail)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			newMail, ok := tt.storage[tt.id]

			if ok != tt.wantOk {
				t.Fatal("expected that mail adree exist")
			}

			if ok && tt.wantErr != errIDUsing {
				if newMail == nil {
					t.Fatal("unexpected nil mail address")
				}
				if *newMail != tt.mail {
					t.Errorf("expected address mail ID: %d address", tt.id)
				}
				return
			}

			if ok && tt.wantErr == errIDUsing {
				if newMail == nil {
					t.Fatal("expected old mail")
				}

				if *newMail != oldMail {
					t.Errorf("expected mail %v, got %v", tt.mail, *newMail)
				}
			}

		})
	}
}

package deliveries

import (
	"delivery/internal/history"
	"delivery/internal/mail"
	"errors"
	"slices"
	"testing"
)

var (
	errFakeDelivery = errors.New("interface validating")
)

type fakeDeliverier struct {
}

func (fd fakeDeliverier) Delivery(mail *mail.Mail) (int, error) {
	return 0, errFakeDelivery
}

func TestDelivery(t *testing.T) {
	var emptyMail *mail.Mail
	tests := []struct {
		name        string
		mail        *mail.Mail
		deliverier  delivery
		history     history.History
		wantMail    mail.Mail
		wantHistory history.History
		wantErr     error
	}{
		{
			name: "success person case",
			mail: &mail.Mail{
				Recipient: "vlad",
				Weight:    6,
				State:     "undelivered",
			},
			deliverier: Person{Price: 10},
			history:    history.History{"mail to valera Delivered successfully"},
			wantMail: mail.Mail{
				Recipient: "vlad",
				Weight:    6,
				State:     "delivered",
				Price:     160,
			},
			wantHistory: history.History{"mail to valera Delivered successfully", "mail to vlad Delivered successfully"},
			wantErr:     nil,
		},

		{
			name: "success service case",
			mail: &mail.Mail{
				Recipient: "dima",
				Weight:    30,
				State:     "undelivered",
			},
			history:    history.History{"mail to valera Delivered successfully"},
			deliverier: Service{Price: 10},
			wantMail: mail.Mail{
				Recipient: "dima",
				Weight:    30,
				State:     "delivered",
				Price:     600,
			},

			wantHistory: history.History{"mail to valera Delivered successfully", "mail to dima Delivered successfully"},
			wantErr:     nil,
		},

		{
			name:        "empty mail address",
			deliverier:  Person{Price: 10},
			history:     history.History{"mail to valera Delivered successfully"},
			wantHistory: history.History{"mail to valera Delivered successfully"},
			wantErr:     errAddressIncorrect,
			mail:        emptyMail,
		},

		{
			name: "interface check",
			mail: &mail.Mail{
				Recipient: "vlad",
				Weight:    20,
				State:     "undelivered",
			},
			deliverier: fakeDeliverier{},
			wantMail: mail.Mail{
				Recipient: "vlad",
				Weight:    20,
				State:     "undelivered",
			},

			history:     history.History{"mail to dima Delivered successfully"},
			wantHistory: history.History{"mail to dima Delivered successfully"},
			wantErr:     errFakeDelivery,
		},

		{
			name: "big weight to person delivery",
			mail: &mail.Mail{
				Recipient: "vlad",
				Weight:    20,
				State:     "undelivered",
			},
			deliverier:  Person{Price: 10},
			history:     history.History{"mail to valera Delivered successfully"},
			wantHistory: history.History{"mail to valera Delivered successfully"},
			wantMail: mail.Mail{
				Recipient: "vlad",
				Weight:    20,
				State:     "undelivered",
			},
			wantErr: errDeliveryBigWeight,
		},

		{
			name: "already delivered mail case",
			mail: &mail.Mail{
				Recipient: "vlad",
				Weight:    9,
				State:     "delivered",
			},
			deliverier:  Person{Price: 10},
			history:     history.History{"mail to valera Delivered successfully"},
			wantHistory: history.History{"mail to valera Delivered successfully"},
			wantMail: mail.Mail{
				Recipient: "vlad",
				Weight:    9,
				State:     "delivered",
			},
			wantErr: mail.ErrMailDelivered,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Delivery(tt.deliverier, tt.mail, &tt.history)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if !slices.Equal(tt.history, tt.wantHistory) {
				t.Errorf("expected history %v, got %v", tt.wantHistory, tt.history)
			}

			if tt.wantErr != errAddressIncorrect {
				if *tt.mail != tt.wantMail {
					t.Errorf("expected mail %v, got %v", tt.wantMail, tt.mail)
				}
			}

		})
	}
}

//успешный кейс для двух служб, кейс пустого адреса, кейс большого веса для доставщика,
//везде проверять чтобы была посылка которую мы хотим и историю

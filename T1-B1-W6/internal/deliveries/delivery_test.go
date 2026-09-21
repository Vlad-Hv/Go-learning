package deliveries

import (
	"delivery/internal/mail"
	"errors"
	"testing"
)

func TestDeliveryPerson(t *testing.T) {
	tests := []struct {
		name      string
		mail      mail.Mail
		person    Person
		wantPrice int
		wantErr   error
	}{
		{
			name: "success case",
			mail: mail.Mail{
				Recipient: "vlad",
				Weight:    5,
				State:     "undelivered",
			},
			person:    Person{Price: 20},
			wantPrice: 200,
			wantErr:   nil,
		},

		{
			name: "big weight",
			mail: mail.Mail{
				Recipient: "vlad",
				Weight:    12,
				State:     "udelivered",
			},
			person:    Person{Price: 20},
			wantPrice: 0,
			wantErr:   errDeliveryBigWeight,
		},

		{
			name: "edge case",
			mail: mail.Mail{
				Recipient: "vlad",
				Weight:    10,
				State:     "undelivery",
			},
			person:    Person{Price: 20},
			wantPrice: 300,
			wantErr:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			price, err := tt.person.Delivery(&tt.mail)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if price != tt.wantPrice {
				t.Errorf("expected price %d, actual %d", tt.wantPrice, price)
			}
		})
	}
}

func TestDeliveryService(t *testing.T) {
	tests := []struct {
		name      string
		service   Service
		mail      mail.Mail
		wantPrice int
		wantErr   error
	}{
		{
			name:    "success case",
			service: Service{Price: 10},
			mail: mail.Mail{
				Recipient: "clad",
				Weight:    50,
				State:     "undeliverid",
			},
			wantPrice: 800,
			wantErr:   nil,
		},

		{
			name:    "edge case",
			service: Service{Price: 10},
			mail: mail.Mail{
				Recipient: "vlad",
				Weight:    100,
				State:     "undelivered",
			},
			wantPrice: 1300,
			wantErr:   nil,
		},

		{
			name:    "too big weight",
			service: Service{Price: 10},
			mail: mail.Mail{
				Recipient: "vlad",
				Weight:    150,
				State:     "undelivered",
			},
			wantPrice: 0,
			wantErr:   errServiceBigWeight,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			price, err := tt.service.Delivery(&tt.mail)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if price != tt.wantPrice {
				t.Errorf("expected price %d, actual %d", tt.wantPrice, price)
			}
		})
	}
}

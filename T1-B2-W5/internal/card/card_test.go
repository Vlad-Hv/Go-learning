package card

import (
	"errors"
	"testing"
)

func TestPay(t *testing.T) {
	tests := []struct {
		name            string
		balance         int
		price           int
		isAvailable     bool
		expectedBalance int
		wantErr         error
	}{
		{name: "success pay case", balance: 100, price: 30, isAvailable: true, expectedBalance: 70, wantErr: nil},
		{name: "success pay full balance", balance: 100, price: 100, isAvailable: true, expectedBalance: 0, wantErr: nil},
		{name: "unavailable card pay", balance: 100, price: 30, isAvailable: false, expectedBalance: 100, wantErr: errBlockedCard},
		{name: "price is zero", balance: 100, price: 0, isAvailable: true, expectedBalance: 100, wantErr: errInvalidPrice},
		{name: "price less than zero", balance: 100, price: -1, isAvailable: true, expectedBalance: 100, wantErr: errInvalidPrice},
		{name: "price more than balance", balance: 100, price: 1000, isAvailable: true, expectedBalance: 100, wantErr: errNotEnoughMoney},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := Card{Balance: tt.balance, IsAvailable: tt.isAvailable}
			err := card.Pay(tt.price)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}
			if card.Balance != tt.expectedBalance {
				t.Errorf("expected %d, actual %d", tt.expectedBalance, card.Balance)
			}
		})
	}
}

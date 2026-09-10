package card

import (
	"errors"
	"testing"
)

func TestPay(t *testing.T) {
	tests := []struct {
		name        string
		balance     int
		price       int
		isBlocked   bool
		wantBalance int
		wantErr     error
	}{
		{name: "success case", balance: 100, price: 30, isBlocked: false, wantBalance: 70, wantErr: nil},
		{name: "success full balance case", balance: 100, price: 100, isBlocked: false, wantBalance: 0, wantErr: nil},
		{name: "blocked card case", balance: 100, price: 30, isBlocked: true, wantBalance: 100, wantErr: errCardBlocked},
		{name: "zero price", balance: 100, price: 0, isBlocked: false, wantBalance: 100, wantErr: errInvalidPrice},
		{name: "less than zero price", balance: 100, price: -1, isBlocked: false, wantBalance: 100, wantErr: errInvalidPrice},
		{name: "price more than balance", balance: 100, price: 1000, isBlocked: false, wantBalance: 100, wantErr: errNotEnoughMoney},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			card := Card{Balance: tt.balance, IsBlocked: tt.isBlocked}
			err := card.Pay(tt.price)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}
			if card.Balance != tt.wantBalance {
				t.Errorf("expected %d, actual %d", tt.wantBalance, card.Balance)
			}
		})
	}
}

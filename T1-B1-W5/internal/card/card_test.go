package card

import (
	"testing"
)

func TestPaySuccess(t *testing.T) {
	card := Card{Balance: 150, IsAvailable: true}

	err := card.Pay(100)

	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}

	if card.Balance != 50 {
		t.Errorf("expected 50, actual %d", card.Balance)
	}
}

func TestPayInsufficientBalance(t *testing.T) {
	card := Card{Balance: 50, IsAvailable: true}
	err := card.Pay(100)

	if err == nil {
		t.Fatalf("expected an error, got %v", err)
	}

	if card.Balance != 50 {
		t.Errorf("expected 50, got %d", card.Balance)
	}
}

func TestPayWithBlockedCard(t *testing.T) {
	card := Card{Balance: 50, IsAvailable: false}

	err := card.Pay(20)

	if err == nil {
		t.Fatalf("expected an error, got %v", err)
	}

	if card.Balance != 50 {
		t.Errorf("expected 50, actual %d", card.Balance)
	}

}

func TestPayWithZeroPrice(t *testing.T) {
	card := Card{Balance: 50, IsAvailable: true}

	err := card.Pay(0)

	if err == nil {
		t.Fatalf("expected an error, got %v", err)
	}

	if card.Balance != 50 {
		t.Errorf("expected 50, actual %d", card.Balance)
	}
}

func TestPayWithPriceLessThanZero(t *testing.T) {
	card := Card{Balance: 50, IsAvailable: true}

	err := card.Pay(-10)

	if err == nil {
		t.Fatalf("expected an error, got %v", err)
	}

	if card.Balance != 50 {
		t.Errorf("expected 50, actual %d", card.Balance)
	}
}

func TestPayWithFullBalance(t *testing.T) {
	card := Card{Balance: 50, IsAvailable: true}

	err := card.Pay(50)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if card.Balance != 0 {
		t.Errorf("expected 0, actual %d", card.Balance)
	}
}

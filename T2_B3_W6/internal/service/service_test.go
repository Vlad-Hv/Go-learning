package service

import (
	"errors"
	"maps"
	"product/internal/storage"
	"testing"
)

var (
	errTestError = errors.New("test")
)

func TestReserveProducte(t *testing.T) {
	tests := []struct {
		name          string
		reserveAmount int
		storage       storage.ProductList
		ID            int
		wantStorage   storage.ProductList
		wantActual    int
		wantErr       error
	}{
		{
			name:          "succesfull case",
			reserveAmount: 10,
			storage: storage.ProductList{20: storage.Product{
				Name:       "Tea",
				FreeAmount: 20,
			},
			},
			ID: 20,
			wantStorage: storage.ProductList{20: storage.Product{
				Name:       "Tea",
				FreeAmount: 10,
			},
			},
			wantActual: 10,
			wantErr:    nil,
		},

		{
			name:          "unexist ID",
			reserveAmount: 10,
			storage: storage.ProductList{20: storage.Product{
				Name:       "Tea",
				FreeAmount: 20,
			},
			},
			ID: 19,
			wantStorage: storage.ProductList{20: storage.Product{
				Name:       "Tea",
				FreeAmount: 20,
			},
			},
			wantActual: 0,
			wantErr:    storage.ErrProductUnexist,
		},

		{
			name:          "not enough amount",
			reserveAmount: 30,
			storage: storage.ProductList{20: storage.Product{
				Name:       "Tea",
				FreeAmount: 20,
			},
			},
			ID: 20,
			wantStorage: storage.ProductList{20: storage.Product{
				Name:       "Tea",
				FreeAmount: 20,
			},
			},
			wantActual: 0,
			wantErr:    errAmountNotEnought,
		},

		{
			name:          "invalid amount",
			reserveAmount: -10,
			storage: storage.ProductList{20: storage.Product{
				Name:       "Tea",
				FreeAmount: 20,
			},
			},
			ID: 20,
			wantStorage: storage.ProductList{20: storage.Product{
				Name:       "Tea",
				FreeAmount: 20,
			},
			},
			wantActual: 0,
			wantErr:    errAmountInvalid,
		},

		{
			name:          "zero invalid amount ",
			reserveAmount: 0,
			storage: storage.ProductList{20: storage.Product{
				Name:       "Tea",
				FreeAmount: 20,
			},
			},
			ID: 20,
			wantStorage: storage.ProductList{20: storage.Product{
				Name:       "Tea",
				FreeAmount: 20,
			},
			},
			wantActual: 0,
			wantErr:    errAmountInvalid,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := New(&tt.storage)
			actual, err := service.ReserveProducts(tt.reserveAmount, tt.ID)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}

			if actual != tt.wantActual {
				t.Errorf("expected actual %v, got %v", tt.wantActual, actual)
			}

			/*if tt.wantErr == storage.ErrProductUnexist {
				if tt.storage[20] != tt.wantStorage[20] {
					t.Errorf("expected storage %v, got %v", tt.wantStorage[tt.ID], tt.storage[tt.ID])

				}
			} else {
				if tt.storage[tt.ID] != tt.wantStorage[tt.ID] {
					t.Errorf("expected storage %v, got %v", tt.wantStorage[tt.ID], tt.storage[tt.ID])
				}
			}*/

			if !maps.Equal(tt.storage, tt.wantStorage) {
				t.Errorf("expected storage %v, got %v", tt.wantStorage[tt.ID], tt.storage[tt.ID])
			}
		})
	}
}

//усппешный кейс, отсутствующий айди, недостаточно количество, неправильное количество, неправильное количество с 0
//новый тип вызывающий ошибку в новой функции

type testStorage struct{}

func (t testStorage) CheckProduct(ID int) (int, error) {
	return 12, nil
}

func (t testStorage) Reserve(id, amount int) error {
	return errTestError
}
func TestReserveProducteNewType(t *testing.T) {
	storage := testStorage{}
	service := New(storage)
	actual, err := service.ReserveProducts(10, 10)
	wantActual := 0
	wantErr := errTestError

	if !errors.Is(err, wantErr) {
		t.Fatalf("expected error %v, got %v", wantErr, err)
	}

	if actual != wantActual {
		t.Errorf("expected actual %d, got %d", wantActual, actual)
	}
}

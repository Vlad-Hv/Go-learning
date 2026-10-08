package storage

import "errors"

var (
	ErrProductUnexist = errors.New("product unexist")
)

type Storage interface {
	CheckProduct(ID int) (int, error)
	Reserve(ID, amount int) error
}

type ProductList map[int]Product

func CreateMap() ProductList {
	return ProductList{
		10: {
			Name:       "Coffee",
			FreeAmount: 20,
		},

		11: {
			Name:       "Tea",
			FreeAmount: 10,
		},

		12: {
			Name:       "Water",
			FreeAmount: 100,
		},
	}
}

func (pl ProductList) CheckProduct(ID int) (int, error) {
	product, ok := pl[ID]
	if ok {
		return product.FreeAmount, nil
	} else {
		return product.FreeAmount, ErrProductUnexist
	}
}

func (pl *ProductList) Reserve(ID int, amount int) error {
	product := (*pl)[ID]
	product.FreeAmount = amount
	(*pl)[ID] = product
	return nil
}

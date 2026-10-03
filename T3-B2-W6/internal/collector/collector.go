package collector

import (
	"collector/internal/components"
)

type Sender interface {
	Send()
}

func Create() []Sender {
	return []Sender{
		&components.DataBase{
			BaseInfo: components.BaseInfo{
				Name:      "MySQL",
				IsCheking: true,
			},
			IsConect: true,
		},

		&components.Queue{
			BaseInfo: components.BaseInfo{
				Name:      "INFO",
				IsCheking: true,
			},
			Limit:        20,
			WaitMessages: 10,
		},

		&components.Storage{
			BaseInfo: components.BaseInfo{
				Name:      "file one",
				IsCheking: true,
			},
			Min:    15,
			Actual: 25,
		},

		&components.Queue{
			BaseInfo: components.BaseInfo{
				Name:      "ILFO",
				IsCheking: false,
			},
			Limit:        20,
			WaitMessages: 10,
		},

		&components.Storage{
			BaseInfo: components.BaseInfo{
				Name:      "file two",
				IsCheking: true,
			},
			Min:    25,
			Actual: 15,
		},

		&components.DataBase{
			BaseInfo: components.BaseInfo{
				Name:      "MongoDB",
				IsCheking: true,
			},
			IsConect: false,
		},
	}
}

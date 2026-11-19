package order

type Order struct {
	ID          int
	Sum         int
	HandlerTime int
	WasHandled  bool
}

func Create() []Order {
	result := []Order{
		{
			ID:          112,
			Sum:         10,
			HandlerTime: 30,
		},

		{
			ID:          113,
			Sum:         13,
			HandlerTime: 30,
		},

		{
			ID:          114,
			Sum:         -10,
			HandlerTime: 20,
		},

		{
			ID:          115,
			Sum:         49,
			HandlerTime: 80,
		},

		{
			ID:          116,
			Sum:         10,
			HandlerTime: 30,
		},

		{
			ID:          117,
			Sum:         13,
			HandlerTime: 300,
		},

		{
			ID:          118,
			Sum:         -10,
			HandlerTime: 210,
		},

		{
			ID:          119,
			Sum:         49,
			HandlerTime: 80,
		},
		{
			ID:          120,
			Sum:         10,
			HandlerTime: 30,
		},

		{
			ID:          121,
			Sum:         13,
			HandlerTime: 300,
		},

		{
			ID:          122,
			Sum:         -10,
			HandlerTime: 210,
		},

		{
			ID:          123,
			Sum:         49,
			HandlerTime: 80,
		},
	}

	return result
}

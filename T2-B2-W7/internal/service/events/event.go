package event

type Event struct {
	Name  string
	Text  string
	Delay int
}

func CreateAFlow() []Event {
	return []Event{
		{
			Name:  "A",
			Text:  "a1",
			Delay: 50,
		},

		{
			Name:  "A",
			Text:  "a2",
			Delay: 100,
		},
	}
}

func CreateBFlow() []Event {
	return []Event{
		{
			Name:  "B",
			Text:  "b1",
			Delay: 100,
		},

		{
			Name:  "B",
			Text:  "b2",
			Delay: 320,
		},

		{
			Name:  "B",
			Text:  "b3",
			Delay: 450,
		},
	}
}

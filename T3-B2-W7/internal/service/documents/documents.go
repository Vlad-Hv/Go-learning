package documents

type Document struct {
	ID          string
	Text        string
	SourceDelay int
	HandleDelay int
}

func CreateA() []Document {
	return []Document{
		{
			ID:          "A1",
			Text:        "go channels",
			HandleDelay: 60,
		},
		{
			ID:          "A2",
			Text:        "cancel the operation",
			SourceDelay: 100,
			HandleDelay: 800,
		},
	}
}

func CreateB() []Document {
	return []Document{
		{
			ID:          "B1",
			Text:        "hello",
			SourceDelay: 30,
			HandleDelay: 60,
		},
		{
			ID:          "B2",
			Text:        "wait for all workers",
			SourceDelay: 180,
			HandleDelay: 800,
		},
	}
}

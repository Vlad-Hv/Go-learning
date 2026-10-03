package materials

type Profile struct {
	Name        string
	Description string
}

type Material struct {
	//ID     int
	Header    string
	Status    string
	ActionLog []string
}

func (m *Material) Add(message string) error {
	if err := validateAdd(message); err != nil {
		return err
	}

	m.ActionLog = append(m.ActionLog, message)
	return nil
}

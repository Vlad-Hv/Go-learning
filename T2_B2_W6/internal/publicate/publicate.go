package publicate

type Edit interface {
	Edit(string, string, string, int) error
}

type Material interface {
	Edit
	Publicate() error
	Add(string) error
}

func Publicate(material Material) error {
	if err := material.Publicate(); err != nil {
		return err
	}

	if err := material.Add("publicate done successfully"); err != nil {
		return err
	}
	return nil
}

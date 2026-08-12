package profile

type Profile struct {
	Name  string
	Level int
}

func CreateProfile(name string, level int) Profile {
	return Profile{
		Name:  name,
		Level: level,
	}
}

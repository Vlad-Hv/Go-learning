package internal

type Storageer interface {
	Find(int) (bool, error)
}

type Service struct {
	Storage Storageer
}

func NewService(storage Storageer) *Service {
	return &Service{Storage: storage}
}

func (s Service) Deciding(ID int) (string, error) {
	active, err := s.Storage.Find(ID)
	if err != nil {
		return "", err
	}

	if active {
		return "access granted", nil
	} else {
		return "access denied", nil
	}
}

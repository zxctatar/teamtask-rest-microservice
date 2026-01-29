package projectdomain

import "time"

type ProjectDomain struct {
	Id        uint32
	OwnerId   uint32
	Name      string
	CreatedAt time.Time
}

func NewProjectDomain(ownerId uint32, name string) (*ProjectDomain, error) {
	if err := validateOwnerId(ownerId); err != nil {
		return nil, err
	}
	if err := validateName(name); err != nil {
		return nil, err
	}

	return &ProjectDomain{
		OwnerId: ownerId,
		Name:    name,
	}, nil
}

func RestoreProjectDomain(id, ownerId uint32, name string, createdAt time.Time) *ProjectDomain {
	return &ProjectDomain{
		Id:        id,
		OwnerId:   ownerId,
		Name:      name,
		CreatedAt: createdAt,
	}
}

func validateOwnerId(ownerId uint32) error {
	if ownerId == 0 {
		return ErrInvalidOwnerId
	}
	return nil
}

func validateName(name string) error {
	rName := []rune(name)
	if len(rName) > 255 {
		return ErrInvalidName
	}
	return nil
}

package createmodel

type CreateProjectInput struct {
	OwnerId uint32
	Name    string
}

func NewCreateProjectInput(ownertId uint32, name string) *CreateProjectInput {
	return &CreateProjectInput{
		OwnerId: ownertId,
		Name:    name,
	}
}

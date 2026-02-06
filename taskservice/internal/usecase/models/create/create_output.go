package createmodel

type CreateOutput struct {
	IsCreated bool
}

func NewCreateOutput(isCreated bool) *CreateOutput {
	return &CreateOutput{
		IsCreated: isCreated,
	}
}

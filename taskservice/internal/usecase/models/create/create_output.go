package createmodel

type CreateOutput struct {
	ProjectId uint32
}

func NewCreateOutput(projectId uint32) *CreateOutput {
	return &CreateOutput{
		ProjectId: projectId,
	}
}

package createmodel

type CreateTaskOutput struct {
	ProjectId uint32
}

func NewCreateOutput(projectId uint32) *CreateTaskOutput {
	return &CreateTaskOutput{
		ProjectId: projectId,
	}
}

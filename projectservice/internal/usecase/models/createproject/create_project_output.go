package createmodel

type CreateProjectOutput struct {
	ProjectId uint32
}

func NewCreateProjectOutput(projectId uint32) *CreateProjectOutput {
	return &CreateProjectOutput{
		ProjectId: projectId,
	}
}

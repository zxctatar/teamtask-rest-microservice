package updatemodel

type UpdateProjectOutput struct {
	IsUpdated bool
}

func NewUpdateProjectOutput(isUpdated bool) *UpdateProjectOutput {
	return &UpdateProjectOutput{
		IsUpdated: isUpdated,
	}
}

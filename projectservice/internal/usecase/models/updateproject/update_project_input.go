package updatemodel

type UpdateProjectInput struct {
	OwnerId   uint32
	ProjectId uint32
	NewName   *string
}

func NewUpdateProjectInput(ownerId uint32, projectId uint32, newName *string) *UpdateProjectInput {
	return &UpdateProjectInput{
		OwnerId:   ownerId,
		ProjectId: projectId,
		NewName:   newName,
	}
}

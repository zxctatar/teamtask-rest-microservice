package createmodel

import "time"

type CreateInput struct {
	ProjectId   uint32
	Description string
	Deadline    time.Time
}

func NewCreateInput(projectId uint32, descriprion string, deadline time.Time) *CreateInput {
	return &CreateInput{
		ProjectId:   projectId,
		Description: descriprion,
		Deadline:    deadline,
	}
}

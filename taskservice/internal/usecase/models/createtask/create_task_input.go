package createmodel

import "time"

type CreateTaskInput struct {
	ProjectId   uint32
	Description string
	Deadline    time.Time
}

func NewCreateInput(projectId uint32, descriprion string, deadline time.Time) *CreateTaskInput {
	return &CreateTaskInput{
		ProjectId:   projectId,
		Description: descriprion,
		Deadline:    deadline,
	}
}

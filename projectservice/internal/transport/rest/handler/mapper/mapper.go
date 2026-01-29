package handlmapper

import (
	createdto "projectservice/internal/transport/rest/handler/dto/create"
	createmodel "projectservice/internal/usecase/models/createproject"
)

func CreateRequestToInput(cr *createdto.CreateRequest, userId uint32) *createmodel.CreateProjectInput {
	return createmodel.NewCreateProjectInput(userId, cr.Name)
}

func CreateOutputToResponse(co *createmodel.CreateProjectOutput) *createdto.CreateResponse {
	return &createdto.CreateResponse{
		IsCreated: co.IsCreated,
	}
}

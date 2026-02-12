package interfaces

import (
	"context"
	updatemodel "projectservice/internal/usecase/models/updateproject"
)

type UpdateProjectUsecase interface {
	Execute(ctx context.Context, in *updatemodel.UpdateProjectInput) (*updatemodel.UpdateProjectOutput, error)
}

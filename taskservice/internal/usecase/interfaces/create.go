package interfaces

import (
	"context"
	createmodel "taskservice/internal/usecase/models/create"
)

type CreateUsecase interface {
	Execute(ctx context.Context, in *createmodel.CreateInput) (*createmodel.CreateOutput, error)
}

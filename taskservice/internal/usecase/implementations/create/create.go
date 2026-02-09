package createuc

import (
	"context"
	"log/slog"
	taskdomain "taskservice/internal/domain/task"
	"taskservice/internal/repository/storage"
	createmodel "taskservice/internal/usecase/models/create"
)

type CreateUC struct {
	log *slog.Logger

	stor storage.StorageRepo
}

func NewCreateUC(log *slog.Logger, stor storage.StorageRepo) *CreateUC {
	return &CreateUC{
		log:  log,
		stor: stor,
	}
}

func (c *CreateUC) Execute(ctx context.Context, in *createmodel.CreateInput) (*createmodel.CreateOutput, error) {
	const op = "createuc.Execute"

	log := c.log.With(slog.String("op", op), slog.Int("projectId", int(in.ProjectId)))

	log.Info("starting create task")

	td, err := taskdomain.NewTaskDomain(
		in.ProjectId,
		in.Description,
		in.Deadline,
	)
	if err != nil {
		log.Info("cannot create task", slog.String("error", err.Error()))
		return nil, err
	}

	id, err := c.stor.Save(ctx, td)
	if err != nil {
		log.Warn("cannot save task", slog.String("error", err.Error()))
		return nil, err
	}

	log.Info("task created successfully")

	return createmodel.NewCreateOutput(id), nil
}

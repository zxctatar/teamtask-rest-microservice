package deleteproject

import (
	"context"
	"errors"
	"log/slog"
	projectdomain "projectservice/internal/domain/project"
	"projectservice/internal/repository/storage"
	deleteerr "projectservice/internal/usecase/error/deleteproject"
	deletemodel "projectservice/internal/usecase/models/deleteproject"
)

type DeleteProjectUC struct {
	log *slog.Logger

	stor storage.Storage
}

func NewDeleteProjectUC(log *slog.Logger, stor storage.Storage) *DeleteProjectUC {
	return &DeleteProjectUC{
		log:  log,
		stor: stor,
	}
}

func (d *DeleteProjectUC) Execute(ctx context.Context, in *deletemodel.DeleteProjectInput) (*deletemodel.DeleteProjectOutput, error) {
	const op = "deleteproject.Execute"

	log := d.log.With(slog.String("op", op), slog.Int("userId", int(in.OwnerId)))

	log.Info("starting delete project")

	proj, err := projectdomain.NewProjectDomain(in.OwnerId, in.Name)
	if err != nil {
		log.Info("can't delete a project", slog.String("error", err.Error()))
		return deletemodel.NewDeleteProjectOutput(false), err
	}

	err = d.stor.Delete(ctx, proj)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			log.Info("project not found")
			return deletemodel.NewDeleteProjectOutput(false), deleteerr.ErrProjectNotFound
		}
		log.Warn("error delete project", slog.String("error", err.Error()))
		return deletemodel.NewDeleteProjectOutput(false), err
	}

	log.Info("project deleted")

	return deletemodel.NewDeleteProjectOutput(true), nil
}

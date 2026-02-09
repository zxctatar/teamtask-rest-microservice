package getallprojects

import (
	"context"
	"errors"
	"log/slog"
	"projectservice/internal/repository/storage"
	getallerr "projectservice/internal/usecase/error/getallprojects"
	getallmodel "projectservice/internal/usecase/models/getallprojects"
)

type GetAllProjectsUC struct {
	log *slog.Logger

	stor storage.StorageRepo
}

func NewGetAllProjectsUC(log *slog.Logger, stor storage.StorageRepo) *GetAllProjectsUC {
	return &GetAllProjectsUC{
		log:  log,
		stor: stor,
	}
}

func (g *GetAllProjectsUC) Execute(ctx context.Context, in *getallmodel.GetAllProjectsInput) (*getallmodel.GetAllProjectsOutput, error) {
	const op = "getallprojects.Execute"

	log := g.log.With(slog.String("op", op), slog.Int("userId", int(in.OwnerId)))

	log.Info("starting get all projects request")

	projects, err := g.stor.GetAll(ctx, in.OwnerId)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			log.Info("projects not found")
			return getallmodel.NewGetAllProjectsOutput(nil), getallerr.ErrProjectsNotFound
		}
		log.Warn("cannot get projects", slog.String("error", err.Error()))
		return getallmodel.NewGetAllProjectsOutput(nil), err
	}

	log.Info("projects received")

	return getallmodel.NewGetAllProjectsOutput(projects), nil
}

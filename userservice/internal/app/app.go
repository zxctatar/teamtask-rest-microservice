package app

import (
	"context"
	"database/sql"
	"log/slog"
	"time"
	"userservice/internal/config"
	bcrypthash "userservice/internal/infrastructure/bcrypt"
	"userservice/internal/infrastructure/postgres"
	"userservice/internal/transport/rest"
	resthandler "userservice/internal/transport/rest/handler"
	"userservice/internal/usecase/implementations/registration"
)

type App struct {
	log        *slog.Logger
	restServer *rest.RestServer
	cfg        *config.Config
	db         *sql.DB
}

func NewApp(cfg *config.Config, log *slog.Logger) *App {
	db := mustLoadPostgres(cfg)

	pos := postgres.NewPostgres(db)

	hasher := bcrypthash.NewBcryptHasher()

	regUC := registration.NewRegUC(log, pos, hasher)

	handl := resthandler.NewRestHandler(log, regUC)

	restServer := mustLoadHttpServer(cfg, log, handl)

	return &App{
		log:        log,
		restServer: restServer,
		cfg:        cfg,
		db:			db,
	}
}

func (a *App) Run() {
	a.restServer.MustStart()
}

func (a *App) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.RestConf.ShutdownTimeout*time.Second)
	defer cancel()

	a.restServer.Stop(ctx)

	a.db.Close()
}

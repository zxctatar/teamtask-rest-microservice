package postgres

import (
	"context"
	"database/sql"
	"log/slog"
	userdomain "userservice/internal/domain/user"
)

type Postgres struct {
	log *slog.Logger
	db  *sql.DB
}

func NewPostgres(log *slog.Logger, db *sql.DB) *Postgres {
	return &Postgres{
		log: log,
		db:  db,
	}
}

func (p *Postgres) Save(ctx context.Context, ud *userdomain.UserDomain) (uint32, error) {
	panic("not implemented")
}

func (p *Postgres) FindByEmail(ctx context.Context, email string) (*userdomain.UserDomain, error) {
	panic("not implemented")
}

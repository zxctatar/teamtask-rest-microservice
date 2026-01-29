package app

import (
	"database/sql"
	"fmt"
	"projectservice/internal/config"

	_ "github.com/lib/pq"
)

func mustLoadPostgres(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgresConf.Host,
		cfg.PostgresConf.Port,
		cfg.PostgresConf.User,
		cfg.PostgresConf.Password,
		cfg.PostgresConf.DbName,
		cfg.PostgresConf.Sslmode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("cannot open db: " + err.Error())
	}

	if err := db.Ping(); err != nil {
		panic("cannot ping db:" + err.Error())
	}

	return db
}

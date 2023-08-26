package dbrepo

import (
	"database/sql"

	"gitlab.com/jlamb85/bookings/internal/config"
	"gitlab.com/jlamb85/bookings/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}

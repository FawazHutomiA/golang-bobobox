package app

import (
	"bobobox/pkg/log"

	"github.com/jmoiron/sqlx"
)

type AppConfig struct {
	Db     *sqlx.DB
	Logger log.Logger
}

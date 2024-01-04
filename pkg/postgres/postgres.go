package postgres

import (
	"database/sql"
	"time"
)

const (
	_defaultConnAttempts = 3
	_defaultConnTimeout  = time.Second
)

type postgres struct {
	connAttempts int
	connTimeout  time.Duration

	db *sql.DB
}

func New() {

}

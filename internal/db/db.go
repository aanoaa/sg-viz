package db

import (
	"database/sql"

	"github.com/pkg/errors"
)

func Conn(path string) (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		return nil, errors.Wrap(err, "db open fail")
	}
	return
}

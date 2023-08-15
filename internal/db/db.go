package db

import (
	"database/sql"
	"os"
	"path"

	"github.com/pkg/errors"
)

var DBFilePath = path.Join(os.Getenv("HOME"), ".local", "share", "sgviz", "sqlite3.db")

func Conn() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite3", DBFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "db open fail")
	}
	return
}

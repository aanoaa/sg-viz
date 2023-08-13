package repo

import (
	"context"
	"database/sql"

	v1 "github.com/aanoaa/sg-viz/repo/v1"
)

type Group interface {
	Upsert(ctx context.Context, record []string) error
}

func NewGroupRepo(db *sql.DB) Group {
	return v1.NewGroup(db)
}

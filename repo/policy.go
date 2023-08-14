package repo

import (
	"context"
	"database/sql"

	v1 "github.com/aanoaa/sg-viz/repo/v1"
)

type Policy interface {
	Upsert(ctx context.Context, record []string) error
}

func NewPolicyRepo(db *sql.DB) Policy {
	return v1.NewPolicy(db)
}

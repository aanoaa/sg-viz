package repo

import (
	"context"
	"database/sql"

	"github.com/aanoaa/sg-viz/models"
	v1 "github.com/aanoaa/sg-viz/repo/v1"
)

type Group interface {
	FindByName(ctx context.Context, name string) (*models.Sgroup, error)
	Upsert(ctx context.Context, record []string) error
}

func NewGroupRepo(db *sql.DB) Group {
	return v1.NewGroup(db)
}

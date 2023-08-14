package repo

import (
	"context"
	"database/sql"

	"github.com/aanoaa/sg-viz/models"
	v1 "github.com/aanoaa/sg-viz/repo/v1"
)

type Host interface {
	FindByHostname(ctx context.Context, hostname string) (*models.Host, error)
	Upsert(ctx context.Context, record []string) error
}

func NewHostRepo(db *sql.DB) Host {
	return v1.NewHost(db)
}

type Group interface {
	FindByName(ctx context.Context, name string) (*models.Sgroup, error)
	Upsert(ctx context.Context, record []string) error
}

func NewGroupRepo(db *sql.DB) Group {
	return v1.NewGroup(db)
}

type Policy interface {
	Upsert(ctx context.Context, record []string) error
}

func NewPolicyRepo(db *sql.DB) Policy {
	return v1.NewPolicy(db)
}

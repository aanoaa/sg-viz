package repo

import (
	"context"
	"database/sql"

	"github.com/aanoaa/sgviz/models"
	v1 "github.com/aanoaa/sgviz/repo/v1"
	"github.com/aanoaa/sgviz/types"
)

type Host interface {
	FindByHostname(ctx context.Context, hostname string) (*models.Host, error)
	ListByGroupID(ctx context.Context, groupID int64) (models.HostSlice, error)
	Upsert(ctx context.Context, record []string, zone string) error
	List(ctx context.Context, match string) (models.HostSlice, error)
}

func NewHostRepo(db *sql.DB) Host {
	return v1.NewHost(db)
}

type Group interface {
	FindByName(ctx context.Context, name string) (*models.Sgroup, error)
	Upsert(ctx context.Context, record []string) error
	List(ctx context.Context, match string) (models.SgroupSlice, error)
}

func NewGroupRepo(db *sql.DB) Group {
	return v1.NewGroup(db)
}

type Policy interface {
	ListHost(ctx context.Context) ([]types.PolicyByHost, error)
	ListGroup(ctx context.Context) ([]types.PolicyByGroup, error)
	Upsert(ctx context.Context, record []string) error
}

func NewPolicyRepo(db *sql.DB) Policy {
	return v1.NewPolicy(db)
}

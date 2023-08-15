package v1

import (
	"context"
	"database/sql"

	"github.com/aanoaa/sgviz/models"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Host struct {
	db *sql.DB
}

func NewHost(db *sql.DB) *Host {
	return &Host{db}
}

func (r *Host) FindByHostname(ctx context.Context, hostname string) (*models.Host, error) {
	host, err := models.Hosts(models.HostWhere.Hostname.EQ(hostname)).One(ctx, r.db)
	return host, errors.Wrap(err, "find fail")
}

func (r *Host) Upsert(ctx context.Context, record []string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "begin tx fail")
	}
	defer func() { _ = tx.Rollback() }()

	host := models.Host{
		Hostname: record[0],
		Ipaddr:   record[1],
	}
	if len(record) == 3 {
		host.Desc = record[2]
	}

	if err := host.Upsert(ctx, r.db, true, []string{"hostname"},
		boil.Whitelist("ipaddr", "desc"), boil.Infer()); err != nil {
		return errors.Wrap(err, "upsert fail")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "commit fail")
	}

	return nil
}

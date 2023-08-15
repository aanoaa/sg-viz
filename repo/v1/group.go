package v1

import (
	"context"
	"database/sql"

	"github.com/aanoaa/sgviz/models"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Group struct {
	db *sql.DB
}

func NewGroup(db *sql.DB) *Group {
	return &Group{db}
}

func (r *Group) FindByName(ctx context.Context, name string) (*models.Sgroup, error) {
	group, err := models.Sgroups(models.SgroupWhere.Name.EQ(name)).One(ctx, r.db)
	return group, errors.Wrap(err, "find fail")
}

func (r *Group) Upsert(ctx context.Context, record []string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "begin tx fail")
	}
	defer func() { _ = tx.Rollback() }()

	hr := NewHost(r.db)
	host, err := hr.FindByHostname(ctx, record[1])
	if err != nil {
		return errors.Wrap(err, "find host fail")
	}

	group := &models.Sgroup{
		Name: record[0],
	}

	if err := group.Upsert(ctx, r.db, true, []string{"name"},
		boil.Whitelist("desc"), boil.Infer()); err != nil {
		return errors.Wrap(err, "upsert fail")
	}

	hg := &models.HostSgroup{
		HostID:   host.ID,
		SgroupID: group.ID,
	}

	_, err = models.FindHostSgroup(ctx, r.db, host.ID, group.ID)
	if errors.Is(err, sql.ErrNoRows) {
		if err := hg.Insert(ctx, r.db, boil.Infer()); err != nil {
			return errors.Wrap(err, "insert host_sgroup fail")
		}
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "commit fail")
	}

	return nil
}

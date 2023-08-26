package v1

import (
	"context"
	"database/sql"

	"github.com/aanoaa/sgviz/models"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Group struct {
	db *sql.DB
}

func NewGroup(db *sql.DB) *Group {
	return &Group{db}
}

func (r *Group) FindByName(ctx context.Context, name string) (*models.Sgroup, error) {
	group, err := models.Sgroups(models.SgroupWhere.Name.EQ(name)).One(ctx, r.db)
	return group, errors.Wrapf(err, "find group fail: %s", name)
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
		return errors.Wrapf(err, "find host fail: %s", record[1])
	}

	group := &models.Sgroup{
		Name: record[0],
		Zone: record[2],
	}

	if err := group.Upsert(ctx, r.db, true, []string{"name"},
		boil.Whitelist("ipaddr", "zone", "desc"), boil.Infer()); err != nil {
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

func (r *Group) List(ctx context.Context, match string) (models.SgroupSlice, error) {
	var list models.SgroupSlice
	var err error
	if match == "" {
		list, err = models.Sgroups().All(ctx, r.db)
	} else {
		list, err = models.Sgroups(qm.Where("name LIKE ?", "%"+match+"%")).All(ctx, r.db)
	}
	return list, errors.Wrap(err, "query fail")
}

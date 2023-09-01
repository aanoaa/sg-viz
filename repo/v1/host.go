package v1

import (
	"context"
	"database/sql"

	"github.com/aanoaa/sgviz/models"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type Host struct {
	db *sql.DB
}

func NewHost(db *sql.DB) *Host {
	return &Host{db}
}

func (r *Host) FindByHostname(ctx context.Context, hostname string) (*models.Host, error) {
	host, err := models.Hosts(models.HostWhere.Hostname.EQ(hostname)).One(ctx, r.db)
	return host, errors.Wrapf(err, "find fail: %s", hostname)
}

func (r *Host) Upsert(ctx context.Context, record []string, zone string) error {
	var hostname, ipaddr, desc string
	switch len(record) {
	case 1:
		hostname = record[0]
	case 2:
		hostname, ipaddr = record[0], record[1]
	case 3:
		hostname, ipaddr, desc = record[0], record[1], record[2]
	default:
		return errors.Errorf("Unexpected format: %v", record)
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "begin tx fail")
	}
	defer func() { _ = tx.Rollback() }()

	host := models.Host{
		Hostname: hostname,
		Ipaddr:   ipaddr,
		Desc:     desc,
		Zone:     zone,
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

func (r *Host) List(ctx context.Context, match string) (models.HostSlice, error) {
	var list models.HostSlice
	var err error
	if match == "" {
		list, err = models.Hosts().All(ctx, r.db)
	} else {
		list, err = models.Hosts(qm.Where("hostname LIKE ?", "%"+match+"%")).All(ctx, r.db)
	}
	return list, errors.Wrap(err, "query fail")
}

func (r *Host) ListByGroupID(ctx context.Context, groupID int64) (models.HostSlice, error) {
	hostGroups, err := models.HostSgroups(models.HostSgroupWhere.SgroupID.EQ(groupID)).All(ctx, r.db)
	if err != nil {
		return nil, errors.Wrap(err, "host_sgroup query fail")
	}

	var hosts models.HostSlice
	for _, hg := range hostGroups {
		host, err := models.FindHost(ctx, r.db, hg.HostID)
		if err != nil {
			return nil, errors.Wrapf(err, "find host fail: %d", hg.HostID)
		}
		hosts = append(hosts, host)
	}
	return hosts, nil
}

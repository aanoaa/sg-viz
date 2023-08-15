package v1

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"

	"github.com/aanoaa/sg-viz/models"
	"github.com/aanoaa/sg-viz/types"
)

type Policy struct {
	db *sql.DB
}

func NewPolicy(db *sql.DB) *Policy {
	return &Policy{db}
}

func (r *Policy) Upsert(ctx context.Context, record []string) error {
	port, err := strconv.Atoi(record[2])
	if err != nil {
		return errors.Wrap(err, "wrong port")
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "begin tx fail")
	}
	defer func() { _ = tx.Rollback() }()

	gr := NewGroup(r.db)
	src, err := gr.FindByName(ctx, record[0])
	if err != nil {
		return errors.Wrap(err, "find group fail")
	}

	dst, err := gr.FindByName(ctx, record[1])
	if err != nil {
		return errors.Wrap(err, "find group fail")
	}

	policy := &models.Policy{
		SRC:  src.ID,
		DST:  dst.ID,
		Port: int64(port),
	}

	if err := policy.Upsert(ctx, r.db, true, []string{"src", "dst", "port", "protocol"},
		boil.Whitelist("desc"), boil.Infer()); err != nil {
		return errors.Wrap(err, "upsert fail")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "commit fail")
	}

	return nil
}

func (r *Policy) ListHost(ctx context.Context) ([]types.PolicyByHost, error) {
	q := queries.Raw(`
SELECT sh.hostname AS src, sh.ipaddr AS src_addr, dh.hostname AS dst, dh.ipaddr AS dst_addr, p.port, p.protocol
FROM policy p
JOIN sgroup src ON src.id = p.src
JOIN sgroup dst ON dst.id = p.dst
JOIN host_sgroup shg ON shg.sgroup_id = src.id
JOIN host_sgroup dhg ON dhg.sgroup_id = dst.id
JOIN host sh ON sh.id = shg.host_id
JOIN host dh ON dh.id = dhg.host_id`)

	var rs []types.PolicyByHost
	if err := q.Bind(ctx, r.db, &rs); err != nil {
		return nil, errors.Wrap(err, "bind fail")
	}
	return rs, nil
}

func (r *Policy) ListGroup(ctx context.Context) ([]types.PolicyByGroup, error) {
	q := queries.Raw(`
SELECT src.name AS src, dst.name AS dst, p.port, p.protocol
FROM policy p
JOIN sgroup src ON src.id = p.src
JOIN sgroup dst ON dst.id = p.dst
JOIN host_sgroup shg ON shg.sgroup_id = src.id
JOIN host_sgroup dhg ON dhg.sgroup_id = dst.id
JOIN host sh ON sh.id = shg.host_id
JOIN host dh ON dh.id = dhg.host_id
GROUP BY src.name, dst.name, p.port`)

	var rs []types.PolicyByGroup
	if err := q.Bind(ctx, r.db, &rs); err != nil {
		return nil, errors.Wrap(err, "bind fail")
	}
	return rs, nil
}

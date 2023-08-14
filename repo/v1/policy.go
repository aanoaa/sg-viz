package v1

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/aanoaa/sg-viz/models"
	"github.com/pkg/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Policy struct {
	db *sql.DB
}

func NewPolicy(db *sql.DB) *Policy {
	return &Policy{db}
}

func (r *Policy) Upsert(ctx context.Context, record []string) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return errors.Wrap(err, "begin tx fail")
	}
	defer func() { _ = tx.Rollback() }()

	port, err := strconv.Atoi(record[2])
	if err != nil {
		return errors.Wrap(err, "wrong port")
	}

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
		From: src.ID,
		To:   dst.ID,
		Port: int64(port),
	}

	if err := policy.Upsert(ctx, r.db, true, []string{"id"},
		boil.Whitelist("from", "to", "port", "protocol", "desc"), boil.Infer()); err != nil {
		return errors.Wrap(err, "upsert fail")
	}

	if err := tx.Commit(); err != nil {
		return errors.Wrap(err, "commit fail")
	}

	return nil
}

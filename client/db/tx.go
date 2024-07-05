package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/techx/portal/errors"
)

type TxRunner interface {
	RunInTxContext(ctx context.Context, fn func(*sqlx.Tx) error) error
}

type txRunner struct {
	db *sqlx.DB
}

func newTxRunner(db *sqlx.DB) *txRunner {
	return &txRunner{db: db}
}

func (r *txRunner) RunInTx(fn func(*sqlx.Tx) error) (err error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return err
	}
	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			err = errors.New(fmt.Sprint(r))
		}
	}()
	if err = fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (r *txRunner) RunInTxContext(ctx context.Context, fn func(*sqlx.Tx) error) (err error) {
	tx, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			err = errors.New(fmt.Sprint(r))
		}
	}()

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

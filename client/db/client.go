package db

import (
	"context"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/techx/portal/config"
	"github.com/techx/portal/db"
	"github.com/techx/portal/errors"
)

var ErrZeroRowsAffected = errors.New("no rows affected")

type Repository struct {
	db        *sqlx.DB
	appName   string
	tableName string
	TxRunner  TxRunner
}

func NewRepository(cfg *config.Config, tableName string) *Repository {
	postgresDB, err := db.NewPostgresDB(context.Background(), cfg.DB)
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to postgres")
		panic(err)
	}

	return &Repository{
		appName:   cfg.AppName,
		tableName: tableName,
		db:        postgresDB,
		TxRunner:  NewTxRunner(postgresDB),
	}
}

func (r *Repository) DBGet(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return r.db.GetContext(ctx, dest, query, args...)
}

func (r *Repository) DBGetInTx(ctx context.Context, tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBGet(ctx, dest, query, args...)
	}
	return tx.GetContext(ctx, dest, query, args...)
}

func (r *Repository) DBSelect(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return r.db.SelectContext(ctx, dest, query, args...)
}

func (r *Repository) DBSelectInTx(ctx context.Context, tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBSelect(ctx, dest, query, args...)
	}
	return tx.SelectContext(ctx, dest, query, args...)
}

func (r *Repository) DBExec(ctx context.Context, query string, args ...interface{}) error {
	return r.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.DBExecInTx(ctx, tx, query, args...)
	})
}

func (r *Repository) DBExecInTx(ctx context.Context, tx *sqlx.Tx, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBExec(ctx, query, args...)
	}
	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if rowsAffected == int64(0) {
		return ErrZeroRowsAffected
	}
	return err
}

func (r *Repository) DBSoftExec(ctx context.Context, query string, args ...interface{}) error {
	return r.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.DBSoftExecInTx(ctx, tx, query, args...)
	})
}

func (r *Repository) DBSoftExecInTx(ctx context.Context, tx *sqlx.Tx, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBSoftExec(ctx, query, args...)
	}
	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	return err
}

func (r *Repository) DBNamedExec(ctx context.Context, query string, arg interface{}) error {
	return r.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.DBNamedExecInTx(ctx, tx, query, arg)
	})
}

func (r *Repository) DBNamedExecInTx(ctx context.Context, tx *sqlx.Tx, query string, arg interface{}) error {
	if tx == nil {
		return r.DBNamedExec(ctx, query, arg)
	}

	res, err := tx.NamedExecContext(ctx, query, arg)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if rowsAffected == int64(0) {
		return ErrZeroRowsAffected
	}

	return err
}

func (r *Repository) DBExecReturning(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return r.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.DBExecReturningInTx(ctx, tx, dest, query, args...)
	})
}

func (r *Repository) DBExecReturningInTx(ctx context.Context, tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBExecReturning(ctx, dest, query, args...)
	}
	row := tx.QueryRowxContext(ctx, query, args...)
	return row.StructScan(dest)
}

func (r *Repository) DBNamedExecReturning(ctx context.Context, dest interface{}, query string, arg interface{}) error {
	return r.TxRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.DBNamedExecReturningInTx(ctx, tx, dest, query, arg)
	})
}

func (r *Repository) DBNamedExecReturningInTx(ctx context.Context, tx *sqlx.Tx, dest interface{}, query string, arg interface{}) error {
	if tx == nil {
		return r.DBNamedExecReturning(ctx, dest, query, arg)
	}

	query, args, err := r.db.BindNamed(query, arg)
	if err != nil {
		return err
	}

	row := tx.QueryRowxContext(ctx, query, args...)
	return row.StructScan(dest)
}

func (r *Repository) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	return r.db.BindNamed(query, arg)
}

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

type Client interface {
	DBGet(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	DBGetInTx(ctx context.Context, tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) error
	DBSelect(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	DBSelectInTx(ctx context.Context, tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) error
	DBExec(ctx context.Context, query string, args ...interface{}) error
	DBExecInTx(ctx context.Context, tx *sqlx.Tx, query string, args ...interface{}) error
	DBSoftExec(ctx context.Context, query string, args ...interface{}) error
	DBSoftExecInTx(ctx context.Context, tx *sqlx.Tx, query string, args ...interface{}) error
	DBNamedExec(ctx context.Context, query string, arg interface{}) error
	DBNamedExecInTx(ctx context.Context, tx *sqlx.Tx, query string, arg interface{}) error
	DBExecReturning(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	DBExecReturningInTx(ctx context.Context, tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) error
	DBNamedExecReturning(ctx context.Context, dest interface{}, query string, arg interface{}) error
	DBNamedExecReturningInTx(ctx context.Context, tx *sqlx.Tx, dest interface{}, query string, arg interface{}) error
	DBRunInTxContext(ctx context.Context, fn func(tx *sqlx.Tx) error) error
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
}

type dbClient struct {
	db       *sqlx.DB
	appName  string
	txRunner TxRunner
}

func NewPostgresDBClient(cfg *config.Config) Client {
	postgresDB, err := db.NewPostgresDB(context.Background(), cfg.DB)
	if err != nil {
		log.Error().Err(err).Msg("failed to connect to postgres")
		panic(err)
	}

	return &dbClient{
		appName:  cfg.AppName,
		db:       postgresDB,
		txRunner: newTxRunner(postgresDB),
	}
}

func (r *dbClient) DBGet(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return r.db.GetContext(ctx, dest, query, args...)
}

func (r *dbClient) DBGetInTx(ctx context.Context, tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBGet(ctx, dest, query, args...)
	}
	return tx.GetContext(ctx, dest, query, args...)
}

func (r *dbClient) DBSelect(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return r.db.SelectContext(ctx, dest, query, args...)
}

func (r *dbClient) DBSelectInTx(ctx context.Context, tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBSelect(ctx, dest, query, args...)
	}
	return tx.SelectContext(ctx, dest, query, args...)
}

func (r *dbClient) DBExec(ctx context.Context, query string, args ...interface{}) error {
	return r.txRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.DBExecInTx(ctx, tx, query, args...)
	})
}

func (r *dbClient) DBExecInTx(ctx context.Context, tx *sqlx.Tx, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBExec(ctx, query, args...)
	}
	res, err := tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if rowsAffected == int64(0) {
		return errors.ErrZeroRowsAffected
	}
	return err
}

func (r *dbClient) DBSoftExec(ctx context.Context, query string, args ...interface{}) error {
	return r.txRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.DBSoftExecInTx(ctx, tx, query, args...)
	})
}

func (r *dbClient) DBSoftExecInTx(ctx context.Context, tx *sqlx.Tx, query string, args ...interface{}) error {
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

func (r *dbClient) DBNamedExec(ctx context.Context, query string, arg interface{}) error {
	return r.txRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.DBNamedExecInTx(ctx, tx, query, arg)
	})
}

func (r *dbClient) DBNamedExecInTx(ctx context.Context, tx *sqlx.Tx, query string, arg interface{}) error {
	if tx == nil {
		return r.DBNamedExec(ctx, query, arg)
	}

	res, err := tx.NamedExecContext(ctx, query, arg)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if rowsAffected == int64(0) {
		return errors.ErrZeroRowsAffected
	}

	return err
}

func (r *dbClient) DBExecReturning(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return r.txRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.DBExecReturningInTx(ctx, tx, dest, query, args...)
	})
}

func (r *dbClient) DBExecReturningInTx(ctx context.Context, tx *sqlx.Tx, dest interface{}, query string, args ...interface{}) error {
	if tx == nil {
		return r.DBExecReturning(ctx, dest, query, args...)
	}
	row := tx.QueryRowxContext(ctx, query, args...)
	return row.StructScan(dest)
}

func (r *dbClient) DBNamedExecReturning(ctx context.Context, dest interface{}, query string, arg interface{}) error {
	return r.txRunner.RunInTxContext(ctx, func(tx *sqlx.Tx) error {
		return r.DBNamedExecReturningInTx(ctx, tx, dest, query, arg)
	})
}

func (r *dbClient) DBNamedExecReturningInTx(ctx context.Context, tx *sqlx.Tx, dest interface{}, query string, arg interface{}) error {
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

func (r *dbClient) DBRunInTxContext(ctx context.Context, fn func(tx *sqlx.Tx) error) error {
	return r.txRunner.RunInTxContext(ctx, fn)
}

func (r *dbClient) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	return r.db.BindNamed(query, arg)
}

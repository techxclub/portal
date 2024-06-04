package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/techx/portal/config"
)

func NewPostgresDB(ctx context.Context, dbConfig config.DB) (*sqlx.DB, error) {
	connStr := dbConfig.GetConnectionString()

	// Set up connection pool configuration
	poolConfig, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to parse connection string")
	}

	// Create connection pool
	pool, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to create connection pool")
	}

	// Register pgxpool with database/sql
	sqlDB := stdlib.OpenDB(*poolConfig.ConnConfig)

	// Wrap the sql.DB with sqlx
	db := sqlx.NewDb(sqlDB, "pgx")
	db.SetMaxOpenConns(dbConfig.GetMaxPoolSize())
	if maxIdleConn := dbConfig.GetMaxIdleConnections(); maxIdleConn != 0 {
		db.SetMaxIdleConns(maxIdleConn)
	}

	if maxLifeTime := dbConfig.GetConnectionMaxLifeTime(); maxLifeTime != 0 {
		db.SetConnMaxLifetime(maxLifeTime)
	}

	if maxIdleTime := dbConfig.GetConnectionMaxIdleTime(); maxIdleTime != 0 {
		db.SetConnMaxIdleTime(maxIdleTime)
	}

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := pool.Ping(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to ping database")
	}

	log.Info().Msg("Successfully connected to the database")

	return db, nil
}

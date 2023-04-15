package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type PgConfig struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

func NewPgConfig(username, password, host, port, database string) *PgConfig {
	return &PgConfig{
		Username: username,
		Password: password,
		Host:     host,
		Port:     port,
		Database: database,
	}
}

func NewClient(ctx context.Context, maxAttempts int, maxDelay time.Duration, cfg *PgConfig) (pool *pgxpool.Pool, err error) {
	var dbPool *pgxpool.Pool
	var dbConnectErr error
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)

	for maxAttempts > 0 {
		dbPool, dbConnectErr = initPostgres(ctx, dsn)
		if dbConnectErr != nil {
			time.Sleep(maxDelay)
			maxAttempts--
			continue
		}
		break
	}

	if dbConnectErr != nil {
		log.Fatal("Unable to connect to database")
	}

	return dbPool, dbConnectErr
}

func initPostgres(ctx context.Context, dsn string) (pool *pgxpool.Pool, err error) {

	pgxConfig, errPgConfig := pgxpool.ParseConfig(dsn)
	if errPgConfig != nil {
		log.Fatal("unable config pgx")
	}
	dbPool, dbConnectErr := pgxpool.NewWithConfig(ctx, pgxConfig)

	return dbPool, dbConnectErr
}

package postgres

import (
	"context"
	"fmt"

	"github.com/Noiidor/go-service-template/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const driver = "pgx"

func New(ctx context.Context, cfg *config.Config) (*sqlx.DB, error) {
	dbUrl := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s",
		cfg.GetDbUser(),
		cfg.GetDbPass(),
		cfg.GetDbHost(),
		cfg.GetDbPort(),
		cfg.GetDbName(),
	)

	db, err := sqlx.ConnectContext(ctx, driver, dbUrl)
	if err != nil {
		return nil, fmt.Errorf("postgres connect err: %w", err)
	}

	return db, nil
}

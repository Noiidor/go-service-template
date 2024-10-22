package postgres

import (
	"context"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

const driver = "pgx"

func New(ctx context.Context) (*sqlx.DB, error) {
	db, err := sqlx.ConnectContext(ctx, driver, "postgresql://")
	if err != nil {
		return nil, fmt.Errorf("postgres connect err: %w", err)
	}

	return db, nil
}

package repos

import "github.com/jmoiron/sqlx"

type ExampleRepo interface {
}

type exampleRepo struct {
	db *sqlx.DB
}

func NewExampleRepo(db *sqlx.DB) exampleRepo {
	return exampleRepo{
		db: db,
	}
}

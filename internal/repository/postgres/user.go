package postgres

import "github.com/jmoiron/sqlx"

type userRepo struct {
	db *sqlx.DB
}

func newUserRepo(db *sqlx.DB) UserRepository {
	return &userRepo{db: db}
}

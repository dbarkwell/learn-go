package user

import "github.com/jmoiron/sqlx"

type Repository struct {
	DB *sqlx.DB
}

func ProvideUserRepository(DB *sqlx.DB) Repository {
	return Repository{DB: DB}
}

package authentication

import (
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	DB *sqlx.DB
}

func ProvideAuthenticationRepository(DB *sqlx.DB) Repository {
	return Repository{DB: DB}
}

//func (ar *Repository) SaveCredential()

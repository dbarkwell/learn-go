package user

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	DB *sqlx.DB
}

func ProvideUserRepository(DB *sqlx.DB) Repository {
	return Repository{DB: DB}
}

func (r *Repository) Add(username string, first string, last string, email string) (User, error) {
	id := uuid.New()
	insert := `INSERT INTO user (id, username, first_name, last_name, email) VALUES (UUID_TO_BIN(?), ?, ?, ?, ?)`
	result := r.DB.MustExec(insert, id, username, first, last, email)

	user := User{Username: username, FirstName: first, LastName: last, Email: email}
	if rows, err := result.RowsAffected(); err != nil || rows == 0 {
		return user, err
	}

	user.ID = id

	return user, nil
}

func (r *Repository) Get(username string) (User, error) {
	sel := `SELECT BIN_TO_UUID(id) "id", username, first_name, last_name, email FROM user WHERE username = ?`
	stmt, err := r.DB.Preparex(sel)
	if err != nil {
		return User{}, err
	}

	defer stmt.Close()

	var user User
	err = stmt.Get(&user, username)

	return user, nil
}

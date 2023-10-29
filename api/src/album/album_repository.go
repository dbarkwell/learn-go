package album

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	DB *sqlx.DB
}

func ProvideAlbumRepository(DB *sqlx.DB) Repository {
	return Repository{DB: DB}
}

func (ar *Repository) FindAll() ([]Album, error) {
	var albums []Album
	err := ar.DB.Select(&albums, "SELECT id, title, artist, price  FROM album")
	if err != nil {
		return nil, err
	}

	return albums, nil
}

func (ar *Repository) FindByID(id int64) (Album, error) {
	stmt, err := ar.DB.Preparex(`SELECT id, title, artist, price FROM album WHERE id = ?`)
	if err != nil {
		return Album{}, err
	}

	var album Album
	err = stmt.Get(&album, id)

	return album, err
}

func (ar *Repository) Add(a Album) (Album, error) {
	insert := `INSERT INTO album (title, artist, price) VALUES (?, ?, ?)`
	result := ar.DB.MustExec(insert, a.Title, a.Artist, a.Price)
	id, err := result.LastInsertId()

	if err != nil {
		return Album{}, err
	}

	a.ID = id
	return a, nil
}

func (ar *Repository) Remove(id int64) (bool, error) {
	delete := `DELETE FROM album WHERE id = ?`
	result := ar.DB.MustExec(delete, id)
	rows, err := result.RowsAffected()

	if err != nil {
		return false, err
	}

	if rows != 1 {
		return false, nil
	}

	return true, nil
}

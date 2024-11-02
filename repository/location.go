package repository

import (
	"database/sql"
	"project/model"
)

type Location struct {
	Db *sql.DB
}

func CreateLocation(db *sql.DB) *Location {
	return &Location{Db: db}
}

func (repo *Location) Create(location *model.Location) error {
	query := `INSERT INTO locations (name) SELECT $1 FROM sessions WHERE id=$2 AND expired_at IS NULL RETURNING id,''`
	err := repo.Db.QueryRow(query, location.Name, location.SessionId).Scan(&location.Id, &location.SessionId)
	if err != nil {
		return err
	}
	return nil
}

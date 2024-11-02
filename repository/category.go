package repository

import (
	"database/sql"
	"project/model"
)

type Category struct {
	Db *sql.DB
}

func CreateCategory(db *sql.DB) *Category {
	return &Category{Db: db}
}

func (repo *Category) Create(category *model.Category) error {
	query := `INSERT INTO categories (name) SELECT $1 FROM sessions WHERE id=$2 AND expired_at IS NULL RETURNING id,''`
	err := repo.Db.QueryRow(query, category.Name, category.SessionId).Scan(&category.Id, &category.SessionId)
	if err != nil {
		return err
	}
	return nil
}

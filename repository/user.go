package repository

import (
	"database/sql"
	"project/model"
)

type User struct {
	Db *sql.DB
}

func CreateUser(db *sql.DB) *User {
	return &User{Db: db}
}

func (r *User) Create(user *model.User) error {
	query := `INSERT INTO users (username,password) VALUES ($1, $2) RETURNING id`
	err := r.Db.QueryRow(query, user.Username, user.Password).Scan(&user.Id)
	if err != nil {
		return err
	}
	return nil
}

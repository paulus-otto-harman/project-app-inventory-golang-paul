package repository

import (
	"database/sql"
	"project/model"
)

type Auth struct {
	Db *sql.DB
}

func CreateAuth(db *sql.DB) *Auth {
	return &Auth{Db: db}
}

func (repo *Auth) Authenticate(user *model.User) (*model.Session, error) {
	query := `INSERT INTO sessions (user_id) SELECT id FROM users WHERE username=$1 AND password=$2 RETURNING id`
	var session model.Session

	if err := repo.Db.QueryRow(query, user.Username, user.Password).Scan(&session.Id); err != nil {
		return nil, err
	}
	return &session, nil
}

func (repo *Auth) Logout(session *model.Session) (int, error) {
	query := `UPDATE sessions SET expired_at=now() WHERE id=$1`
	result, err := repo.Db.Exec(query, session.Id)
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	return int(rowsAffected), nil
}

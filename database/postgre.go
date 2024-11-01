package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const DbUser = "postgres"
const DbPassword = "postgres"

func PgConnect(dbName string) (*sql.DB, error) {
	return sql.Open("postgres", fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=localhost", DbUser, DbPassword, dbName))
}

func DbOpen(dbName string) *sql.DB {
	db, err := PgConnect(dbName)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

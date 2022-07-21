package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

const ENGINE_DATABASE = "postgres"
const URL_DATABASE = "postgresql://root:secret@localhost:5432/prueba?sslmode=disable"

func GetConnection() (*sql.DB, error) {
	conn, err := sql.Open(ENGINE_DATABASE, URL_DATABASE)
	if err != nil {
		return nil, err
	}
	return conn, err
}

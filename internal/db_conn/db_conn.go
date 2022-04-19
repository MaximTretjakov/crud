package db_conn

import (
	"database/sql"
)

func NewDBConnection(driver string, connectionString string) (*sql.DB, error) {
	// create database connection
	db, err := sql.Open(driver, connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

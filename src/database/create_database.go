package database

import (
	"database/sql"
	"fmt"
)

func CreateDatabaseIfNotExists(driver, dsn, dbName string) error {
	var db *sql.DB
	var err error

	if driver == "mysql" {
		db, err = sql.Open("mysql", dsn)
	} else if driver == "postgres" {
		db, err = sql.Open("postgres", dsn)
	}

	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	return err
}

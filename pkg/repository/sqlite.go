package repository

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

const (
	userTable  = "user"
	phoneTable = "phone"
)

func NewSQLiteDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./sarkor.db")
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	createSchema(db)

	return db, nil
}

func CloseDB(db *sql.DB) {
	db.Close()
}

func createSchema(db *sql.DB) {
	statement, _ := db.Prepare(
		"CREATE TABLE IF NOT EXISTS user " +
			"(id INTEGER PRIMARY KEY, login TEXT, password_hash TEXT, name TEXT, age INTEGER)" +
			"")
	statement.Exec()

	statement, _ = db.Prepare(
		"CREATE TABLE IF NOT EXISTS phone " +
			"(id INTEGER PRIMARY KEY, phone TEXT, description TEXT, isFax TEXT, userId INTEGER)" +
			"")
	statement.Exec()
}

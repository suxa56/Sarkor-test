package repository

import (
	Sarkor_test "Sarkor-test"
	"database/sql"
	"fmt"
)

type AuthSQLite struct {
	db *sql.DB
}

func NewAuthSQLite(db *sql.DB) *AuthSQLite {
	return &AuthSQLite{db: db}
}

func (a *AuthSQLite) CreateUser(user Sarkor_test.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (login, password_hash, name, age) VALUES ($1, $2, $3, $4) RETURNING id",
		userTable)
	row := a.db.QueryRow(query, user.Login, user.Password, user.Name, user.Age)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

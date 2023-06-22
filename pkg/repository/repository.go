package repository

import (
	Sarkor_test "Sarkor-test"
	"database/sql"
)

type Authorization interface {
	CreateUser(user Sarkor_test.User) (int, error)
	GetUser(login, password string) (int, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSQLite(db),
	}
}

package repository

import (
	Sarkor_test "Sarkor-test"
	"database/sql"
)

type Authorization interface {
	CreateUser(user Sarkor_test.User) (int, error)
	GetUser(login, password string) (int, error)
}

type UserInfo interface {
	GetUserInfo(name string) ([]Sarkor_test.UserDto, error)
}

type Repository struct {
	Authorization
	UserInfo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSQLite(db),
		UserInfo:      NewUserInfoRepo(db),
	}
}

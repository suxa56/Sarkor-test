package repository

import (
	Sarkor_test "Sarkor-test"
	"Sarkor-test/pkg/repository/user"
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
		Authorization: user.NewAuthSQLite(db),
		UserInfo:      user.NewUserInfoRepo(db),
	}
}

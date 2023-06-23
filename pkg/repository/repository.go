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

type PhoneRepo interface {
	CreatePhone(phone Sarkor_test.Phone) (int, error)
	GetPhoneInfo(phone string) (Sarkor_test.PhoneDto, error)
	UpdatePhone(userId int, input Sarkor_test.UpdatePhone) error
	DeletePhone(phoneId, userId int) error
}

type Repository struct {
	Authorization
	UserInfo
	PhoneRepo
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthSQLite(db),
		UserInfo:      NewUserInfoRepo(db),
		PhoneRepo:     NewPhoneRepoImpl(db),
	}
}

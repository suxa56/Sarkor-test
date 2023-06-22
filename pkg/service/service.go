package service

import (
	Sarkor_test "Sarkor-test"
	"Sarkor-test/pkg/repository"
	"Sarkor-test/pkg/service/user"
)

type Authorization interface {
	CreateUser(user Sarkor_test.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type UserInfo interface {
	GetUserInfo(name string) ([]Sarkor_test.UserDto, error)
}

type Service struct {
	Authorization
	UserInfo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: user.NewAuthService(repos.Authorization),
		UserInfo:      user.NewUserInfoService(repos.UserInfo),
	}
}

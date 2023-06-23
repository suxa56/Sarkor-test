package service

import (
	Sarkor_test "Sarkor-test"
	"Sarkor-test/pkg/repository"
)

type Authorization interface {
	CreateUser(user Sarkor_test.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}

type UserInfo interface {
	GetUserInfo(name string) ([]Sarkor_test.UserDto, error)
}

type PhoneService interface {
	CreatePhone(phone Sarkor_test.Phone) (int, error)
	GetPhoneInfo(phone string) (Sarkor_test.PhoneDto, error)
	DeletePhone(phoneId, userId int) error
}

type Service struct {
	Authorization
	UserInfo
	PhoneService
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		UserInfo:      NewUserInfoService(repos.UserInfo),
		PhoneService:  NewPhoneServiceImpl(repos.PhoneRepo),
	}
}

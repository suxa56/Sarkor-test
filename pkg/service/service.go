package service

import (
	Sarkor_test "Sarkor-test"
	"Sarkor-test/pkg/repository"
)

type Authorization interface {
	CreateUser(user Sarkor_test.User) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}

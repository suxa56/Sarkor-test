package service

import (
	Sarkor_test "Sarkor-test"
	"Sarkor-test/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

const (
	salt = "sjhfedvbwuifovbwiujhfgvb"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user Sarkor_test.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	return string(hash)
}

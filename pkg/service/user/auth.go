package user

import (
	Sarkor_test "Sarkor-test"
	"Sarkor-test/pkg/repository"
	"errors"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	salt        = "sjhfedvbwuifovbwiujhfgvb"
	tokenExpire = 12 * time.Hour
	signedKey   = "skjfbhweiovjnsfgkjhgwritlg"
)

type AuthService struct {
	repo repository.Authorization
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId int    `json:"userId"`
	Login  string `json:"login"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) CreateUser(user Sarkor_test.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

func (a *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid sign in method")
		}
		return []byte(signedKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

func (a *AuthService) GenerateToken(login, password string) (string, error) {
	id, err := a.repo.GetUser(login, password+salt)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenExpire).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		id,
		login,
	})

	return token.SignedString([]byte(signedKey))
}

func generatePasswordHash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	return string(hash)
}

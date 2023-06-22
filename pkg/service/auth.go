package service

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

// Create new user with hashed password
// Return user id, error
func (a *AuthService) CreateUser(user Sarkor_test.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return a.repo.CreateUser(user)
}

// Parse JWT token to get user id
// Return user id, error
func (a *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return Sarkor_test.UNDEFINED_ID, errors.New("invalid sign in method")
		}
		return []byte(signedKey), nil
	})

	if err != nil {
		return Sarkor_test.UNDEFINED_ID, err
	}

	// Get custom claims, contains user id
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return Sarkor_test.UNDEFINED_ID, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserId, nil
}

// Generate JWT token to authorized user
// Return JWT token, error
func (a *AuthService) GenerateToken(login, password string) (string, error) {
	// Get authorized user id
	id, err := a.repo.GetUser(login, password+salt)
	if err != nil {
		return "", err
	}

	// Generate token with custom claims
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

// Hash password with bCrypt, salt with default cost (10)
// Return hash password
func generatePasswordHash(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	return string(hash)
}

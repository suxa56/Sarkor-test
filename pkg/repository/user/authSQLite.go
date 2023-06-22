package user

import (
	Sarkor_test "Sarkor-test"
	"Sarkor-test/pkg/repository"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AuthSQLite struct {
	db *sql.DB
}

func NewAuthSQLite(db *sql.DB) *AuthSQLite {
	return &AuthSQLite{db: db}
}

func (a *AuthSQLite) CreateUser(user Sarkor_test.User) (int, error) {
	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (login, password_hash, name, age) VALUES ($1, $2, $3, $4) RETURNING id",
		repository.UserTable)
	row := a.db.QueryRow(query, user.Login, user.Password, user.Name, user.Age)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthSQLite) GetUser(login, password string) (int, error) {
	var id int
	var passHash string
	query := fmt.Sprintf(
		"SELECT id, password_hash FROM %s WHERE login=$1", repository.UserTable)
	row := a.db.QueryRow(query, login)
	if err := row.Scan(&id, &passHash); err != nil {
		return -1, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(passHash), []byte(password)); err != nil {
		return 0, err
	}

	return id, nil
}

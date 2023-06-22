package repository

import (
	Sarkor_test "Sarkor-test"
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

// Create user
// Return user id, error
func (a *AuthSQLite) CreateUser(user Sarkor_test.User) (int, error) {
	var id int
	// Create query
	query := fmt.Sprintf(
		"INSERT INTO %s (login, password_hash, name, age) VALUES ($1, $2, $3, $4) RETURNING id",
		UserTable)
	// Fill query and execute
	row := a.db.QueryRow(query, user.Login, user.Password, user.Name, user.Age)
	// Get id from response
	if err := row.Scan(&id); err != nil {
		return Sarkor_test.UNDEFINED_ID, err
	}
	return id, nil
}

// Get user
// Return user id, error
func (a *AuthSQLite) GetUser(login, password string) (int, error) {
	var id int
	var passHash string
	// Create query -> find user with login
	query := fmt.Sprintf(
		"SELECT id, password_hash FROM %s WHERE login=$1", UserTable)
	// Fill query and execute
	row := a.db.QueryRow(query, login)
	// Get id and password hash from response
	if err := row.Scan(&id, &passHash); err != nil {
		return Sarkor_test.UNDEFINED_ID, err
	}
	// Compare input password with password from db
	if err := bcrypt.CompareHashAndPassword([]byte(passHash), []byte(password)); err != nil {
		return Sarkor_test.UNDEFINED_ID, err
	}

	return id, nil
}

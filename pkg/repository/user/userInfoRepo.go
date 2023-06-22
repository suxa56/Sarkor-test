package user

import (
	Sarkor_test "Sarkor-test"
	"Sarkor-test/pkg/repository"
	"database/sql"
	"fmt"
)

type UserInfoRepo struct {
	db *sql.DB
}

func NewUserInfoRepo(db *sql.DB) *UserInfoRepo {
	return &UserInfoRepo{db: db}
}

// Get user dto
// Return user dto slice, error
func (u *UserInfoRepo) GetUserInfo(name string) ([]Sarkor_test.UserDto, error) {
	var id int
	var age int
	var result = make([]Sarkor_test.UserDto, 0)
	// Create query to get id and age from users having NAME
	query := fmt.Sprintf(
		"SELECT id, age FROM %s WHERE name=$1", repository.UserTable)
	// Fill query and execute
	rows, _ := u.db.Query(query, name)

	// Iterate responses, create dto entity and append to slice
	for rows.Next() {
		err := rows.Scan(&id, &age)
		if err != nil {
			return result, err
		}
		result = append(result, Sarkor_test.UserDto{Id: id, Name: name, Age: age})
	}
	return result, nil
}

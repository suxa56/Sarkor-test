package repository

import (
	Sarkor_test "Sarkor-test"
	"database/sql"
	"fmt"
)

type PhoneRepoImpl struct {
	db *sql.DB
}

func NewPhoneRepoImpl(db *sql.DB) *PhoneRepoImpl {
	return &PhoneRepoImpl{db: db}
}

// Create new phone
// Return phone id, error
func (p *PhoneRepoImpl) CreatePhone(phone Sarkor_test.Phone) (int, error) {
	var id int
	// Create query
	query := fmt.Sprintf(
		"INSERT INTO %s (phone, description, isFax, userId) VALUES ($1, $2, $3, $4) RETURNING id",
		PhoneTable)

	// Fill query and execute
	row := p.db.QueryRow(query, phone.Phone, phone.Description, phone.IsFax, phone.UserId)
	if err := row.Scan(&id); err != nil {
		return Sarkor_test.UNDEFINED_ID, err
	}
	return id, nil
}

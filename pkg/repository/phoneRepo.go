package repository

import (
	Sarkor_test "Sarkor-test"
	"database/sql"
	"errors"
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
	var id = Sarkor_test.UNDEFINED_ID
	// Check phone duplicate
	query := fmt.Sprintf(
		"SELECT id FROM %s WHERE phone=$1",
		PhoneTable)
	row := p.db.QueryRow(query, phone.Phone)
	if err := row.Scan(&id); err != nil && id != Sarkor_test.UNDEFINED_ID {
		return id, err
	} else if id != Sarkor_test.UNDEFINED_ID {
		return id, errors.New("phone duplicate")
	}

	//Create query
	query = fmt.Sprintf(
		"INSERT INTO %s (phone, description, isFax, userId) VALUES ($1, $2, $3, $4) RETURNING id",
		PhoneTable)

	//Fill query and execute
	row = p.db.QueryRow(query, phone.Phone, phone.Description, phone.IsFax, phone.UserId)
	if err := row.Scan(&id); err != nil {
		return Sarkor_test.UNDEFINED_ID, err
	}
	return id, nil
}

// Get phone dto (id, description, isFax, user id) by phone
// Return phone dto, error
func (p *PhoneRepoImpl) GetPhoneInfo(phone string) (Sarkor_test.PhoneDto, error) {
	var id int
	var description string
	var isFax bool
	var userId int

	query := fmt.Sprintf(
		"SELECT id, description, isFax, userId FROM %s WHERE phone=$1", PhoneTable)
	row := p.db.QueryRow(query, phone)
	if err := row.Scan(&id, &description, &isFax, &userId); err != nil {
		return Sarkor_test.PhoneDto{}, err
	}
	return Sarkor_test.PhoneDto{Id: id, Description: description, IsFax: isFax, UserId: userId}, nil
}

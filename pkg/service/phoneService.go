package service

import (
	Sarkor_test "Sarkor-test"
	"Sarkor-test/pkg/repository"
)

type PhoneServiceImpl struct {
	repo repository.PhoneRepo
}

func NewPhoneServiceImpl(repo repository.PhoneRepo) *PhoneServiceImpl {
	return &PhoneServiceImpl{repo: repo}
}

// Create new phone
func (p *PhoneServiceImpl) CreatePhone(phone Sarkor_test.Phone) (int, error) {
	phoneId, err := p.repo.CreatePhone(phone)
	if err != nil {
		return Sarkor_test.UNDEFINED_ID, err
	}
	return phoneId, nil
}

package service

import (
	Sarkor_test "Sarkor-test"
	"Sarkor-test/pkg/repository"
	"errors"
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

// Get phone dto by phone
func (p *PhoneServiceImpl) GetPhoneInfo(phone string) (Sarkor_test.PhoneDto, error) {
	phoneDto, err := p.repo.GetPhoneInfo(phone)
	if err != nil {
		return phoneDto, err
	}
	return phoneDto, nil
}

// Update phone
func (p *PhoneServiceImpl) EditPhone(userId int, input Sarkor_test.UpdatePhone) error {
	if input.Phone == nil && input.IsFax == nil && input.Description == nil {
		return errors.New("empty body")
	}
	return p.repo.UpdatePhone(userId, input)
}

// Delete phone by phone and user id
func (p *PhoneServiceImpl) DeletePhone(phoneId, userId int) error {
	return p.repo.DeletePhone(phoneId, userId)
}

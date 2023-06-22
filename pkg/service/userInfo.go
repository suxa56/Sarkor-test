package service

import (
	Sarkor_test "Sarkor-test"
	"Sarkor-test/pkg/repository"
)

type UserInfoService struct {
	repo repository.UserInfo
}

func NewUserInfoService(repo repository.UserInfo) *UserInfoService {
	return &UserInfoService{repo: repo}
}

// Get users by name
// Return user dto slice, error
func (u *UserInfoService) GetUserInfo(name string) ([]Sarkor_test.UserDto, error) {
	var emptySlice = make([]Sarkor_test.UserDto, 0)
	result, err := u.repo.GetUserInfo(name)
	if err != nil {
		return emptySlice, err
	}

	return result, nil
}

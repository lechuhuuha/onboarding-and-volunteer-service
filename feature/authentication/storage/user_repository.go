package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/domain"
)

func GetUserByEmail(email string, password string) (*domain.User, string) {
	var user domain.User
	err := DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err.Error()
	}
	if user.Status == 0 {
		return nil, "User is inactive"
	}
	if user.Password != password {
		return nil, "Password is incorrect"
	}
	return &user, ""
}

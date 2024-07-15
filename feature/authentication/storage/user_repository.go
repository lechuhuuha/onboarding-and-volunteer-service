package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/domain"
	"gorm.io/gorm"
)

type AuthenticationRepository struct {
	DB *gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) *AuthenticationRepository {
	return &AuthenticationRepository{DB: db}
}
func (r *AuthenticationRepository) GetUserByEmail(email string, password string) (*domain.User, string) {
	var user domain.User
	err := r.DB.Where("email = ?", email).First(&user).Error
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

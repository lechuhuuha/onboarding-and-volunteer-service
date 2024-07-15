package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"gorm.io/gorm"
)

type ApplicantRepository struct {
	DB *gorm.DB
}

func NewApplicantRepository(db *gorm.DB) *ApplicantRepository {
	return &ApplicantRepository{DB: db}
}

// CRUD với các user là applicant
func (r *ApplicantRepository) CreateUser(user *domain.User) error {
	return DB.Create(user).Error
}

func (r *ApplicantRepository) UpdateUser(user *domain.User) error {
	return DB.Save(user).Error
}

func (r *ApplicantRepository) DeleteUser(userID uint) error {
	return DB.Delete(&domain.User{}, userID).Error
}

package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"gorm.io/gorm"
)

type AdminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{db}
}
func (r *AdminRepository) ApproveRequest(id int) (*domain.Request, error) {
	var request domain.Request
	err := r.db.Where("id = ?", id).First(&request).Error
	if err != nil {
		return nil, err
	}
	return &request, nil
}

package storage

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"gorm.io/gorm"
)

type ApplicantRequestRepository struct {
	DB *gorm.DB
}

func NewApplicantRequestRepository(db *gorm.DB) *ApplicantRequestRepository {
	return &ApplicantRequestRepository{DB: db}
}

// Tạo trong bảng request 1 record khi điền xong application form
func (r *ApplicantRequestRepository) CreateRequest(request *domain.Request) error {
	return DB.Create(request).Error
}

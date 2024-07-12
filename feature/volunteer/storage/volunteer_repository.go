package storage

import (
	"gorm.io/gorm"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/domain"
)

type VolunteerRepository interface {
	Create(volunteer *domain.Volunteer) error
	GetAll() ([]*domain.Volunteer, error)
	GetByID(id uint) (*domain.Volunteer, error)
	Update(volunteer *domain.Volunteer) error
	Delete(id uint) error
}

type volunteerRepository struct {
	db *gorm.DB
}

// Create implements VolunteerRepository.
func (v *volunteerRepository) Create(volunteer *domain.Volunteer) error {
	return v.db.Create(volunteer).Error
}

// GetAll implements VolunteerRepository.
func (v *volunteerRepository) GetAll() ([]*domain.Volunteer, error) {
	var volunteers []*domain.Volunteer
	if err := v.db.Find(&volunteers).Error; err != nil {
        return nil, err
    }
    return volunteers, nil
}

// GetByID implements VolunteerRepository.
func (v *volunteerRepository) GetByID(id uint) (*domain.Volunteer, error) {
	var volunteer *domain.Volunteer
	if error := v.db.First(&volunteer, id).Error; error != nil {
		return nil, error
	}
	return volunteer, nil
}

// Update implements VolunteerRepository.
func (v *volunteerRepository) Update(volunteer *domain.Volunteer) error {
	return v.db.Save(volunteer).Error
}

// Delete implements VolunteerRepository.
func (v *volunteerRepository) Delete(id uint) error {
	return v.db.Delete(&domain.Volunteer{}, id).Error
}

func NewVolunteerRepository(db *gorm.DB) VolunteerRepository {
	return &volunteerRepository{db: db}
}

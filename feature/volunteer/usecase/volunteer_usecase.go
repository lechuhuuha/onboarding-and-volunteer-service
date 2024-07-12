package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/storage"
)

type VolunteerUsecase interface {
	CreateVolunteer(input dto.VolunteerCreateDTO) (*domain.Volunteer, error)
	GetAllVolunteers() ([]*domain.Volunteer, error)
	GetVolunteerByID(id uint) (*domain.Volunteer, error)
	UpdateVolunteer(id uint, input dto.VolunteerUpdateDTO) (*domain.Volunteer, error)
	DeleteVolunteer(id uint) error
}

type volunteerUsecase struct {
	repo storage.VolunteerRepository
}

// CreateVolunteer implements VolunteerUsecase.
func (v *volunteerUsecase) CreateVolunteer(input dto.VolunteerCreateDTO) (*domain.Volunteer, error) {
	volunteer := &domain.Volunteer{
		UserId: input.UserId,
		DepartmentId: input.DepartmentId,
		Status: input.Status,
	}
	err := v.repo.Create(volunteer)
	return volunteer, err
}

// GetAllVolunteers implements VolunteerUsecase.
func (v *volunteerUsecase) GetAllVolunteers() ([]*domain.Volunteer, error) {
	return v.repo.GetAll()
}

// GetVolunteerByID implements VolunteerUsecase.
func (v *volunteerUsecase) GetVolunteerByID(id uint) (*domain.Volunteer, error) {
	return v.repo.GetByID(id)
}

// UpdateVolunteer implements VolunteerUsecase.
func (v *volunteerUsecase) UpdateVolunteer(id uint, input dto.VolunteerUpdateDTO) (*domain.Volunteer, error) {
	volunteer, err := v.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	volunteer.DepartmentId = input.DepartmentId
	volunteer.Status = input.Status
	err = v.repo.Update(volunteer)
	return volunteer, err
}

// DeleteVolunteer implements VolunteerUsecase.
func (v *volunteerUsecase) DeleteVolunteer(id uint) error {
	return v.repo.Delete(id)
}


func NewVolunteerUsecase(repo storage.VolunteerRepository) VolunteerUsecase {
	return &volunteerUsecase{repo: repo}
}

package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
)

type VolunteerRequestUsecase struct {
	VolRequestRepo *storage.VolunteerRequestRepository
}

func NewVolunteerRequestUsecase(volRequestRepo *storage.VolunteerRequestRepository) *VolunteerRequestUsecase {
	return &VolunteerRequestUsecase{VolRequestRepo: volRequestRepo}
}

func (u *VolunteerRequestUsecase) CreateRequest(request dto.VoluteerRequestCreatingDTO) error {
	req := &domain.VolunteerRequest{
		UserID: request.UserID,
		Type:   request.Type,
		Status: request.Status,
	}
	return u.VolRequestRepo.CreateVolunteerRequest(req)
}

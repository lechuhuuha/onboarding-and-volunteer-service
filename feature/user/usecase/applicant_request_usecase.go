package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
)

type ApplicantRequestUsecase struct {
	RequestRepo *storage.ApplicantRequestRepository
}

func NewApplicantRequestUsecase(requestRepo *storage.ApplicantRequestRepository) *ApplicantRequestUsecase {
	return &ApplicantRequestUsecase{RequestRepo: requestRepo}
}

func (u *ApplicantRequestUsecase) CreateRequest(request dto.ApplicantRequestCreatingDTO) error {
	req := &domain.ApplicantRequestDomain{
		UserID: request.UserID,
		Type:   request.Type,
		Status: request.Status,
	}
	return u.RequestRepo.CreateApplicantRequest(req)
}

package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
)

type ApplicantRequestUsecase struct {
	repo *storage.ApplicantRequestRepository
}

// NewCountryUsecase creates a new instance of CountryUsecase.
func NewApplicantRequestUsecase(repo *storage.ApplicantRequestRepository) *ApplicantRequestUsecase {
	return &ApplicantRequestUsecase{repo: repo}
}

// Function submit application form, tạo ra 1 record trong bảng request
func (u *ApplicantRequestUsecase) SubmitApplicationForm(appFormDTO dto.ApplicationFormDTO) error {
	//Function parse thời gian về đúng format mình cần dùng

	request := domain.Request{
		UserID: appFormDTO.UserID,
		Type:   "application form",
		Status: 0, // Dat 0 lam default
	}

	return u.repo.CreateRequest(&request)
}

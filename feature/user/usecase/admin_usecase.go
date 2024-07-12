package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
)

type AdminApproveUseCase struct {
	adminRepository *storage.AdminRepository
}

func NewAuthUseCase(userRepo *storage.AdminRepository) *AdminApproveUseCase {
	return &AdminApproveUseCase{userRepo}
}

func (u *AdminApproveUseCase) ApproveRequest(req dto.ApproveRequest) (*dto.ApproveResponse, error) {
	msg, err := u.adminRepository.ApproveRequest(req.ID)
	// if user does not exist in the database, insert i
	if err != nil {
		return nil, err
	}

	return &dto.ApproveResponse{Message: msg}, nil
}

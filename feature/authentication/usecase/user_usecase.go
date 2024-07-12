package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/storage"
)

func Login(req dto.LoginUserRequest) (*dto.LoginUserResponse, string) {
	user, msg := storage.GetUserByEmail(req.Email, req.Password)
	if user != nil {
		return &dto.LoginUserResponse{
			ID:                 user.ID,
			RoleID:             user.RoleID,
			DepartmentID:       user.DepartmentID,
			Email:              user.Email,
			Name:               user.Name,
			Surname:            user.Surname,
			Gender:             user.Gender,
			DOB:                user.DOB,
			Mobile:             user.Mobile,
			CountryID:          user.CountryID,
			ResidentCountryID:  user.ResidentCountryID,
			Avatar:             user.Avatar,
			VerificationStatus: user.VerificationStatus,
			Status:             user.Status,
		}, ""
	}
	return nil, msg
}

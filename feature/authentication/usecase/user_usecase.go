package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/storage"
)

type UserUsecase struct {
	repo *storage.AuthenticationRepository
}

func NewUserUsecase(repo *storage.AuthenticationRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}
func (u *UserUsecase) Login(req dto.LoginUserRequest) (*dto.LoginUserResponse, string) {
	user, msg := u.repo.GetUserByEmail(req.Email, req.Password)
	if user != nil {
		return &dto.LoginUserResponse{
			ID:                 user.ID,
			RoleID:             user.RoleID,
			DepartmentID:       user.DepartmentID,
			Email:              user.Email,
			Name:               user.Name,
			Surname:            user.Surname,
			Gender:             user.Gender,
			DOB:                user.Dob,
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

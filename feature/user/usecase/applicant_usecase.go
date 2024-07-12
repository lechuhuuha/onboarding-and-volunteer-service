package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/domain"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/storage"
)

type ApplicantUsecase struct {
	repo *storage.ApplicantRepository
}

// NewCountryUsecase creates a new instance of CountryUsecase.
func NewCountryUsecase(repo *storage.ApplicantRepository) *ApplicantUsecase {
	return &ApplicantUsecase{repo: repo}
}

// Tao 1 user lần đầu tiên
func (u *ApplicantUsecase) SignupUser(userDTO dto.UserSignupDTO) error {
	user := domain.User{
		Email:   userDTO.Email,
		Name:    userDTO.Name,
		Surname: userDTO.Surname,
	}
	return u.repo.CreateUser(&user)
}

// Update user khi điền xong application form
func (u *ApplicantUsecase) UpdateUser(userDTO dto.UserUpdateDTO) error {
	user := domain.User{
		ID:      userDTO.ID,
		Email:   userDTO.Email,
		Name:    userDTO.Name,
		Surname: userDTO.Surname,
	}
	return u.repo.UpdateUser(&user)
}

// Delete
func (u *ApplicantUsecase) DeleteUser(userID uint) error {
	return u.repo.DeleteUser(userID)
}

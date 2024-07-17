package usecase

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/storage"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type UserUsecase struct {
	repo      *storage.AuthenticationRepository
	secretKey string
}

func NewUserUsecase(repo *storage.AuthenticationRepository, secretKey string) *UserUsecase {
	return &UserUsecase{repo: repo,
		secretKey: secretKey}
}
func (u *UserUsecase) Login(req dto.LoginUserRequest) (*dto.LoginUserTokenResponse, string) {
	user, msg := u.repo.GetUserByEmail(req.Email, req.Password)
	if user != nil {
		claims := jwt.MapClaims{
			"userId": user.ID,
			"roleId": user.RoleID,
			"exp":    time.Now().Add(time.Hour * 72).Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString([]byte(u.secretKey))
		if err != nil {
			return nil, "Could not generate token"
		}
		return &dto.LoginUserTokenResponse{
			Token: tokenString,
		}, ""
	}
	return nil, msg
}

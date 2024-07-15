package transport

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/authentication/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthenticationHandler struct {
	usecase *usecase.UserUsecase
}

func NewAuthenticationHandler(usecase *usecase.UserUsecase) *AuthenticationHandler {
	return &AuthenticationHandler{usecase: usecase}
}
func (h *AuthenticationHandler) Login(c *gin.Context) {
	var req dto.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, msg := h.usecase.Login(req)
	if msg != "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": msg})
		return
	}

	c.JSON(http.StatusOK, resp)
}

package transport

import (
	"net/http"
	"strconv"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/usecase"
	"github.com/gin-gonic/gin"
)

type ApplicantHandler struct {
	usecase *usecase.ApplicantUsecase
}

// NewCountryHandler creates a new instance of CountryHandler.
func NewApplicantHandler(usecase *usecase.ApplicantUsecase) *ApplicantHandler {
	return &ApplicantHandler{usecase: usecase}
}

// Register một người dùng mới qua signup form
func (h *ApplicantHandler) SignupUser(c *gin.Context) {
	var userDTO dto.UserSignupDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.SignupUser(userDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User signed up successfully"})
}

// Update thông tin người dùng ở bảng table khi sử dụng application form
func (h *ApplicantHandler) UpdateUser(c *gin.Context) {
	var userDTO dto.UserUpdateDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.usecase.UpdateUser(userDTO); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

// xóa dữ liệu người dùng ở table user
func (h *ApplicantHandler) DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := h.usecase.DeleteUser(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

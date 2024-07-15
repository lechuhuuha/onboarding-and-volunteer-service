package http

import (
	"net/http"
	"strconv"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user_identity/usecase"
	"github.com/gin-gonic/gin"
)

type UserIdentityHandler struct {
	UserIdentityUsecase *usecase.UserIdentityUsecase
}

func NewUserIdentityHandler(userIdentityUsecase *usecase.UserIdentityUsecase) *UserIdentityHandler {
	return &UserIdentityHandler{UserIdentityUsecase: userIdentityUsecase}
}

func (h *UserIdentityHandler) CreateUserIdentity(c *gin.Context) {
	var request dto.CreateUserIdentityRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserIdentityUsecase.CreateUserIdentity(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User identity created successfully"})
}

func (h *UserIdentityHandler) UpdateUserIdentity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid identity ID"})
		return
	}

	var request dto.UpdateUserIdentityRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UserIdentityUsecase.UpdateUserIdentity(id, request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User identity updated successfully"})
}

func (h *UserIdentityHandler) FindUserIdentity(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid identity ID"})
		return
	}

	identity, err := h.UserIdentityUsecase.FindUserIdentityByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, identity)
}

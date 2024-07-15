package transport

import (
	"net/http"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/usecase"
	"github.com/gin-gonic/gin"
)

type VolunteerRequestHandler struct {
	VolRequestUsecase *usecase.VolunteerRequestUsecase
}

func NewVolunteerRequestHandler(volRequestUsecase *usecase.VolunteerRequestUsecase) *VolunteerRequestHandler {
	return &VolunteerRequestHandler{VolRequestUsecase: volRequestUsecase}
}

func (h *VolunteerRequestHandler) CreateRequest(c *gin.Context) {
	var request dto.VoluteerRequestCreatingDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.VolRequestUsecase.CreateRequest(request); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Request created successfully"})
}
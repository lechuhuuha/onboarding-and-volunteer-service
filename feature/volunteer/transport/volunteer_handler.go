package transport

import (
	"net/http"
	"strconv"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/usecase"
	"github.com/gin-gonic/gin"
)

type VolunteerHandler struct {
	VolUsecaseH usecase.VolunteerUsecaseInterface
}

func NewVolunteerHandler(volUsecase usecase.VolunteerUsecaseInterface) *VolunteerHandler {
	return &VolunteerHandler{VolUsecaseH: volUsecase}
}

func (h *VolunteerHandler) CreateVolunteer(c *gin.Context) {
	var input dto.VolunteerCreateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.VolUsecaseH.CreateVolunteer(input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func (h *VolunteerHandler) UpdateVolunteer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid volunteer ID"})
		return
	}

	var input dto.VolunteerUpdateDTO
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.VolUsecaseH.UpdateVolunteer(id, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Volunteer updated successfully"})
}

func (h *VolunteerHandler) DeleteVolunteer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid volunteer ID"})
		return
	}

	if err = h.VolUsecaseH.DeleteVolunteer(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Volunteer deleted successfully"})
}

func (h *VolunteerHandler) FindVolunteerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	volunteer, err := h.VolUsecaseH.FindVolunteerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Volunteer not found"})
		return
	}

	c.JSON(http.StatusOK, volunteer)
}

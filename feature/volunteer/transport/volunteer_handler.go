package transport

import (
	"net/http"
	"strconv"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/volunteer/usecase"
	"github.com/gin-gonic/gin"
)

type VolunteerHandler interface {
	CreateVolunteer(c *gin.Context)
	GetAllVolunteers(c *gin.Context)
	GetVolunteerByID(c *gin.Context)
	UpdateVolunteer(c *gin.Context)
	DeleteVolunteer(c *gin.Context)
}

type volunteerHandler struct {
	usecase usecase.VolunteerUsecase
}

// CreateVolunteer implements VolunteerHandler.
func (v *volunteerHandler) CreateVolunteer(c *gin.Context) {
	var input dto.VolunteerCreateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	volunteer, err := v.usecase.CreateVolunteer(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, volunteer)
}

// GetAllVolunteers implements VolunteerHandler.
func (v *volunteerHandler) GetAllVolunteers(c *gin.Context) {
	volunteers, err := v.usecase.GetAllVolunteers()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No volunteers found"})
	}

	c.JSON(http.StatusOK, volunteers)
}

// GetVolunteerByID implements VolunteerHandler.
func (v *volunteerHandler) GetVolunteerByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	volunteer, err := v.usecase.GetVolunteerByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Volunteer not found"})
		return
	}

	c.JSON(http.StatusOK, volunteer)
}

// UpdateVolunteer implements VolunteerHandler.
func (v *volunteerHandler) UpdateVolunteer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	var input dto.VolunteerUpdateDTO
	if err = c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	volunteer, err := v.usecase.UpdateVolunteer(uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, volunteer)
}

// DeleteVolunteer implements VolunteerHandler.
func (v *volunteerHandler) DeleteVolunteer(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	err = v.usecase.DeleteVolunteer(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func NewVolunteerHandler(usecase usecase.VolunteerUsecase) VolunteerHandler {
	return &volunteerHandler{usecase: usecase}
}

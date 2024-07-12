package transport

import (
	"net/http"
	"strconv"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/country/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/country/usecase"
	"github.com/gin-gonic/gin"
)

// CountryHandler handles the HTTP requests for countries.
type CountryHandler struct {
	usecase *usecase.CountryUsecase
}

// NewCountryHandler creates a new instance of CountryHandler.
func NewCountryHandler(usecase *usecase.CountryUsecase) *CountryHandler {
	return &CountryHandler{usecase: usecase}
}

// CreateCountry handles the HTTP POST request to create a new country.
func (h *CountryHandler) CreateCountry(c *gin.Context) {
	var input dto.CountryCreateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	country, err := h.usecase.CreateCountry(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, country)
}

// GetCountryByID handles the HTTP GET request to retrieve a country by its ID.
func (h *CountryHandler) GetCountryByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID"})
		return
	}

	country, err := h.usecase.GetCountryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Country not found"})
		return
	}

	c.JSON(http.StatusOK, country)
}

// UpdateCountry handles the HTTP PUT request to update a country.
func (h *CountryHandler) UpdateCountry(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID"})
		return
	}

	var input dto.CountryUpdateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	country, err := h.usecase.UpdateCountry(uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, country)
}

// DeleteCountry handles the HTTP DELETE request to delete a country.
func (h *CountryHandler) DeleteCountry(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid country ID"})
		return
	}

	err = h.usecase.DeleteCountry(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

package transport

import (
	"net/http"
	"strconv"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/department/usecase"
	"github.com/gin-gonic/gin"
)

// DepartmentHandler handles the HTTP requests for departments.
type DepartmentHandler struct {
	usecase *usecase.DepartmentUsecase
}

// NewDepartmentHandler creates a new instance of DepartmentHandler.
func NewDepartmentHandler(usecase *usecase.DepartmentUsecase) *DepartmentHandler {
	return &DepartmentHandler{usecase: usecase}
}

// CreateDepartment handles the HTTP POST request to create a new department.
func (h *DepartmentHandler) CreateDepartment(c *gin.Context) {
	var input dto.DepartmentCreateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	department, err := h.usecase.CreateDepartment(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, department)
}

// GetDepartmentByID handles the HTTP GET request to retrieve a department by its ID.
func (h *DepartmentHandler) GetDepartmentByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}

	department, err := h.usecase.GetDepartmentByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Department not found"})
		return
	}

	c.JSON(http.StatusOK, department)
}

// UpdateDepartment handles the HTTP PUT request to update a department.
func (h *DepartmentHandler) UpdateDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}

	var input dto.DepartmentUpdateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	department, err := h.usecase.UpdateDepartment(uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, department)
}

// DeleteDepartment handles the HTTP DELETE request to delete a department.
func (h *DepartmentHandler) DeleteDepartment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}

	err = h.usecase.DeleteDepartment(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

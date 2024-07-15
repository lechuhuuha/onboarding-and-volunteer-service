package transport

import (
	"net/http"
	"strconv"

	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/role/usecase"
	"github.com/gin-gonic/gin"
)

// RoleHandler handles the HTTP requests for roles.
type RoleHandler struct {
	usecase *usecase.RoleUsecase
}

// NewRoleHandler creates a new instance of RoleHandler.
func NewRoleHandler(usecase *usecase.RoleUsecase) *RoleHandler {
	return &RoleHandler{usecase: usecase}
}

// CreateRole handles the HTTP POST request to create a new role.
func (h *RoleHandler) CreateRole(c *gin.Context) {
	var input dto.RoleCreateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := h.usecase.CreateRole(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}

// GetRoleByID handles the HTTP GET request to retrieve a role by its ID.
func (h *RoleHandler) GetRoleByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	role, err := h.usecase.GetRoleByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Role not found"})
		return
	}

	c.JSON(http.StatusOK, role)
}

// UpdateRole handles the HTTP PUT request to update a role.
func (h *RoleHandler) UpdateRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	var input dto.RoleUpdateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := h.usecase.UpdateRole(uint(id), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

// DeleteRole handles the HTTP DELETE request to delete a role.
func (h *RoleHandler) DeleteRole(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	err = h.usecase.DeleteRole(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

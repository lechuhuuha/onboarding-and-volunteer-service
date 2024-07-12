package transport

import (
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/dto"
	"github.com/cesc1802/onboarding-and-volunteer-service/feature/user/usecase"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AdminHandler struct {
	usecase *usecase.AdminUsecase
}

func NewAuthenticationHandler(usecase *usecase.AdminUsecase) *AdminHandler {
	return &AdminHandler{usecase: usecase}
}
func (h *AdminHandler) GetListRequest(c *gin.Context) {
	resp, msg := h.usecase.GetListRequest()
	if msg != "" {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *AdminHandler) GetRequestById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	resp, msg := h.usecase.GetRequestById(id)
	if msg != "" {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *AdminHandler) ApproveRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	userId := c.MustGet("userId").(int)
	msg := h.usecase.ApproveRequest(id, userId)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (h *AdminHandler) RejectRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	userId := c.MustGet("userId").(int)
	msg := h.usecase.RejectRequest(id, userId)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (h *AdminHandler) AddRejectNotes(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	var req dto.AddRejectNoteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	msg := h.usecase.AddRejectNotes(id, req.Notes)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func (h *AdminHandler) DeleteRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request ID"})
		return
	}
	msg := h.usecase.DeleteRequest(id)
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

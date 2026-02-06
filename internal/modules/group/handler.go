package group

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/middleware"
)

type createGroupRequest struct {
	GroupName string `json:"groupName" binding:"required"`
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// CreateGroup handles POST /group
func (h *Handler) CreateGroup(c *gin.Context) {
	var req createGroupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "groupName is required"})
		return
	}

	userID, ok := middleware.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	group, err := h.service.CreateGroup(req.GroupName, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": group})
}

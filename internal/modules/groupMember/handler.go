package groupMember

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

// GetAllGroupMembers handles GET /group/:groupId/members
func (h *Handler) GetAllGroupMembers(c *gin.Context) {
	groupIDParam := c.Param("groupId")
	gid, err := strconv.ParseUint(groupIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid groupId"})
		return
	}

	members, err := h.service.GetAllGroupMembers(gid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": members})
}

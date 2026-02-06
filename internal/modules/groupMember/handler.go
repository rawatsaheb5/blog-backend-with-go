package groupMember

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/middleware"
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

// GetUserGroups handles GET /group to list group IDs for the authenticated user
func (h *Handler) GetUserGroups(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	groupIDs, err := h.service.GetUserGroupIDs(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": groupIDs})
}

// LeaveGroup handles POST /group/:groupId/leave to set status = LEFT for the current user
func (h *Handler) LeaveGroup(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	groupIDParam := c.Param("groupId")
	gid, err := strconv.ParseUint(groupIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid groupId"})
		return
	}
	success, err := h.service.LeaveGroup(gid, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusNotFound, gin.H{"error": "membership not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "left group"})
}

// InviteMember handles POST /group/:groupId/invite to generate an invite link and (stub) send email
func (h *Handler) InviteMember(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	groupIDParam := c.Param("groupId")
	gid, err := strconv.ParseUint(groupIDParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid groupId"})
		return
	}
	var body struct{ Email string `json:"email" binding:"required,email"` }
	if err := c.ShouldBindJSON(&body); err != nil || body.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "valid email is required"})
		return
	}
	link, err := h.service.GenerateInviteLink(gid, userID, body.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// TODO: integrate actual email sending service here. For now, return the link.
	c.JSON(http.StatusOK, gin.H{"message": "invite generated", "inviteLink": link})
}

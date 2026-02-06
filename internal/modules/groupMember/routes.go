package groupMember

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	h := NewHandler(svc)

	r.GET("/group", h.GetUserGroups)
	r.GET("/group/:groupId/members", h.GetAllGroupMembers)
}

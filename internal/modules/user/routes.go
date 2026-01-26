package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
    repo := NewRepository(db)
    svc := NewService(repo)

    r.POST("/users", Register(svc))
}

package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB, jwtSecret string) {
    repo := NewRepository(db)
    svc := NewService(repo, jwtSecret)

    r.POST("/register", Register(svc))
	r.POST("/login", Login(svc))
}

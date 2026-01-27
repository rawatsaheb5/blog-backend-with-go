package post

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePostRoutes(r *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	r.POST("/", CreatePost(svc))
	r.GET("/:id", GetPostByID(svc))
	r.GET("/", GetPosts(svc))
	r.PUT("/:id", UpdatePost(svc))
	r.DELETE("/:id", DeletePost(svc))
}
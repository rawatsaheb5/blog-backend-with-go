package post

import (
	"github.com/gin-gonic/gin"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/middleware"
)

func CreatePost(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Title   string `json:"title" binding:"required"`
			Content string `json:"content" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "title and content are required"})
			return
		}

		// Extract user ID from context (set by auth middleware)
		userID, exists := middleware.GetUserID(c)
		if !exists {
			c.JSON(401, gin.H{"error": "user not authenticated"})
			return
		}

		if err := svc.CreatePost(req.Title, req.Content, userID); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, gin.H{"message": "post created"})
	}
}

func GetPostByID(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		post, err := svc.GetPostByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "post not found"})
			return
		}
		c.JSON(200, post)
	}
}

func GetPosts(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		posts, err := svc.GetPosts()
		if err != nil {
			c.JSON(404, gin.H{"error": "posts not found"})
			return
		}
		c.JSON(200, posts)
	}
}

func UpdatePost(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		
		// Extract user ID from context
		userID, exists := middleware.GetUserID(c)
		if !exists {
			c.JSON(401, gin.H{"error": "user not authenticated"})
			return
		}

		var req struct {
			Title   string `json:"title" binding:"required"`
			Content string `json:"content" binding:"required"`
		}
		
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "title and content are required"})
			return
		}

		// Get the post first to check ownership
		post, err := svc.GetPostByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "post not found"})
			return
		}

		// Verify user owns the post
		if post.UserID != userID {
			c.JSON(403, gin.H{"error": "you can only update your own posts"})
			return
		}

		// Update the post
		post.Title = req.Title
		post.Content = req.Content
		if err := svc.UpdatePost(post); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "post updated"})
	}
}

func DeletePost(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		
		// Extract user ID from context
		userID, exists := middleware.GetUserID(c)
		if !exists {
			c.JSON(401, gin.H{"error": "user not authenticated"})
			return
		}

		// Get the post first to check ownership
		post, err := svc.GetPostByID(id)
		if err != nil {
			c.JSON(404, gin.H{"error": "post not found"})
			return
		}

		// Verify user owns the post
		if post.UserID != userID {
			c.JSON(403, gin.H{"error": "you can only delete your own posts"})
			return
		}

		if err := svc.DeletePost(id); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"message": "post deleted"})
	}
}
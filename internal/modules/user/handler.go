package user

import (
	"log"
	"github.com/gin-gonic/gin"
)

func Register(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "email and password are required"})
			return
		}

		if err := svc.Register(req.Email, req.Password); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		c.JSON(201, gin.H{"message": "user created"})
	}
}

func Login(svc *Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": "email and password are required"})
			return
		}

		user, token, err := svc.Login(req.Email, req.Password)
		if err != nil {
			c.JSON(401, gin.H{"error": err.Error()})
			return
		}
		// Set token in cookie
		c.SetCookie(
			"token",           // Cookie name
			token,             // Cookie value
			3600*24,          // Max age in seconds (24 hours)
			"/",               // Path
			"",                // Domain (empty = current domain)
			false,             // Secure (set to true in production with HTTPS)
			true,              // HttpOnly (prevents JavaScript access)
		)

		c.JSON(200, gin.H{
			"message": "login successful",
			"data": gin.H{
				"user": gin.H{
					"id":       user.ID,
					"email":    user.Email,
					"username": user.Username,
				},
			},
		})
	}
}
package expense

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.RouterGroup, db *gorm.DB) {
	h := NewHandler(db)
	r.GET("/group/:groupId/expense", h.ListGroupExpenses)
	r.POST("/group/:groupId/expenses", h.CreateExpense)
}

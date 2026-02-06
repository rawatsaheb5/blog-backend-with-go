package expense

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/middleware"
	expensesplit "github.com/rawatsaheb5/blog-backend-with-go/internal/modules/expenseSplit"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/modules/groupMember"
	"gorm.io/gorm"
)

type Handler struct { svc Service }

func NewHandler(db *gorm.DB) *Handler {
	expRepo := NewRepository(db)
	splitRepo := expensesplit.NewRepository(db)
	gmRepo := groupMember.NewRepository(db)
	svc := NewService(expRepo, splitRepo, gmRepo)
	return &Handler{svc: svc}
}

type createExpenseRequest struct {
	Title       string  `json:"title" binding:"required"`
	Amount      float64 `json:"amount" binding:"required,gt=0"`
	PaidBy      uint64  `json:"paidBy" binding:"required"`
	SplitType   string  `json:"splitType" binding:"required"`
	Note        string  `json:"note"`
	ExpenseDate string  `json:"expenseDate"`
	Participants []struct {
		UserID uint64  `json:"userId" binding:"required"`
		Share  float64 `json:"share,omitempty"`
	} `json:"participants" binding:"required,min=1"`
}

func (h *Handler) ListGroupExpenses(c *gin.Context) {
	groupIDStr := c.Param("groupId")
	var gid uint64
	_, err := fmt.Sscanf(groupIDStr, "%d", &gid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid groupId"})
		return
	}
	exps, err := h.svc.ListGroupExpenses(gid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": exps})
}

func (h *Handler) CreateExpense(c *gin.Context) {
	userID, ok := middleware.GetUserID(c)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	groupIDStr := c.Param("groupId")
	var gid uint64
	_, err := fmt.Sscanf(groupIDStr, "%d", &gid)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid groupId"})
		return
	}
	var req createExpenseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var expDate time.Time
	if req.ExpenseDate != "" {
		if t, err := time.Parse(time.RFC3339, req.ExpenseDate); err == nil {
			expDate = t
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid expenseDate; use RFC3339 format"})
			return
		}
	} else {
		expDate = time.Now()
	}

	participants := make([]ParticipantInput, 0, len(req.Participants))
	for _, p := range req.Participants {
		participants = append(participants, ParticipantInput{UserID: p.UserID, Share: p.Share})
	}
	input := CreateExpenseInput{
		GroupID: gid,
		Title: req.Title,
		Amount: req.Amount,
		PaidBy: req.PaidBy,
		SplitType: req.SplitType,
		Note: req.Note,
		ExpenseDate: expDate,
		Participants: participants,
		CreatedBy: userID,
	}
	exp, err := h.svc.CreateExpense(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"data": exp})
}

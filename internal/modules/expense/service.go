package expense

import (
	"errors"
	"time"

	expensesplit "github.com/rawatsaheb5/blog-backend-with-go/internal/modules/expenseSplit"
	"github.com/rawatsaheb5/blog-backend-with-go/internal/modules/groupMember"
)

type Service interface {
	CreateExpense(input CreateExpenseInput) (*Expense, error)
	ListGroupExpenses(groupID uint64) ([]Expense, error)
	GetExpenseByID(expenseID uint64) (*Expense, error)
}

type service struct {
	expRepo Repository
	splitRepo expensesplit.Repository
	gmRepo   groupMember.Repository
}

func NewService(expRepo Repository, splitRepo expensesplit.Repository, gmRepo groupMember.Repository) Service {
	return &service{expRepo: expRepo, splitRepo: splitRepo, gmRepo: gmRepo}
}

func (s *service) ListGroupExpenses(groupID uint64) ([]Expense, error) {
	return s.expRepo.ListByGroupID(groupID)
}

func (s *service) GetExpenseByID(expenseID uint64) (*Expense, error) {
	return s.expRepo.GetByID(expenseID)
}

type ParticipantInput struct {
	UserID uint64  `json:"userId" binding:"required"`
	Share  float64 `json:"share,omitempty"`    // used for unequal or percentage
}

type CreateExpenseInput struct {
	GroupID    uint64             `json:"-"`
	Title      string             `json:"title" binding:"required"`
	Amount     float64            `json:"amount" binding:"required,gt=0"`
	PaidBy     uint64             `json:"paidBy" binding:"required"`
	SplitType  string             `json:"splitType" binding:"required"`
	Note       string             `json:"note"`
	ExpenseDate time.Time         `json:"expenseDate"`
	Participants []ParticipantInput `json:"participants" binding:"required,min=1"`
	CreatedBy  uint64             `json:"-"`
}

func (s *service) CreateExpense(input CreateExpenseInput) (*Expense, error) {
	// Validate membership for paidBy and each participant
	members, err := s.gmRepo.ListByGroupID(input.GroupID)
	if err != nil { return nil, err }
	memberSet := map[uint64]bool{}
	for _, m := range members { memberSet[m.UserID] = m.Status == "active" }
	if !memberSet[input.PaidBy] { return nil, errors.New("paidBy is not an active member of the group") }
	for _, p := range input.Participants {
		if !memberSet[p.UserID] { return nil, errors.New("one or more participants are not active group members") }
	}

	// Create Expense
	exp := &Expense{
		GroupID: input.GroupID,
		Title: input.Title,
		TotalAmount: input.Amount,
		PaidBy: input.PaidBy,
		ExpenseDate: input.ExpenseDate,
		SplitType: input.SplitType,
		Note: input.Note,
		CreatedBy: input.CreatedBy,
		Status: "ACTIVE",
	}
	if err := s.expRepo.Create(exp); err != nil { return nil, err }

	// Build splits excluding the payer
	splits := make([]expensesplit.ExpenseSplit, 0)
	n := float64(len(input.Participants))
	if n == 0 { return exp, nil }

	switch input.SplitType {
	case "EQUAL", "equal", "Equal":
		share := int64(input.Amount / n)
		for _, p := range input.Participants {
			if p.UserID == input.PaidBy { continue }
			splits = append(splits, expensesplit.ExpenseSplit{ExpenseID: exp.ID, UserID: p.UserID, ShareAmount: share})
		}
	case "UNEQUAL", "PERCENTAGE", "unequal", "percentage":
		for _, p := range input.Participants {
			if p.UserID == input.PaidBy { continue }
			amount := int64(p.Share)
			splits = append(splits, expensesplit.ExpenseSplit{ExpenseID: exp.ID, UserID: p.UserID, ShareAmount: amount})
		}
	default:
		return nil, errors.New("unsupported splitType")
	}

	if err := s.splitRepo.BulkCreate(splits); err != nil { return nil, err }
	return exp, nil
}

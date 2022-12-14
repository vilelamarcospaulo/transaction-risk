package context

import "github.com/vilelamarcospaulo/risk/internal/transaction"

type userContext struct {
	TotalSpend int
	UserCards  map[int]bool
}

func newUserContext() *userContext {
	return &userContext{
		TotalSpend: 0,
		UserCards:  make(map[int]bool),
	}
}

type EvalContext struct {
	userContext map[int]*userContext
}

func NewEvalContext() *EvalContext {
	return &EvalContext{
		userContext: make(map[int]*userContext),
	}
}

func (c *EvalContext) UserProcessSpend(transaction transaction.Transaction) {
	userContext := c.GetUserContext(transaction.User_id)

	userContext.TotalSpend += transaction.Amount_us_cents
	userContext.UserCards[transaction.Card_id] = true
}

func (c *EvalContext) GetUserContext(userId int) *userContext {
	if c.userContext[userId] == nil {
		c.userContext[userId] = newUserContext()
	}

	return c.userContext[userId]
}

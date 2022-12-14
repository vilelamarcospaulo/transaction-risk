package context_test

import (
	"testing"

	context "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_context"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func TestTransactionContext_ShouldIncreaseUserSpendAfterEachConsult(t *testing.T) {
	// ARRANGE
	context := context.NewEvalContext()

	// ACT
	context.UserProcessSpend(transaction.Transaction{
		User_id:         1,
		Card_id:         1,
		Amount_us_cents: 100,
	})

	context.UserProcessSpend(transaction.Transaction{
		User_id:         1,
		Card_id:         2,
		Amount_us_cents: 100,
	})

	context.UserProcessSpend(transaction.Transaction{
		User_id:         2,
		Card_id:         1,
		Amount_us_cents: 100,
	})

	// ASSERT
	user1 := context.GetUserContext(1)
	user2 := context.GetUserContext(2)

	if user1.TotalSpend != 200 {
		t.Errorf("TotalSpend %d but expected 200", user1.TotalSpend)
	}

	userCards := len(user1.UserCards)
	if userCards != 2 {
		t.Errorf("Context has %d cards but expected 2", userCards)
	}

	if user1.TotalSpend != 200 {
		t.Errorf("TotalSpend %d but expected 100", user2.TotalSpend)
	}
}

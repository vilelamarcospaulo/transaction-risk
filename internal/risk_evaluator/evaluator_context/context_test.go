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
		UserId:        1,
		CardId:        1,
		AmountUsCents: 100,
	})

	context.UserProcessSpend(transaction.Transaction{
		UserId:        1,
		CardId:        2,
		AmountUsCents: 100,
	})

	context.UserProcessSpend(transaction.Transaction{
		UserId:        2,
		CardId:        1,
		AmountUsCents: 100,
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

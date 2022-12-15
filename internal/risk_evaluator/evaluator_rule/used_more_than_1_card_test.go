package rule_test

import (
	"testing"

	context "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_context"
	rule "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_rule"
	level "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/risk_level"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func TestShouldReturnErrWhenCalledWithoutContext_1card(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		AmountUsCents: 4_000_00,
	}
	evaluatorNode := rule.MoreThan1Cards()

	// ACT
	_, err := evaluatorNode.EvaluateTransactionRisk(transaction, nil)

	// ASSERT
	if err == nil {
		t.Errorf("should receive error %s", err)
	}
}

func TestTransactionRiskLevelShouldBeLowWhenOnlyOneCardUsed(t *testing.T) {
	// ARRANGE
	t1 := transaction.Transaction{
		UserId:        1,
		CardId:        1,
		AmountUsCents: 4_000_00,
	}
	t2 := transaction.Transaction{
		UserId:        1,
		CardId:        1,
		AmountUsCents: 4_000_00,
	}

	context := context.NewEvalContext()
	context.UserProcessSpend(t1)
	context.UserProcessSpend(t2)

	// ACT
	evaluatorNode := rule.MoreThan1Cards()
	risk, err := evaluatorNode.EvaluateTransactionRisk(t2, context)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != level.Low {
		t.Errorf("received %d but expected Low", risk)
	}
}

func TestTransactionRiskLevelShouldBeMediumWhenTwoCardsUsed(t *testing.T) {
	// ARRANGE
	t1 := transaction.Transaction{
		UserId:        1,
		CardId:        1,
		AmountUsCents: 4_000_00,
	}
	t2 := transaction.Transaction{
		UserId:        1,
		CardId:        2,
		AmountUsCents: 4_000_00,
	}
	t3 := transaction.Transaction{
		UserId:        1,
		CardId:        2,
		AmountUsCents: 4_000_00,
	}

	context := context.NewEvalContext()
	context.UserProcessSpend(t1)
	context.UserProcessSpend(t2)
	context.UserProcessSpend(t3)

	// ACT
	evaluatorNode := rule.MoreThan1Cards()
	risk, err := evaluatorNode.EvaluateTransactionRisk(t3, context)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	// t.Errorf("fooo, %v \n %d\n", context.GetUserContext(1), len(context.GetUserContext(1).UserCards))
	if risk != level.Medium {
		t.Errorf("received %d but expected Medium", risk)
	}
}

package rule_test

import (
	"testing"

	context "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_context"
	rule "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_rule"
	level "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/risk_level"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func TestTransactionRiskLevelShouldBeHighWhenThreeCardsUsed(t *testing.T) {
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
		CardId:        3,
		AmountUsCents: 4_000_00,
	}

	context := context.NewEvalContext()
	context.UserProcessSpend(t1)
	context.UserProcessSpend(t2)
	context.UserProcessSpend(t3)

	// ACT
	evaluatorNode := rule.MoreThan2Cards()
	risk, err := evaluatorNode.EvaluateTransactionRisk(t3, context)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != level.High {
		t.Errorf("received %d but expected High", risk)
	}
}

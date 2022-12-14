package rule_test

import (
	"testing"

	evaluator "github.com/vilelamarcospaulo/risk/internal/risk_evaluator"
	rule "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_rule"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func TestTransactionRiskLevelShouldBeLowWhenAmountBellow10K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		Amount_us_cents: 9_000_00,
	}
	evaluatorNode := rule.AmountAbove10K()

	// ACT
	risk, err := evaluatorNode.EvaluateTransactionRisk(transaction, nil)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != evaluator.Low {
		t.Errorf("received %d but expected Low", risk)
	}
}

func TestTransactionRiskLevelShouldBeHighWhenAmountAbove10K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		Amount_us_cents: 10_001_00,
	}
	evaluatorNode := rule.AmountAbove10K()

	// ACT
	risk, err := evaluatorNode.EvaluateTransactionRisk(transaction, nil)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != evaluator.High {
		t.Errorf("received %d but expected High", risk)
	}
}

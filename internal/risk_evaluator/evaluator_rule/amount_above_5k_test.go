package rule_test

import (
	"testing"

	evaluator "github.com/vilelamarcospaulo/risk/internal/risk_evaluator"
	rule "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_rule"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func TestTransactionRiskLevelShouldBeLowWhenAmountBellow5K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		Amount_us_cents: 4_000_00,
	}
	evaluatorNode := rule.AmountAbove5K()

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

func TestTransactionRiskLevelShouldBeMediumWhenAmountAbove5K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		Amount_us_cents: 5_001_00,
	}
	evaluatorNode := rule.AmountAbove5K()

	// ACT
	risk, err := evaluatorNode.EvaluateTransactionRisk(transaction, nil)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != evaluator.Medium {
		t.Errorf("received %d but expected Medium", risk)
	}
}

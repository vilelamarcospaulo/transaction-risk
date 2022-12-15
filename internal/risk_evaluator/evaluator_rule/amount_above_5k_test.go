package rule_test

import (
	"testing"

	chain "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_chain"
	rule "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_rule"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func TestTransactionRiskLevelShouldBeLowWhenAmountBellow5K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		AmountUsCents: 4_000_00,
	}
	evaluatorNode := rule.AmountAbove5K()

	// ACT
	risk, err := evaluatorNode.EvaluateTransactionRisk(transaction, nil)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != chain.Low {
		t.Errorf("received %d but expected Low", risk)
	}
}

func TestTransactionRiskLevelShouldBeMediumWhenAmountAbove5K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		AmountUsCents: 5_001_00,
	}
	evaluatorNode := rule.AmountAbove5K()

	// ACT
	risk, err := evaluatorNode.EvaluateTransactionRisk(transaction, nil)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != chain.Medium {
		t.Errorf("received %d but expected Medium", risk)
	}
}

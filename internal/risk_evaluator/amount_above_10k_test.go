package riskevaluator_test

import (
	"testing"

	riskevaluator "github.com/vilelamarcospaulo/risk/internal/risk_evaluator"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func TestTransactionRiskLevelShouldBeLowWhenAmountBellow10K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		Amount_us_cents: 9_000_00,
	}
	evaluatorNode := riskevaluator.AmountAbove10K()

	// ACT
	risk, err := evaluatorNode.EvaluateTransactionRisk(transaction)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != riskevaluator.Low {
		t.Errorf("received %d but expected Low", risk)
	}
}

func TestTransactionRiskLevelShouldBeHighWhenAmountAbove10K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		Amount_us_cents: 10_001_00,
	}
	evaluatorNode := riskevaluator.AmountAbove10K()

	// ACT
	risk, err := evaluatorNode.EvaluateTransactionRisk(transaction)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != riskevaluator.High {
		t.Errorf("received %d but expected High", risk)
	}
}

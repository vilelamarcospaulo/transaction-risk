package riskevaluator_test

import (
	"testing"

	riskevaluator "github.com/vilelamarcospaulo/risk/internal/risk_evaluator"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func TestTransactionRiskLevelShouldBeLowWhenAmountBellow5K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		Amount_us_cents: 4_000_00,
	}
	evaluatorNode := riskevaluator.AmountAbove5K()

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

func TestTransactionRiskLevelShouldBeMediumWhenAmountAbove5K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		Amount_us_cents: 5_001_00,
	}
	evaluatorNode := riskevaluator.AmountAbove5K()

	// ACT
	risk, err := evaluatorNode.EvaluateTransactionRisk(transaction)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != riskevaluator.Medium {
		t.Errorf("received %d but expected Medium", risk)
	}
}

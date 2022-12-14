package riskevaluator_test

import (
	"testing"

	riskevaluator "github.com/vilelamarcospaulo/risk/internal/risk_evaluator"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func TestShouldProcessAllChainForTransaction_Bellow5K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		Amount_us_cents: 4_000_00,
	}
	head := riskevaluator.AmountAbove10K()
	node1 := riskevaluator.AmountAbove5K()

	head.Append(node1)

	// ACT
	risk, err := head.EvaluateTransactionRisk(transaction, nil)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != riskevaluator.Low {
		t.Errorf("received %d but expected Low", risk)
	}
}

func TestShouldProcessChainForTransaction_Bellow10K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		Amount_us_cents: 7_000_00,
	}
	head := riskevaluator.AmountAbove10K()
	node1 := riskevaluator.AmountAbove5K()

	head.Append(node1)

	// ACT
	risk, err := head.EvaluateTransactionRisk(transaction, nil)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != riskevaluator.Medium {
		t.Errorf("received %d but expected Medium", risk)
	}
}

func TestShouldProcessChainForTransaction_Above10K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		Amount_us_cents: 10_500_00,
	}
	head := riskevaluator.AmountAbove10K()
	node1 := riskevaluator.AmountAbove5K()

	head.Append(node1)

	// ACT
	risk, err := head.EvaluateTransactionRisk(transaction, nil)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != riskevaluator.High {
		t.Errorf("received %d but expected High", risk)
	}
}

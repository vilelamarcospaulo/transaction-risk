package chain_test

import (
	"testing"

	chain "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_chain"
	rule "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_rule"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func TestShouldProcessAllChainForTransaction_Bellow5K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		Amount_us_cents: 4_000_00,
	}
	head := rule.AmountAbove10K()
	node1 := rule.AmountAbove5K()

	head.Append(node1)

	// ACT
	risk, err := head.EvaluateTransactionRisk(transaction, nil)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != chain.Low {
		t.Errorf("received %d but expected Low", risk)
	}
}

func TestShouldProcessChainForTransaction_Bellow10K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		Amount_us_cents: 7_000_00,
	}
	head := rule.AmountAbove10K()
	node1 := rule.AmountAbove5K()

	head.Append(node1)

	// ACT
	risk, err := head.EvaluateTransactionRisk(transaction, nil)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != chain.Medium {
		t.Errorf("received %d but expected Medium", risk)
	}
}

func TestShouldProcessChainForTransaction_Above10K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		Amount_us_cents: 10_500_00,
	}
	head := rule.AmountAbove10K()
	node1 := rule.AmountAbove5K()

	head.Append(node1)

	// ACT
	risk, err := head.EvaluateTransactionRisk(transaction, nil)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != chain.High {
		t.Errorf("received %d but expected High", risk)
	}
}

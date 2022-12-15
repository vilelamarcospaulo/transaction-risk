package rule_test

import (
	"testing"

	context "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_context"
	rule "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_rule"
	level "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/risk_level"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func TestShouldReturnErrWhenCalledWithoutContext(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		AmountUsCents: 4_000_00,
	}
	evaluatorNode := rule.SpendAbove10K()

	// ACT
	_, err := evaluatorNode.EvaluateTransactionRisk(transaction, nil)

	// ASSERT
	if err == nil {
		t.Errorf("should receive error %s", err)
	}
}

func TestTransactionRiskLevelShouldBeLowWhenAmountBellows10K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		AmountUsCents: 4_000_00,
	}

	context := context.NewEvalContext()
	context.UserProcessSpend(transaction)

	// ACT
	evaluatorNode := rule.SpendAbove10K()
	risk, err := evaluatorNode.EvaluateTransactionRisk(transaction, context)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != level.Low {
		t.Errorf("received %d but expected Low", risk)
	}
}

func TestTransactionRiskLevelShouldBeMediumWhenAmountAbove10K(t *testing.T) {
	// ARRANGE
	transaction := transaction.Transaction{
		AmountUsCents: 15_000_00,
	}

	context := context.NewEvalContext()
	context.UserProcessSpend(transaction)

	// ACT
	evaluatorNode := rule.SpendAbove10K()
	risk, err := evaluatorNode.EvaluateTransactionRisk(transaction, context)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	if risk != level.Medium {
		t.Errorf("received %d but expected Medium", risk)
	}
}

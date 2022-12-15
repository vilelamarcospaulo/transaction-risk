package evaluator_test

import (
	"reflect"
	"testing"

	evaluator "github.com/vilelamarcospaulo/risk/internal/risk_evaluator"
	chain "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_chain"
	rule "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_rule"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func TestShouldProcessAllTransactionInEachRule(t *testing.T) {
	// ARRANGE
	transactions := []transaction.Transaction{
		{
			UserId:        1,
			CardId:        1,
			AmountUsCents: 4_000_00,
		},
		{
			UserId:        1,
			CardId:        1,
			AmountUsCents: 3_000_00,
		},
		{
			UserId:        2,
			CardId:        1,
			AmountUsCents: 10_000_01,
		},
	}

	rules := []*chain.Node{
		rule.AmountAbove10K(),
		rule.AmountAbove5K(),
	}

	// ACT
	evaluator := evaluator.NewEvaluator(rules)
	result, err := evaluator.Eval(transactions)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	expected := []chain.RiskLevel{chain.Low, chain.Low, chain.High}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("received %v expetected %v", result, expected)
	}
}

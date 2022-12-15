package httpz_test

import (
	"reflect"
	"testing"

	evaluator "github.com/vilelamarcospaulo/risk/internal/risk_evaluator"
	chain "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_chain"
	rule "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_rule"
	httpz "github.com/vilelamarcospaulo/risk/ports/http"
)

func TestShouldProcessAllTransactionAndBuildReturn(t *testing.T) {
	// ARRANGE
	req := httpz.RiskRequest{
		Transactions: []httpz.Transaction{
			{
				Id:              1,
				User_id:         1,
				Card_id:         1,
				Amount_us_cents: 4_000_00,
			},
			{
				Id:              2,
				User_id:         1,
				Card_id:         1,
				Amount_us_cents: 3_000_00,
			},
			{
				Id:              3,
				User_id:         2,
				Card_id:         1,
				Amount_us_cents: 10_000_01,
			},
		},
	}

	rules := []*chain.Node{
		rule.AmountAbove10K(),
		rule.AmountAbove5K(),
	}

	evaluator := evaluator.NewEvaluator(rules)

	// ACT
	controller := httpz.NewRiskController(evaluator)
	result, err := controller.Process(req)

	// ASSERT
	if err != nil {
		t.Errorf("received error %s", err)
	}

	expected := httpz.RiskResponse{
		Risk_ratings: []string{
			"Low", "Low", "High",
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("received %v expetected %v", result, expected)
	}
}

package rule

import (
	"errors"

	chain "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_chain"
	context "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_context"
	level "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/risk_level"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func SpendAbove10K() *chain.Node {
	return spendAbove(10_000_00, level.Medium)
}

func SpendAbove20K() *chain.Node {
	return spendAbove(20_000_00, level.High)
}

func spendAbove(valueInUsCents int, riskLevel level.RiskLevel) *chain.Node {
	return chain.NewNode(checkSpendAbove(valueInUsCents), riskLevel)
}

func checkSpendAbove(valueInUsCents int) func(t transaction.Transaction, context *context.EvalContext) (bool, error) {
	return func(t transaction.Transaction, context *context.EvalContext) (bool, error) {
		if context == nil {
			return false, errors.New("can't define spend without context")
		}

		return context.GetUserContext(t.UserId).TotalSpend > valueInUsCents, nil
	}
}

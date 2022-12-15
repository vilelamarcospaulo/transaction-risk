package rule

import (
	chain "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_chain"
	context "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_context"
	level "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/risk_level"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func AmountAbove5K() *chain.Node {
	return amountAbove(5_000_00, level.Medium)
}

func AmountAbove10K() *chain.Node {
	return amountAbove(10_000_00, level.High)
}

func amountAbove(valueInUsCents int, riskLevel level.RiskLevel) *chain.Node {
	return chain.NewNode(checkValueAbove(valueInUsCents), riskLevel)
}

func checkValueAbove(valueInUsCents int) func(transaction transaction.Transaction, globalContext *context.EvalContext) (bool, error) {
	return func(transaction transaction.Transaction, globalContext *context.EvalContext) (bool, error) {
		return transaction.AmountUsCents > valueInUsCents, nil
	}
}

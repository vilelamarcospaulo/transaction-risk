package rule

import (
	chain "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_chain"
	context "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_context"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func AmountAbove5K() *chain.Node {
	return amountAbove(5_000_00, chain.Medium)
}

func AmountAbove10K() *chain.Node {
	return amountAbove(10_000_00, chain.High)
}

func amountAbove(valueInUsCents int, riskLevel chain.RiskLevel) *chain.Node {
	return chain.NewNode(checkValueAbove(valueInUsCents), riskLevel)
}

func checkValueAbove(valueInUsCents int) func(transaction transaction.Transaction, globalContext *context.EvalContext) (bool, error) {
	return func(transaction transaction.Transaction, globalContext *context.EvalContext) (bool, error) {
		return transaction.Amount_us_cents > valueInUsCents, nil
	}
}

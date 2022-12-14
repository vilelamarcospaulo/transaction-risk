package rule

import (
	evaluator "github.com/vilelamarcospaulo/risk/internal/risk_evaluator"
	chain "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_chain"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func AmountAbove5K() *chain.Node {
	return amountAbove(5_000_00, evaluator.Medium)
}

func AmountAbove10K() *chain.Node {
	return amountAbove(10_000_00, evaluator.High)
}

func amountAbove(valueInUsCents int, riskLevel evaluator.RiskLevel) *chain.Node {
	return chain.NewNode(checkValueAbove(valueInUsCents), riskLevel)
}

func checkValueAbove(valueInUsCents int) func(transaction transaction.Transaction, globalContext []transaction.Transaction) (bool, error) {
	return func(transaction transaction.Transaction, globalContext []transaction.Transaction) (bool, error) {
		return transaction.Amount_us_cents > valueInUsCents, nil
	}
}

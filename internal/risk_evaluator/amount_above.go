package riskevaluator

import "github.com/vilelamarcospaulo/risk/internal/transaction"

func AmountAbove5K() *evaluatorNode {
	return amountAbove(5_000_00, Medium)
}

func AmountAbove10K() *evaluatorNode {
	return amountAbove(10_000_00, High)
}

func amountAbove(valueInUsCents int, riskLevel RiskLevel) *evaluatorNode {
	return &evaluatorNode{
		next:      nil,
		predicate: checkValueAbove(valueInUsCents),
		riskLevel: riskLevel,
	}

}
func checkValueAbove(valueInUsCents int) func(transaction transaction.Transaction, globalContext *[]transaction.Transaction) (bool, error) {
	return func(transaction transaction.Transaction, globalContext *[]transaction.Transaction) (bool, error) {
		return transaction.Amount_us_cents > valueInUsCents, nil
	}
}

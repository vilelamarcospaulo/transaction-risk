package riskevaluator

import "github.com/vilelamarcospaulo/risk/internal/transaction"

func AmountAbove5K() *evaluatorNode {
	return &evaluatorNode{
		next:      nil,
		predicate: checkValueAbove5K,
		riskLevel: Medium,
	}
}

func checkValueAbove5K(transaction transaction.Transaction, globalContext *[]transaction.Transaction) (bool, error) {
	const valueInCents = 5000 * 100
	return transaction.Amount_us_cents > valueInCents, nil
}

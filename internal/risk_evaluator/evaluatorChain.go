package riskevaluator

import "github.com/vilelamarcospaulo/risk/internal/transaction"

type RiskLevel int64

const (
	Low RiskLevel = iota
	Medium
	High
)

// Based on chain of responsibility pattern
type evaluatorNode struct {
	next      *evaluatorNode
	predicate func(transaction.Transaction, *[]transaction.Transaction) (bool, error)
	riskLevel RiskLevel
}

func (e evaluatorNode) EvaluateTransactionRisk(transaction transaction.Transaction) (RiskLevel, error) {
	matched, err := e.predicate(transaction, nil)

	if matched {
		return e.riskLevel, nil
	}

	if err != nil {
		return e.riskLevel, err
	}

	if e.next != nil {
		return e.next.EvaluateTransactionRisk(transaction)
	}

	return Low, nil
}

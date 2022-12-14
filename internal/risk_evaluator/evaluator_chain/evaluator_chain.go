package chain

import (
	evaluator "github.com/vilelamarcospaulo/risk/internal/risk_evaluator"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

// Based on chain of responsibility pattern
type Node struct {
	next      *Node
	predicate func(transaction.Transaction, []transaction.Transaction) (bool, error)
	riskLevel evaluator.RiskLevel
}

func NewNode(
	predicate func(transaction.Transaction, []transaction.Transaction) (bool, error),
	riskLevel evaluator.RiskLevel) *Node {
	return &Node{
		next:      nil,
		predicate: predicate,
		riskLevel: riskLevel,
	}
}

func (e Node) EvaluateTransactionRisk(transaction transaction.Transaction, allTransactions []transaction.Transaction) (evaluator.RiskLevel, error) {
	matched, err := e.predicate(transaction, allTransactions)

	if matched {
		return e.riskLevel, nil
	}

	if err != nil {
		return e.riskLevel, err
	}

	if e.next != nil {
		return e.next.EvaluateTransactionRisk(transaction, allTransactions)
	}

	return evaluator.Low, nil
}

func (e *Node) Append(next *Node) {
	current := e
	for ; current.next != nil; current = current.next {
	}

	current.next = next
}

package chain

import (
	context "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_context"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

type RiskLevel int64

const (
	Low RiskLevel = iota
	Medium
	High
)

// Based on chain of responsibility pattern
type Node struct {
	next      *Node
	predicate func(transaction.Transaction, *context.EvalContext) (bool, error)
	riskLevel RiskLevel
}

func NewNode(
	predicate func(transaction.Transaction, *context.EvalContext) (bool, error),
	riskLevel RiskLevel) *Node {
	return &Node{
		next:      nil,
		predicate: predicate,
		riskLevel: riskLevel,
	}
}

func (e Node) EvaluateTransactionRisk(transaction transaction.Transaction, evalContext *context.EvalContext) (RiskLevel, error) {
	if evalContext != nil {
		evalContext.UserProcessSpend(transaction)
	}

	matched, err := e.predicate(transaction, evalContext)

	if matched {
		return e.riskLevel, nil
	}

	if err != nil {
		return e.riskLevel, err
	}

	if e.next != nil {
		return e.next.EvaluateTransactionRisk(transaction, evalContext)
	}

	return Low, nil
}

func (e *Node) Append(next *Node) {
	current := e
	for ; current.next != nil; current = current.next {
	}

	current.next = next
}

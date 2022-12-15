package chain

import (
	context "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_context"
	level "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/risk_level"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

// Based on chain of responsibility pattern
type Predicate = func(transaction.Transaction, *context.EvalContext) (bool, error)
type Node struct {
	next      *Node
	predicate Predicate
	riskLevel level.RiskLevel
}

func NewNode(p Predicate, riskLevel level.RiskLevel) *Node {
	return &Node{
		next:      nil,
		predicate: p,
		riskLevel: riskLevel,
	}
}

func (e Node) EvaluateTransactionRisk(transaction transaction.Transaction, evalContext *context.EvalContext) (level.RiskLevel, error) {
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

	return level.Low, nil
}

func (e *Node) Append(next *Node) {
	current := e
	for ; current.next != nil; current = current.next {
	}

	current.next = next
}

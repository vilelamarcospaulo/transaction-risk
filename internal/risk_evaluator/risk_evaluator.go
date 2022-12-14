package evaluator

import (
	chain "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_chain"
	context "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_context"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

type Evaluator struct {
	head    *chain.Node
	context *context.EvalContext
}

func NewEvaluator(nodes []*chain.Node) *Evaluator {
	// Make sure every single node in attached to to list
	head := nodes[0]
	for i := 1; i < len(nodes); i++ {
		head.Append(nodes[i])
	}

	return &Evaluator{
		head:    head,
		context: context.NewEvalContext(),
	}
}

func (e *Evaluator) Eval(transactions []transaction.Transaction) ([]chain.RiskLevel, error) {
	var result []chain.RiskLevel
	for _, t := range transactions {
		tRisk, err := e.head.EvaluateTransactionRisk(t, e.context)
		if err != nil {
			return nil, err
		}

		result = append(result, tRisk)
	}

	return result, nil
}

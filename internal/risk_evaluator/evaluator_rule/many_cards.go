package rule

import (
	"errors"

	chain "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_chain"
	context "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_context"
	level "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/risk_level"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

func MoreThan1Cards() *chain.Node {
	return manyCards(1, level.Medium)
}

func MoreThan2Cards() *chain.Node {
	return manyCards(2, level.High)
}

func manyCards(cardsLimit int, riskLevel level.RiskLevel) *chain.Node {
	return chain.NewNode(
		func(t transaction.Transaction, context *context.EvalContext) (bool, error) {
			if context == nil {
				return false, errors.New("can't define cards number without context")
			}

			return cardsLimit < len(context.GetUserContext(t.UserId).UserCards), nil
		}, riskLevel)
}

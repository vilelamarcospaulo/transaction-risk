package httpz

import (
	evaluator "github.com/vilelamarcospaulo/risk/internal/risk_evaluator"
	level "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/risk_level"
	"github.com/vilelamarcospaulo/risk/internal/transaction"
)

type RiskController struct {
	evaluator *evaluator.Evaluator
}

func NewRiskController(e *evaluator.Evaluator) *RiskController {
	return &RiskController{
		evaluator: e,
	}
}

func (c *RiskController) Process(req RiskRequest) (RiskResponse, error) {
	transactions := requestToEntity(req)
	result, err := c.evaluator.Eval(transactions)

	if err != nil {
		return RiskResponse{}, err
	}

	return RiskResponse{
		Risk_ratings: evaluationsToStr(result),
	}, nil

}

func requestToEntity(req RiskRequest) []transaction.Transaction {
	var transactions []transaction.Transaction
	for _, t := range req.Transactions {
		transactions = append(transactions, transaction.Transaction{
			UserId:        t.User_id,
			CardId:        t.Card_id,
			AmountUsCents: t.Amount_us_cents,
		})
	}

	return transactions
}

func evaluationsToStr(evaluations []level.RiskLevel) []string {
	var r []string
	for _, evaluation := range evaluations {
		r = append(r, evaluation.String())
	}

	return r
}

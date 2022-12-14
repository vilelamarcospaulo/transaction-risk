package main

import (
	evaluator "github.com/vilelamarcospaulo/risk/internal/risk_evaluator"
	chain "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_chain"
	rule "github.com/vilelamarcospaulo/risk/internal/risk_evaluator/evaluator_rule"
	httpz "github.com/vilelamarcospaulo/risk/ports/http"
)

func main() {
	rules := []*chain.Node{
		rule.MoreThan2Cards(),
		rule.SpendAbove20K(),
		rule.AmountAbove10K(),
		rule.MoreThan1Cards(),
		rule.SpendAbove10K(),
		rule.AmountAbove5K(),
	}
	evaluator := evaluator.NewEvaluator(rules)
	controller := httpz.NewRiskController(evaluator)

	transactionRiskServer := httpz.NewServer(controller)

	transactionRiskServer.Start(3000)
}

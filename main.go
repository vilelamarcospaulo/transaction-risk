package main

import httpz "github.com/vilelamarcospaulo/risk/ports/http"

func main() {
	transactionRiskServer := httpz.NewServer(Controller)

	transactionRiskServer.Start(3000)
}

func Controller(req httpz.RiskRequest) httpz.RiskResponse {
	return httpz.RiskResponse{
		Risk_ratings: []string{"foo", "bar"},
	}
}

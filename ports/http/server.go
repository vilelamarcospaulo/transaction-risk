package httpz

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Transaction struct {
	Id              int
	User_id         int
	Amount_us_cents int
	Card_id         int
}

type RiskRequest struct {
	Transactions []Transaction
}

type RiskResponse struct {
	Risk_ratings []string
}

type RiskController = func(req RiskRequest) RiskResponse

type TransactionRiskHttpServer struct {
	httpHandler *http.ServeMux
	controller  RiskController
}

func NewServer(controller RiskController) *TransactionRiskHttpServer {
	return &TransactionRiskHttpServer{
		httpHandler: http.NewServeMux(),
		controller:  controller,
	}
}

func (s *TransactionRiskHttpServer) Start(port int) {
	// Unique route
	s.httpHandler.HandleFunc("/transaction-risk", processHttpRequest(s.controller))

	// health check end-point
	s.httpHandler.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "healthz")
	})

	http.ListenAndServe(fmt.Sprintf(":%d", port), s.httpHandler)
}

func processHttpRequest(controller RiskController) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var payload RiskRequest
		err := decoder.Decode(&payload)
		if err != nil {
			panic(err)
		}

		result := controller(payload)
		json.NewEncoder(w).Encode(result)
	}
}

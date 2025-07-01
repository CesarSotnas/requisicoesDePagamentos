package handlers

import (
	"encoding/json"
	"github.com/CesarSotnas/requisicoesDePagamentos.git/stats"
	"net/http"

	"github.com/CesarSotnas/requisicoesDePagamentos.git/models"
	"github.com/CesarSotnas/requisicoesDePagamentos.git/services"
)

func ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var req models.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	result, err := services.ProcessPayment(req)
	if err != nil {
		http.Error(w, "error while processing the payment: "+err.Error(), http.StatusServiceUnavailable)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GetStats(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats.Snapshot())
}

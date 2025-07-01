package processors

import (
	"errors"
	"math/rand"
	"time"

	"github.com/CesarSotnas/requisicoesDePagamentos.git/models"
)

func FastPay(req models.PaymentRequest) (*models.PaymentResult, error) {
	// latency simulation between 300~1000 ms
	delay := time.Duration(300+rand.Intn(700)) * time.Millisecond
	time.Sleep(delay)

	// fail chance simulation: 10%
	if rand.Float64() < 0.1 {
		return nil, errors.New("failed FastPay process")
	}

	return &models.PaymentResult{
		ID:     req.ID,
		Amount: req.Amount,
	}, nil
}

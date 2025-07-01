package processors

import (
	"errors"
	"math/rand"
	"time"

	"github.com/CesarSotnas/requisicoesDePagamentos.git/models"
)

func CheapPay(req models.PaymentRequest) (*models.PaymentResult, error) {
	// latency simulation between 1~3 seconds
	delay := time.Duration(1+rand.Intn(3)) * time.Second
	time.Sleep(delay)

	// fail chance simulation: 30%
	if rand.Float64() < 0.3 {
		return nil, errors.New("failed CheapPay process")
	}

	return &models.PaymentResult{
		ID:     req.ID,
		Amount: req.Amount,
	}, nil
}

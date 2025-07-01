package services

import (
	"errors"
	"time"

	"github.com/CesarSotnas/requisicoesDePagamentos.git/models"
	"github.com/CesarSotnas/requisicoesDePagamentos.git/processors"
	"github.com/CesarSotnas/requisicoesDePagamentos.git/stats"
)

const (
	timeoutPerProcessor = 2 * time.Second
)

func ProcessPayment(req models.PaymentRequest) (*models.PaymentResult, error) {
	type resultWithError struct {
		result *models.PaymentResult
		err    error
	}

	processorsInfo := []struct {
		Name     string
		Fee      float64
		Function func(models.PaymentRequest) (*models.PaymentResult, error)
	}{
		{
			Name:     "CheapPay",
			Fee:      0.02, // 2%
			Function: processors.CheapPay,
		},
		{
			Name:     "FastPay",
			Fee:      0.03, // 3%
			Function: processors.FastPay,
		},
	}

	if processorsInfo[1].Fee < processorsInfo[0].Fee {
		processorsInfo[0], processorsInfo[1] = processorsInfo[1], processorsInfo[0]
	}

	for _, p := range processorsInfo {
		resultCh := make(chan resultWithError, 1)

		go func(pInfo struct {
			Name     string
			Fee      float64
			Function func(models.PaymentRequest) (*models.PaymentResult, error)
		}) {
			res, err := pInfo.Function(req)
			resultCh <- resultWithError{result: res, err: err}
		}(p)

		select {
		case r := <-resultCh:
			if r.err == nil {
				r.result.ProcessorUsed = p.Name
				r.result.FeeApplied = req.Amount * p.Fee
				r.result.Success = true
				stats.Register(r.result)
				return r.result, nil
			}
		case <-time.After(timeoutPerProcessor):
			continue
		}
	}

	return nil, errors.New("fail in both processors")
}

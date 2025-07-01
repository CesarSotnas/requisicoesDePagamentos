package stats

import (
	"sync"

	"github.com/CesarSotnas/requisicoesDePagamentos.git/models"
)

type ProcessorStats struct {
	Count int     `json:"count"`
	Total float64 `json:"total"`
	Fee   float64 `json:"fee"`
}

type Stats struct {
	mu         sync.RWMutex
	TotalCount int                        `json:"total_count"`
	TotalFee   float64                    `json:"total_fee"`
	Processors map[string]*ProcessorStats `json:"processors"`
}

var globalStats = &Stats{
	Processors: make(map[string]*ProcessorStats),
}

// Register register the result of a payment
func Register(result *models.PaymentResult) {
	globalStats.mu.Lock()
	defer globalStats.mu.Unlock()

	p := result.ProcessorUsed
	if _, ok := globalStats.Processors[p]; !ok {
		globalStats.Processors[p] = &ProcessorStats{}
	}

	ps := globalStats.Processors[p]
	ps.Count++
	ps.Total += result.Amount
	ps.Fee += result.FeeApplied

	globalStats.TotalCount++
	globalStats.TotalFee += result.FeeApplied
}

// Snapshot returns a safe copy with statistics
func Snapshot() Stats {
	globalStats.mu.RLock()
	defer globalStats.mu.RUnlock()

	clone := Stats{
		TotalCount: globalStats.TotalCount,
		TotalFee:   globalStats.TotalFee,
		Processors: make(map[string]*ProcessorStats),
	}

	for k, v := range globalStats.Processors {
		clone.Processors[k] = &ProcessorStats{
			Count: v.Count,
			Total: v.Total,
			Fee:   v.Fee,
		}
	}

	return clone
}

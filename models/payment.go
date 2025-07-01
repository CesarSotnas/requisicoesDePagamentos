package models

type PaymentRequest struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
}

type PaymentResult struct {
	ID            string  `json:"id"`
	Amount        float64 `json:"amount"`
	ProcessorUsed string  `json:"processor"`
	FeeApplied    float64 `json:"fee"`
	Success       bool    `json:"success"`
}

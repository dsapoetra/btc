package http_model

type Transaction struct {
	Amount    float64 `json:"amount"`
	CreatedAt string  `json:"created_at"`
}

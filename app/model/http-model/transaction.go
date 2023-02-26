package http_model

type Transaction struct {
	Amount   float64 `json:"amount"`
	DateTime string  `json:"datetime"`
}

package http_model

type Transaction struct {
	Amount    float64 `json:"amount"`
	CreatedAt string  `json:"created_at"`
}

type InquiryTransaction struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

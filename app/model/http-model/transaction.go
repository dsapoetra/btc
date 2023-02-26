package http_model

type Transaction struct {
	Amount   float64 `json:"amount"`
	DateTime string  `json:"datetime"`
}

type InquiryTransaction struct {
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

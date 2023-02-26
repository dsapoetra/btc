package repo

import "time"

type Transaction struct {
	Amount    float64   `db:"amount" json:"amount"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

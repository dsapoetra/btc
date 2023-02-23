package model

import "time"

type Transaction struct {
	ID        int64     `db:"id" json:"id"`
	Amount    float64   `db:"amount" json:"amount"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

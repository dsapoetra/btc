package repository

import (
	"btc/app/model/repo"
	"github.com/jmoiron/sqlx"
)

type TransactionRepository struct {
	db *sqlx.DB
}

type ITransactionRepository interface {
	AddTransaction(trx repo.Transaction) error
	//ListTransaction(startTime time.Time, endTime time.Time) ([]model.Transaction, error)
}

func NewTransactionRepository(db *sqlx.DB) ITransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (t *TransactionRepository) AddTransaction(trx repo.Transaction) error {
	query := `INSERT INTO transactions(amount, created_at) VALUES ($1, $2)`

	_, err := t.db.Exec(query, trx.Amount, trx.CreatedAt)

	return err
}

package repository

import (
	"btc/app/model/repo"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

type TransactionRepository struct {
	db *sqlx.DB
}

type ITransactionRepository interface {
	AddTransaction(trx repo.Transaction) error
	ListTransaction(startTime time.Time, endTime time.Time) (*[]repo.Transaction, error)
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

func (t *TransactionRepository) ListTransaction(startTime time.Time, endTime time.Time) (*[]repo.Transaction, error) {
	rows, err := t.db.Query(`SELECT date_trunc('hour', created_at) as Hour, sum(amount) as Avg FROM transactions where created_at  >= $1 and created_at  <= $2 GROUP BY date_trunc('hour', created_at);`, startTime, endTime)
	if err != nil {
		log.Println(err)
	}
	var results []repo.Transaction // creating empty slice
	defer rows.Close()
	for rows.Next() {
		trx := repo.Transaction{} // creating new struct for every row
		err = rows.Scan(&trx.CreatedAt, &trx.Amount)
		if err != nil {
			log.Println(err)
		}
		results = append(results, trx) // add new row information
	}

	//_, err := t.db.Exec(query, trx.Amount, trx.CreatedAt)

	return &results, nil
}

package service

import (
	"btc/app/model"
	"btc/app/repository"
	"log"
)

type TransactionService struct {
	db repository.ITransactionRepository
}

type ITransactionService interface {
	AddTransaction(trx model.Transaction) (*model.Transaction, error)
}

func NewTransactionService(repo repository.ITransactionRepository) ITransactionService {
	return &TransactionService{
		db: repo,
	}
}

func (t *TransactionService) AddTransaction(trx model.Transaction) (*model.Transaction, error) {
	log.Println("HERE 2")
	err := t.db.AddTransaction(trx)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

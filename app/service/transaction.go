package service

import (
	"btc/app/model/repo"
	"btc/app/repository"
	"errors"
)

type TransactionService struct {
	db repository.ITransactionRepository
}

type ITransactionService interface {
	AddTransaction(trx repo.Transaction) error
}

func NewTransactionService(repo repository.ITransactionRepository) ITransactionService {
	return &TransactionService{
		db: repo,
	}
}

func (t *TransactionService) AddTransaction(trx repo.Transaction) error {

	if trx.Amount <= 1 {
		return errors.New("amount must be greater than 1")
	}

	err := t.db.AddTransaction(trx)

	return err
}

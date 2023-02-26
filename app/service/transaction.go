package service

import (
	"btc/app/model/repo"
	"btc/app/repository"
	"context"
	"errors"
	"time"
)

type TransactionService struct {
	db repository.ITransactionRepository
}

type ITransactionService interface {
	AddTransaction(ctx context.Context, trx repo.Transaction) error
	ListTransaction(ctx context.Context, startTime time.Time, endTime time.Time) (*[]repo.Transaction, error)
}

func NewTransactionService(repo repository.ITransactionRepository) ITransactionService {
	return &TransactionService{
		db: repo,
	}
}

func (t *TransactionService) AddTransaction(ctx context.Context, trx repo.Transaction) error {

	if trx.Amount < 1 {
		return errors.New("amount must be greater than 1")
	}

	err := t.db.AddTransaction(ctx, trx)

	return err
}

func (t *TransactionService) ListTransaction(ctx context.Context, startTime time.Time, endTime time.Time) (*[]repo.Transaction, error) {

	res, err := t.db.ListTransaction(ctx, startTime, endTime)

	return res, err
}

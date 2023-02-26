package service

import (
	http_model "btc/app/model/http-model"
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
	AddTransaction(ctx context.Context, trx http_model.Transaction) error
	ListTransaction(ctx context.Context, startTimeStr, endTimeStr string) (*[]repo.Transaction, error)
}

func NewTransactionService(repo repository.ITransactionRepository) ITransactionService {
	return &TransactionService{
		db: repo,
	}
}

func (t *TransactionService) AddTransaction(ctx context.Context, trx http_model.Transaction) error {

	createdAt, err := time.Parse(time.RFC3339, trx.DateTime)

	if err != nil {
		return errors.New("invalid date time format, should be RFC3339 (2006-01-02T15:04:05+07:00)")
	}

	param := repo.Transaction{
		CreatedAt: createdAt.UTC(),
		Amount:    trx.Amount,
	}

	if param.Amount == 0 {
		return errors.New("amount can't be zero")
	}

	err = t.db.AddTransaction(ctx, param)

	return err
}

func (t *TransactionService) ListTransaction(ctx context.Context, startTimeStr, endTimeStr string) (*[]repo.Transaction, error) {

	startTime, err := time.Parse(time.RFC3339, startTimeStr)

	if err != nil {
		return nil, errors.New("invalid start time format, should be RFC3339 (2006-01-02T15:04:05+07:00)")
	}

	endTime, err := time.Parse(time.RFC3339, endTimeStr)

	if err != nil {
		return nil, errors.New("invalid end time format, should be RFC3339 (2006-01-02T15:04:05+07:00)")
	}

	res, err := t.db.ListTransaction(ctx, startTime.UTC(), endTime.UTC())

	return res, err
}

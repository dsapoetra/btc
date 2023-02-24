// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/repository/transaction.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	repo "btc/app/model/repo"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITransactionRepository is a mock of ITransactionRepository interface.
type MockITransactionRepository struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionRepositoryMockRecorder
}

// MockITransactionRepositoryMockRecorder is the mock recorder for MockITransactionRepository.
type MockITransactionRepositoryMockRecorder struct {
	mock *MockITransactionRepository
}

// NewMockITransactionRepository creates a new mock instance.
func NewMockITransactionRepository(ctrl *gomock.Controller) *MockITransactionRepository {
	mock := &MockITransactionRepository{ctrl: ctrl}
	mock.recorder = &MockITransactionRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionRepository) EXPECT() *MockITransactionRepositoryMockRecorder {
	return m.recorder
}

// AddTransaction mocks base method.
func (m *MockITransactionRepository) AddTransaction(trx repo.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTransaction", trx)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTransaction indicates an expected call of AddTransaction.
func (mr *MockITransactionRepositoryMockRecorder) AddTransaction(trx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTransaction", reflect.TypeOf((*MockITransactionRepository)(nil).AddTransaction), trx)
}

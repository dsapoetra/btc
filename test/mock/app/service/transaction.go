// Code generated by MockGen. DO NOT EDIT.
// Source: ./app/service/transaction.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	repo "btc/app/model/repo"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockITransactionService is a mock of ITransactionService interface.
type MockITransactionService struct {
	ctrl     *gomock.Controller
	recorder *MockITransactionServiceMockRecorder
}

// MockITransactionServiceMockRecorder is the mock recorder for MockITransactionService.
type MockITransactionServiceMockRecorder struct {
	mock *MockITransactionService
}

// NewMockITransactionService creates a new mock instance.
func NewMockITransactionService(ctrl *gomock.Controller) *MockITransactionService {
	mock := &MockITransactionService{ctrl: ctrl}
	mock.recorder = &MockITransactionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockITransactionService) EXPECT() *MockITransactionServiceMockRecorder {
	return m.recorder
}

// AddTransaction mocks base method.
func (m *MockITransactionService) AddTransaction(trx repo.Transaction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTransaction", trx)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTransaction indicates an expected call of AddTransaction.
func (mr *MockITransactionServiceMockRecorder) AddTransaction(trx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTransaction", reflect.TypeOf((*MockITransactionService)(nil).AddTransaction), trx)
}

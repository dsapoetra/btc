package handler

import (
	"btc/app/model/repo"
	mock_service "btc/test/mock/app/service"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Test_Handler_Transaction(t *testing.T) {
	type TransactionHandler struct {
		mockService *mock_service.MockITransactionService
	}

	tests := []struct {
		description string

		route string

		expectedError bool
		expectedCode  int
		expectedBody  string
		requestMethod string
		requestBody   string
		prepare       func(f *TransactionHandler)
	}{
		{
			description:   "success add transaction route",
			route:         "/v1/transaction",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"error\":false,\"msg\":\"Success inserted transaction\"}",
			requestMethod: "POST",
			requestBody:   "{\"amount\":100,\"created_at\":\"2006-01-02T00:00:00+00:00\"}",
			prepare: func(f *TransactionHandler) {
				createdAt := "2006-01-02T00:00:00+00:00"
				timeCreated, _ := time.Parse(time.RFC3339, createdAt)

				trxRepo := repo.Transaction{CreatedAt: timeCreated, Amount: 100}
				f.mockService.EXPECT().AddTransaction(trxRepo).Return(nil)

			},
		},
		{
			description:   "wrong datetime fotrmat add transaction route",
			route:         "/v1/transaction",
			expectedError: false,
			expectedCode:  400,
			expectedBody:  "{\"error\":true,\"msg\":\"parsing time \\\"2006-01-02T00:00:00\\\" as \\\"2006-01-02T15:04:05Z07:00\\\": cannot parse \\\"\\\" as \\\"Z07:00\\\"\"}",
			requestMethod: "POST",
			requestBody:   "{\"amount\":100,\"created_at\":\"2006-01-02T00:00:00\"}",
		},
		{
			description:   "amount is less than 1 add transaction route",
			route:         "/v1/transaction",
			expectedError: false,
			expectedCode:  400,
			expectedBody:  "{\"error\":true,\"msg\":\"amount must be greater than 1\"}",
			requestMethod: "POST",
			requestBody:   "{\"amount\":0,\"created_at\":\"2006-01-02T00:00:00+00:00\"}",
			prepare: func(f *TransactionHandler) {
				createdAt := "2006-01-02T00:00:00+00:00"
				timeCreated, _ := time.Parse(time.RFC3339, createdAt)

				trxRepo := repo.Transaction{CreatedAt: timeCreated, Amount: 0}
				f.mockService.EXPECT().AddTransaction(trxRepo).Return(errors.New("amount must be greater than 1"))

			},
		},
	}

	for _, test := range tests {

		t.Run(test.description, func(t *testing.T) {

			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()
			f := TransactionHandler{
				mockService: mock_service.NewMockITransactionService(mockCtrl),
			}

			if test.prepare != nil {
				test.prepare(&f)
			}

			app := fiber.New()
			api := app.Group("/v1")
			NewTransactionHandler(api.Group("/"), f.mockService)

			req, _ := http.NewRequest(test.requestMethod, test.route, strings.NewReader(test.requestBody))
			req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))
			req.Header.Set("Content-Type", "application/json")

			res, err := app.Test(req)

			assert.Equalf(t, test.expectedError, err != nil, test.description)

			assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

			body, err := io.ReadAll(res.Body)

			assert.Nilf(t, err, test.description)

			assert.Equalf(t, test.expectedBody, string(body), test.description)
		})
	}
}

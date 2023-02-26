package handler

import (
	http_model "btc/app/model/http-model"
	"btc/app/model/repo"
	mock_service "btc/test/mock/app/service"
	"context"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Test_Post_Transaction(t *testing.T) {
	type TransactionHandler struct {
		mockService *mock_service.MockITransactionService
	}

	timeNow, _ := time.Parse(time.RFC3339, "2023-02-26T22:05:52+00:00")
	timeNowInput := timeNow.Format(time.RFC3339)

	tests := []struct {
		description string

		route string

		expectedCode int
		expectedBody string
		requestBody  string
		prepare      func(f *TransactionHandler)
	}{
		{
			description:  "success add transaction route",
			route:        "/v1/transaction",
			expectedCode: 200,
			expectedBody: "{\"error\":false,\"msg\":\"Success inserted transaction\"}",
			requestBody:  "{\"amount\":100,\"datetime\":\"" + timeNowInput + "\"}",
			prepare: func(f *TransactionHandler) {
				trxHttp := http_model.Transaction{DateTime: timeNowInput, Amount: 100}
				f.mockService.EXPECT().AddTransaction(mock.MatchedBy(func(ctx context.Context) bool { return true }), trxHttp).Return(nil)

			},
		},
		{
			description:  "wrong datetime format add transaction route",
			route:        "/v1/transaction",
			expectedCode: 400,
			expectedBody: "{\"error\":true,\"msg\":\"invalid date time format, should be RFC3339 (2006-01-02T15:04:05+07:00)\"}",
			requestBody:  "{\"amount\":100,\"datetime\":\"2006-01-02T00:07:00\"}",
			prepare: func(f *TransactionHandler) {
				//timeFormat, _:= time.Parse("2006-01-02T15:04:05", )
				trxHttp := http_model.Transaction{DateTime: "2006-01-02T00:07:00", Amount: 100}
				f.mockService.EXPECT().AddTransaction(mock.MatchedBy(func(ctx context.Context) bool { return true }), trxHttp).Return(errors.New("invalid date time format, should be RFC3339 (2006-01-02T15:04:05+07:00)"))

			},
		},
		{
			description:  "amount is 0 add transaction route",
			route:        "/v1/transaction",
			expectedCode: 400,
			expectedBody: "{\"error\":true,\"msg\":\"amount can't be zero\"}",
			requestBody:  "{\"amount\":0,\"datetime\":\"" + timeNowInput + "\"}",
			prepare: func(f *TransactionHandler) {
				trxHttp := http_model.Transaction{DateTime: timeNowInput, Amount: 0}
				f.mockService.EXPECT().AddTransaction(mock.MatchedBy(func(ctx context.Context) bool { return true }), trxHttp).Return(errors.New("amount can't be zero"))

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

			req, _ := http.NewRequest("POST", test.route, strings.NewReader(test.requestBody))
			req.Header.Add("Content-Length", strconv.FormatInt(req.ContentLength, 10))
			req.Header.Set("Content-Type", "application/json")

			res, err := app.Test(req)

			assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

			body, err := io.ReadAll(res.Body)

			assert.Nilf(t, err, test.description)

			assert.Equalf(t, test.expectedBody, string(body), test.description)
		})
	}
}

func Test_Get_Transaction(t *testing.T) {
	type TransactionHandler struct {
		mockService *mock_service.MockITransactionService
	}

	timeNow, _ := time.Parse(time.RFC3339, "2023-02-26T22:05:52+00:00")
	timeNowInput := timeNow.Format(time.RFC3339)
	timeHourLater, _ := time.Parse(time.RFC3339, "2023-02-26T23:05:52+00:00")
	timeHourLaterInput := timeHourLater.Format(time.RFC3339)

	trxRepo := &[]repo.Transaction{
		{
			CreatedAt: timeNow, Amount: 100,
		},
		{
			CreatedAt: timeNow, Amount: 100,
		},
		{
			CreatedAt: timeHourLater, Amount: 100,
		},
		{
			CreatedAt: timeHourLater, Amount: 100,
		},
	}

	trxRepoJson, _ := json.Marshal(trxRepo)

	tests := []struct {
		description string

		route string

		expectedError         bool
		expectedCode          int
		expectedBody          string
		requestMethod         string
		requestQueryStartTime string
		requestQueryEndTime   string
		prepare               func(f *TransactionHandler)
	}{
		{
			description:           "success add transaction route",
			route:                 "/v1/transaction",
			expectedError:         false,
			expectedCode:          200,
			expectedBody:          string(trxRepoJson),
			requestQueryStartTime: timeNowInput,
			requestQueryEndTime:   timeHourLaterInput,
			prepare: func(f *TransactionHandler) {
				f.mockService.EXPECT().ListTransaction(mock.MatchedBy(func(ctx context.Context) bool { return true }), timeNowInput, timeHourLaterInput).Return(trxRepo, nil)
			},
		},
	}

	for _, test := range tests {

		t.Run(test.description, func(t *testing.T) {
			log.Println()
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

			req, err := http.NewRequest("GET", test.route, nil)
			if err != nil {
				log.Print(err)
				os.Exit(1)
			}

			q := req.URL.Query()
			q.Add("startDateTime", test.requestQueryStartTime)
			q.Add("endDateTime", test.requestQueryEndTime)
			req.URL.RawQuery = q.Encode()
			res, err := app.Test(req)

			assert.Equalf(t, test.expectedError, err != nil, test.description)

			assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

			body, err := io.ReadAll(res.Body)

			assert.Nilf(t, err, test.description)

			assert.Equalf(t, test.expectedBody, string(body), test.description)
		})
	}
}

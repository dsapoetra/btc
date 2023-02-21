package handler

import (
	mock_service "btc/test/mock/app/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func Test_Handler_createBankAccount(t *testing.T) {
	tests := []struct {
		description string

		route string

		expectedError bool
		expectedCode  int
		expectedBody  string
	}{
		{
			description:   "health route",
			route:         "/v1/health",
			expectedError: false,
			expectedCode:  200,
			expectedBody:  "{\"error\":false,\"msg\":null}",
		},
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockService := mock_service.NewMockIHealthService(mockCtrl)
	mockService.EXPECT().HealthCheck().Return(nil)

	app := fiber.New()
	api := app.Group("/v1")
	NewHealthHandler(api.Group("/"), mockService)

	for _, test := range tests {
		req, _ := http.NewRequest(
			"GET",
			test.route,
			nil,
		)

		res, err := app.Test(req, -1)

		assert.Equalf(t, test.expectedError, err != nil, test.description)

		if test.expectedError {
			continue
		}

		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)

		body, err := io.ReadAll(res.Body)

		assert.Nilf(t, err, test.description)

		assert.Equalf(t, test.expectedBody, string(body), test.description)
	}
}

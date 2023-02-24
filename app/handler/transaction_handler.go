package handler

import (
	"btc/app/model/http-model"
	"btc/app/model/repo"
	"btc/app/service"
	"github.com/gofiber/fiber/v2"
	"time"
)

func NewTransactionHandler(app fiber.Router, transactionService service.ITransactionService) {
	app.Post("/transaction", CreateTransaction(transactionService))
}

// CreateTransaction func create transaction.
// @Description Create transaction.
// @Summary create transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param transaction body model.TransactionRequest true "Transaction"
// @Success 200 {object} model.Transaction
// @Security ApiKeyAuth
// @Router /v1/transaction [post]
func CreateTransaction(transactionService service.ITransactionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := &http_model.Transaction{}

		if err := c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		createdAt, err := time.Parse(time.RFC3339, body.CreatedAt)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		param := repo.Transaction{
			CreatedAt: createdAt,
			Amount:    body.Amount,
		}

		err = transactionService.AddTransaction(param)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		// Return status 200 OK.
		return c.JSON(fiber.Map{
			"error": false,
			"msg":   "Success inserted transaction",
		})
	}
}

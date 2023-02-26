package handler

import (
	"btc/app/model/http-model"
	"btc/app/service"
	"github.com/gofiber/fiber/v2"
)

func NewTransactionHandler(app fiber.Router, transactionService service.ITransactionService) {
	app.Post("/transaction", CreateTransaction(transactionService))
	app.Get("/transaction", GetTransaction(transactionService))
}

// CreateTransaction func create transaction.
// @Description Create transaction.
// @Summary create transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param transaction body http_model.Transaction true "Transaction"
// @Success 200 {object} repo.Transaction
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

		err := transactionService.AddTransaction(c.Context(), *body)
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

// GetTransaction func get transaction.
// @Description Get transaction.
// @Summary get transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param startDateTime query string true "Start Time"
// @Param endDateTime query string true "End Time"
// @Success 200 {object} []repo.Transaction
// @Security ApiKeyAuth
// @Router /v1/transaction [get]
func GetTransaction(transactionService service.ITransactionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		st := c.Query("startDateTime")
		et := c.Query("endDateTime")

		res, err := transactionService.ListTransaction(c.Context(), st, et)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		// Return status 200 OK.
		return c.JSON(res)
	}
}

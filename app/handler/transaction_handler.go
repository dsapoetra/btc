package handler

import (
	"btc/app/model"
	"btc/app/service"
	"github.com/gofiber/fiber/v2"
	"log"
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
// @Param transaction body model.Transaction true "Transaction"
// @Success 200 {object} model.Transaction
// @Security ApiKeyAuth
// @Router /v1/transaction [post]
func CreateTransaction(transactionService service.ITransactionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := &model.Transaction{}

		if err := c.BodyParser(body); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
		log.Println("HERE 1")
		trx, err := transactionService.AddTransaction(*body)
		if err != nil {
			// Return, if book not found.
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   "error when inputting author",
			})
		}
		log.Println("HERE 2")

		// Return status 200 OK.
		return c.JSON(fiber.Map{
			"error": false,
			"msg":   trx,
		})
	}
}

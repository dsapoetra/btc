package handler

import (
	"btc/app/model/http-model"
	"btc/app/model/repo"
	"btc/app/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
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

// GetTransaction func create transaction.
// @Description Create transaction.
// @Summary create transaction
// @Tags Transaction
// @Accept json
// @Produce json
// @Param start_time query string true "Start Time"
// @Param end_time query string true "End Time"
// @Success 200 {object} repo.Transaction
// @Security ApiKeyAuth
// @Router /v1/transaction [get]
func GetTransaction(transactionService service.ITransactionService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		st := c.Query("start_time")
		et := c.Query("end_time")

		log.Println("Start Time string: ", st)

		startTime, err := time.Parse(time.RFC3339, st)

		log.Println("Start Time: ", startTime)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		endTime, err := time.Parse(time.RFC3339, et)
		log.Println("End Time: ", endTime)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		res, err := transactionService.ListTransaction(startTime, endTime)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}

		// Return status 200 OK.
		return c.JSON(fiber.Map{
			"error": false,
			"msg":   res,
		})
	}
}

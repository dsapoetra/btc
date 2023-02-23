package main

import (
	"btc/app/handler"
	"btc/app/repository"
	"btc/app/routes"
	"btc/app/service"
	"btc/pkg"
	"btc/pkg/db"
	"btc/pkg/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"

	_ "btc/docs"                          // load API Docs files (Swagger)
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func main() {
	// Define Fiber config.
	config := pkg.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	db, err := db.PostgreSQLConnection()

	if err != nil {
		log.Println("Error connecting to database: ", err)
	}

	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)

	healthService := service.NewHealthService()

	// Routes.
	routes.SwaggerRoute(app) // Register a route for API Docs (Swagger).

	api := app.Group("/v1")

	// Prepare our endpoints for the API.
	handler.NewHealthHandler(api.Group("/"), healthService)
	handler.NewTransactionHandler(api.Group("/"), transactionService)

	app.Use(
		// Add CORS to each route.
		cors.New(),
		// Add simple logger.
		logger.New(),
	)

	// Start server (with graceful shutdown).
	server.StartServerWithGracefulShutdown(app)
}

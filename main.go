package main

import (
	"btc/app/handler"
	"btc/app/routes"
	"btc/app/service"
	"btc/pkg/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	_ "btc/docs"                          // load API Docs files (Swagger)
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

func main() {
	// Define Fiber config.
	//config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New()

	//db, _ := database.PostgreSQLConnection()
	//articleRepo := repository.NewArticleRepository(db)
	//articleSvc := service.NewArticleService(articleRepo)

	//authorRepo := repository.NewAuthorRepository(nil)
	//authorSvc := service.NewAuthorService(authorRepo)

	healthService := service.NewHealthService()

	//// Middlewares.
	//middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	//// Routes.
	routes.SwaggerRoute(app) // Register a route for API Docs (Swagger).
	//routes.PrivateRoutes(app) // Register a private routes for app.
	//routes.NotFoundRoute(app) // Register route for 404 Error.

	api := app.Group("/v1")
	//app.Get("/health", func(ctx *fiber.Ctx) error {
	//	return ctx.JSON("ok")
	//})

	// Prepare our endpoints for the API.
	handler.NewHealthHandler(api.Group("/"), healthService)

	app.Use(
		// Add CORS to each route.
		cors.New(),
		// Add simple logger.
		logger.New(),
	)

	// Start server (with graceful shutdown).
	server.StartServerWithGracefulShutdown(app)
}

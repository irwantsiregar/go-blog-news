package app

import (
	"bwanews/config"
	"bwanews/internal/adapter/handler"
	"bwanews/internal/adapter/repository"
	"bwanews/internal/core/service"
	"bwanews/lib/auth"
	"bwanews/lib/middleware"
	"bwanews/lib/pagination"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

func RunServer() {
	cfg := config.NewConfig()

	db, err := cfg.ConnectionPostgres()

	if err != nil {
		log.Fatal().Msgf("Error to connecting to database %v", err)
		return
	}

	// CloudflareR2
	crfR2 := cfg.LoadAwsConfig()
	_ = s3.NewFromConfig(crfR2)

	jwt := auth.NewJwt(cfg)

	middlewareAuth := middleware.NewMiddleware(cfg)

	_ = pagination.NewPagination()


	// Repository
	authrepo := repository.NewAuthRepository(db.DB)
	categoryRepo := repository.NewCategoryRepository(db.DB)

	// Service
	authService := service.NewAuthService(authrepo, cfg, jwt)
	categoryService := service.NewCategoryService(categoryRepo)

	// Handler
 	authHandler := handler.NewAuthHandler(authService)
	categoryHandler := handler.NewCategoryHandler(categoryService)


	// Instance Go Fiber
	app := fiber.New()
	app.Use(cors.New())
	app.Use(recover.New())

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${ip} ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		TimeZone:   "Asia/Jakarta",
	}))

	api := app.Group("/api")
	api.Post("/login", authHandler.Login)

	adminApp := api.Group("/admin")
	adminApp.Use(middlewareAuth.CheckToken())

	// Category
	categoryApp := adminApp.Group("/categories")
	categoryApp.Get("/", categoryHandler.GetCategories)
	categoryApp.Post("/", categoryHandler.CreateCategory)
	categoryApp.Put("/:categoryID", categoryHandler.EditCategoryByID)
	categoryApp.Post("/:categoryID", categoryHandler.GetCategoryByID)

	go func() {
		if cfg.App.AppPort == "" {
			cfg.App.AppPort = os.Getenv(("APP_PORT"))
		}

		err := app.Listen(":" + cfg.App.AppPort)

		if err != nil {
			log.Fatal().Msgf("Error to starting server %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)

	<-quit

	log.Info().Msg("server shutdown of 5 seconds")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	app.ShutdownWithContext(ctx)
}
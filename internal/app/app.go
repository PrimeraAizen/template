package app

import (
	"context"
	"fmt"

	"github.com/PrimeraAizen/template/config"
	"github.com/PrimeraAizen/template/internal/delivery"
	"github.com/PrimeraAizen/template/internal/repository"
	"github.com/PrimeraAizen/template/internal/server"
	"github.com/PrimeraAizen/template/internal/service"
	postgres "github.com/PrimeraAizen/template/pkg/adapter"
	"github.com/PrimeraAizen/template/pkg/logger"
)

func StartWebServer(ctx context.Context, cfg *config.Config, appLogger *logger.Logger) error {
	appLogger.WithComponent("app").Info("Initializing web server")

	// Initialize database connection
	appLogger.WithComponent("database").Info("Connecting to database")
	pg, err := postgres.New(ctx, &cfg.PG)
	if err != nil {
		appLogger.WithComponent("database").WithError(err).Error("Failed to initialize database connection")
		return fmt.Errorf("could not init postgres connection: %w", err)
	}
	defer func() {
		appLogger.WithComponent("database").Info("Closing database connection")
		pg.Close()
	}()

	appLogger.WithComponent("database").Info("Database connection established")

	// Initialize repositories
	appLogger.WithComponent("repository").Info("Initializing repositories")
	repos := repository.NewRepositories(pg)

	// Initialize services
	appLogger.WithComponent("service").Info("Initializing services")
	services := service.NewServices(service.Deps{
		Repos:  repos,
		Config: cfg,
	})

	// Initialize handlers
	appLogger.WithComponent("handler").Info("Initializing handlers")
	handlers := delivery.NewHandler(services, appLogger)

	// Initialize server
	appLogger.WithComponent("server").Info("Initializing HTTP server")
	srv := server.NewServer(cfg, handlers.Init(cfg), appLogger)

	// Start server
	appLogger.WithComponent("server").WithFields(logger.Fields{
		"host": cfg.Http.Host,
		"port": cfg.Http.Port,
	}).Info("Starting HTTP server")

	defer func() {
		appLogger.WithComponent("server").Info("Stopping HTTP server")
		srv.Stop()
	}()

	srv.Run()
	appLogger.WithComponent("server").Info("HTTP server started successfully")

	// Wait for context cancellation
	<-ctx.Done()
	appLogger.WithComponent("app").Info("Received shutdown signal")

	return nil
}

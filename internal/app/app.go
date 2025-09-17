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
)

func StartWebServer(ctx context.Context, cfg *config.Config) error {
	pg, err := postgres.New(ctx, &cfg.PG)
	if err != nil {
		return fmt.Errorf("could not init postgres connection: %w", err)
	}
	defer pg.Close()
	defer pg.Close()

	repos := repository.NewRepositories(pg)
	services := service.NewServices(service.Deps{
		Repos:  repos,
		Config: cfg,
	})

	handlers := delivery.NewHandler(services)

	srv := server.NewServer(cfg, handlers.Init(cfg))

	srv.Run()
	defer srv.Stop()
	fmt.Println("Web server started!")

	<-ctx.Done()

	return nil
}

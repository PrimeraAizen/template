package service

import (
	"github.com/PrimeraAizen/template/config"
	"github.com/PrimeraAizen/template/internal/repository"
)

type Service struct {
	ExampleService Example
	HealthService  Health
}

type Deps struct {
	Repos  *repository.Repository
	Config *config.Config
}

func NewServices(deps Deps) *Service {
	return &Service{
		ExampleService: NewExampleService(deps.Repos.Example),
		HealthService:  NewHealthService(deps.Repos.Health),
	}
}

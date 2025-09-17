package repository

import postgres "github.com/PrimeraAizen/template/pkg/adapter"

type Repository struct {
	Example Example
	Health  Health
}

func NewRepositories(pg *postgres.Postgres) *Repository {
	return &Repository{
		Example: NewExampleRepository(pg),
		Health:  NewHealthRepository(pg),
	}
}

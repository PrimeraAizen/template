package repository

import (
	"context"

	postgres "github.com/PrimeraAizen/template/pkg/adapter"
)

type Example interface {
	ExampleMethod() error
}

type Health interface {
	Ping(ctx context.Context) error
}

type ExampleRepository struct {
	pg *postgres.Postgres
}

func NewExampleRepository(pg *postgres.Postgres) *ExampleRepository {
	return &ExampleRepository{
		pg: pg,
	}
}

func (e *ExampleRepository) ExampleMethod() error {
	return nil
}

type HealthRepository struct {
	pg *postgres.Postgres
}

func NewHealthRepository(pg *postgres.Postgres) *HealthRepository {
	return &HealthRepository{pg: pg}
}

func (r *HealthRepository) Ping(ctx context.Context) error {
	// Cheap ping using pool.
	return r.pg.Pool.Ping(ctx)
}

package graph

import "github.com/umerm-work/hatch-assignment/infrastructure/postgres"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

func BuildResolver(
	s *postgres.DB,
) *Resolver {
	return &Resolver{
		storage: s,
	}
}

type Resolver struct {
	storage *postgres.DB
}

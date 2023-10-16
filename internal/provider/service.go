package provider

import (
	"context"

	"github.com/andsanchez/DERES_Back/internal/domain"
)

type Service interface {
	Get(ctx context.Context, name string) (domain.Provider, error)
	Create(ctx context.Context, name string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Get(ctx context.Context, name string) (domain.Provider, error) {
	provider, err := s.repository.Get(ctx, name)
	if err != nil {
		return domain.Provider{}, err
	}
	return provider, nil
}

func (s *service) Create(ctx context.Context, name string) error {
	providerToCreate := domain.Provider{
		Name: name,
	}

	err := s.repository.Create(ctx, providerToCreate)
	if err != nil {
		return err
	}
	return nil
}

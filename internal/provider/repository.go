package provider

import (
	"context"
	"errors"

	"github.com/andsanchez/DERES_Back/internal/domain"
)

// Repository encapsulates the storage of a section.
type Repository interface {
	Get(ctx context.Context, name string) (domain.Provider, error)
	Create(ctx context.Context, user domain.Provider) error
}

type repository struct {
	providerList []domain.Provider
}

func NewRepository() Repository {
	return &repository{
		providerList: []domain.Provider{},
	}
}

func (r *repository) Get(ctx context.Context, name string) (domain.Provider, error) {
	for _, provider := range r.providerList {
		if provider.Name == name {
			return provider, nil
		}
	}
	return domain.Provider{}, errors.New("provider was not found")
}

func (r *repository) Create(ctx context.Context, provider domain.Provider) error {
	r.providerList = append(r.providerList, provider)
	return nil
}

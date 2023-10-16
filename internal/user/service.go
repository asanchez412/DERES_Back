package user

import (
	"context"

	"github.com/andsanchez/DERES_Back/internal/domain"
)

type Service interface {
	Login(ctx context.Context, username, password string) error
	SignUp(ctx context.Context, username, password string) error
}

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Login(ctx context.Context, username, password string) error {
	_, err := s.repository.Get(ctx, username, password)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) SignUp(ctx context.Context, username, password string) error {
	userToCreate := domain.User{
		Username: username,
		Password: password,
	}

	err := s.repository.Create(ctx, userToCreate)
	if err != nil {
		return err
	}
	return nil
}

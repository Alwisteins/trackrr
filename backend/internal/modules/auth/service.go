package auth

import (
	"context"
	"errors"

	"backend/internal/utils"
)

type Service interface {
	Login(ctx context.Context, email, password string) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Login(ctx context.Context, identifier, password string) (*User, error) {
	user, err := s.repo.FindByIdentifier(ctx, identifier)
	if err != nil {
		return nil, errors.New("email atau password salah")
	}

	// compare password
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("email atau password salah")
	}

	return user, nil
}

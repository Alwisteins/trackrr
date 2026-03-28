package auth

import (
	"context"

	"backend/internal/errors"
	"backend/internal/utils"
)

type Service interface {
	Register(ctx context.Context, req *RegisterRequest) (*User, string, error)
	Login(ctx context.Context, email, password string) (*User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) Register(ctx context.Context, req *RegisterRequest) (*User, string, error) {
	user, err := s.repo.Register(ctx, req)
	if err != nil {
		return nil, "", err
	}

	token, err := utils.GenerateToken(user.UUID)
	if err != nil {
		return nil, "", errors.NewInternalError("TOKEN_ERROR", "Failed to generate token", err)
	}

	return user, token, nil
}

func (s *service) Login(ctx context.Context, identifier, password string) (*User, error) {
	user, err := s.repo.FindByIdentifier(ctx, identifier)
	if err != nil {
		return nil, errors.NewUnauthorizedError("INVALID_CREDENTIALS", "email atau password salah")
	}

	// compare password
	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.NewUnauthorizedError("INVALID_CREDENTIALS", "email atau password salah")
	}

	return user, nil
}

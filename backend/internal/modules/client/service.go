package client

import (
	"context"
)

type Service interface {
	CreateClient(ctx context.Context, req CreateClientRequest) (*Client, error)
	FindAllClients(ctx context.Context) ([]Client, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateClient(ctx context.Context, req CreateClientRequest) (*Client, error) {
	client, err := s.repo.CreateClient(ctx, req)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (s *service) FindAllClients(ctx context.Context) ([]Client, error) {
	clients, err := s.repo.FindAllClients(ctx)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

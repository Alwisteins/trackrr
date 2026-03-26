package client

import (
	"context"
)

type Service interface {
	CreateClient(ctx context.Context, req ClientRequest) (*Client, error)
	FindAllClients(ctx context.Context) ([]*Client, error)
	FindClientByUUID(ctx context.Context, uuid string) (*Client, error)
	UpdateClient(ctx context.Context, uuid string, req ClientRequest) (*Client, error)
	DeleteClient(ctx context.Context, uuid string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo}
}

func (s *service) CreateClient(ctx context.Context, req ClientRequest) (*Client, error) {
	client, err := s.repo.CreateClient(ctx, req)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (s *service) FindAllClients(ctx context.Context) ([]*Client, error) {
	clients, err := s.repo.FindAllClients(ctx)
	if err != nil {
		return nil, err
	}

	return clients, nil
}

func (s *service) FindClientByUUID(ctx context.Context, uuid string) (*Client, error) {
	client, err := s.repo.FindClientByUUID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (s *service) UpdateClient(ctx context.Context, uuid string, req ClientRequest) (*Client, error) {
	client, err := s.repo.UpdateClient(ctx, uuid, req)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (s *service) DeleteClient(ctx context.Context, uuid string) error {
	return s.repo.DeleteClient(ctx, uuid)
}

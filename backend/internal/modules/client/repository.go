package client

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"backend/internal/errors"
)

type Repository interface {
	CreateClient(ctx context.Context, req CreateClientRequest) (*Client, error)
	FindAllClients(ctx context.Context) ([]Client, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateClient(ctx context.Context, req CreateClientRequest) (*Client, error) {
	// 1. Mapping request data to Client model
	newClient := Client{
		UUID:         uuid.New().String(),
		CompanyName:  req.CompanyName,
		ContactName:  req.ContactName,
		ContactEmail: req.ContactEmail,
		ContactPhone: req.ContactPhone,
		Notes:        req.Notes,
	}

	// 2. Saving the new client to the database
	result := r.db.WithContext(ctx).Create(&newClient)
	if result.Error != nil {
		return nil, errors.HandleGormError(result.Error, "client")
	}

	return &newClient, nil
}

func (r *repository) FindAllClients(ctx context.Context) ([]Client, error) {
	var clients []Client
	result := r.db.WithContext(ctx).Find(&clients)
	if result.Error != nil {
		return nil, errors.HandleGormError(result.Error, "client")
	}
	return clients, nil
}

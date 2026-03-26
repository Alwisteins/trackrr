package client

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"backend/internal/errors"
)

type Repository interface {
	CreateClient(ctx context.Context, req ClientRequest) (*Client, error)
	FindAllClients(ctx context.Context) ([]*Client, error)
	FindClientByUUID(ctx context.Context, uuid string) (*Client, error)
	UpdateClient(ctx context.Context, uuid string, req ClientRequest) (*Client, error)
	DeleteClient(ctx context.Context, uuid string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) CreateClient(ctx context.Context, req ClientRequest) (*Client, error) {
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

func (r *repository) FindAllClients(ctx context.Context) ([]*Client, error) {
	var clients []*Client
	result := r.db.WithContext(ctx).Find(&clients)
	if result.Error != nil {
		return nil, errors.HandleGormError(result.Error, "client")
	}
	return clients, nil
}

func (r *repository) FindClientByUUID(ctx context.Context, uuid string) (*Client, error) {
	var client *Client
	result := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&client)
	if result.Error != nil {
		return nil, errors.HandleGormError(result.Error, "client")
	}
	return client, nil
}

func (r *repository) UpdateClient(ctx context.Context, uuid string, req ClientRequest) (*Client, error) {
	// 1. Mapping request data to Client model
	updateClient := Client{
		CompanyName:  req.CompanyName,
		ContactName:  req.ContactName,
		ContactEmail: req.ContactEmail,
		ContactPhone: req.ContactPhone,
		Notes:        req.Notes,
	}

	// 2. Update the client in the database
	result := r.db.WithContext(ctx).
		Model(&Client{}).
		Where("uuid = ?", uuid).
		Updates(updateClient)
	if result.Error != nil {
		return nil, errors.HandleGormError(result.Error, "client")
	}

	if result.RowsAffected == 0 {
		return nil, errors.NewNotFoundError("RECORD_NOT_FOUND", "client not found")
	}

	// 3. Fetch updated client and return
	var updated Client
	findErr := r.db.WithContext(ctx).Where("uuid = ?", uuid).First(&updated).Error
	if findErr != nil {
		return nil, errors.HandleGormError(findErr, "client")
	}

	return &updated, nil
}

func (r *repository) DeleteClient(ctx context.Context, uuid string) error {
	result := r.db.WithContext(ctx).Delete(&Client{}, "uuid = ?", uuid)
	if result.Error != nil {
		return errors.HandleGormError(result.Error, "client")
	}
	if result.RowsAffected == 0 {
		return errors.NewNotFoundError("RECORD_NOT_FOUND", "client not found")
	}
	return nil
}

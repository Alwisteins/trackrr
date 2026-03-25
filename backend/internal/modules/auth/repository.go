package auth

import (
	"context"

	"gorm.io/gorm"

	"backend/internal/errors"
)

type Repository interface {
	FindByIdentifier(ctx context.Context, identifier string) (*User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindByIdentifier(ctx context.Context, identifier string) (*User, error) {
	user := &User{}
	err := r.db.WithContext(ctx).
		Where("email = ? OR username = ?", identifier, identifier).
		First(user).Error

	if err != nil {
		return nil, errors.HandleGormError(err, "user")
	}

	return user, nil
}

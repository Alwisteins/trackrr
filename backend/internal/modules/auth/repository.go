package auth

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"backend/internal/errors"
	"backend/internal/utils"
)

type Repository interface {
	Register(ctx context.Context, req *RegisterRequest) (*User, error)
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

func (r *repository) Register(ctx context.Context, req *RegisterRequest) (*User, error) {
	// 1. Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, errors.NewInternalError("HASH_ERROR", "Failed to process password", err)
	}

	// 2. Mapping request data to User model
	newUser := &User{
		UUID:     uuid.New().String(),
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	// 2. Saving the new user to the database
	result := r.db.WithContext(ctx).Create(&newUser)
	if result.Error != nil {
		return nil, errors.HandleGormError(result.Error, "user")
	}

	return newUser, nil
}

package auth

import (
	"context"
	"database/sql"
)

type Repository interface {
	FindByIdentifier(ctx context.Context, identifier string) (*User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) FindByIdentifier(ctx context.Context, identifier string) (*User, error) {
	query := `
		SELECT id, email, password, role
		FROM users
		WHERE email = $1 OR WHERE username = $1
	`

	user := &User{}
	err := r.db.QueryRowContext(ctx, query, identifier).
		Scan(&user.ID, &user.Email, &user.Password, &user.Role)

	if err != nil {
		return nil, err
	}

	return user, nil
}

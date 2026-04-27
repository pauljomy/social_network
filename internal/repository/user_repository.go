package repository

import (
	"context"

	"github.com/pauljomy/social/internal/models"
)

type UserRepository interface {
	Create(ctx context.Context) error
	FindByID(ctx context.Context, id string) (*models.User, error)
}

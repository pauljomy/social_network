package store

import (
	"context"

	"github.com/pauljomy/social/internal/models"
)

type PostStorer interface {
	Create(ctx context.Context, post *models.Post) error
}

type CommentStorer interface {
	Create(ctx context.Context, comment *models.Comment) error
}

type UserStorer interface {
	Create(ctx context.Context) error
	FindByID(ctx context.Context, id string) (*models.User, error)
}

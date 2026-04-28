package store

import (
	"context"

	"github.com/pauljomy/social/internal/models"
)

type PostStorer interface {
	Create(context.Context, *models.Post) error
}

type CommentStorer interface {
	Create(context.Context, *models.Comment) error
}

type UserStorer interface {
	Create(context.Context, *models.User) error
	FindByID(context.Context, string) (*models.User, error)
}

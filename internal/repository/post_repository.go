package repository

import (
	"context"

	"github.com/pauljomy/social/internal/models"
)

type PostRepository interface {
	Create(ctx context.Context, post *models.Post) error
}

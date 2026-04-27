package repository

import (
	"context"

	"github.com/pauljomy/social/internal/models"
)

type CommentRepository interface {
	Create(ctx context.Context, comment *models.Comment) error
}

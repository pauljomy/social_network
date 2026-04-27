package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/pauljomy/social/internal/models"
)

type postStore struct {
	db *pgxpool.Pool
}

func (s *postStore) Create(ctx context.Context, post *models.Post) error {
	return nil
}
